import reactor, capnp, collections/pprint, collections, reactor/redis, times, isa/erasure_code, isa/hash, securehash, reactor/threading, os, snappy, isa/aes, hashes
import tlog/schema, tlog/util

type
  CoreState = ref object
    dbConnections: seq[RedisClient]
    hasher: Sha512Pool
    volumes: Table[string, VolumeHandler]

  VolumeHandler = ref object
    volumeId: string
    lastHash: string
    waitingBlocks: seq[TlogBlock]
    waitingSegments: seq[ByteView]
    flushing: bool

var coreState {.threadvar.}: CoreState

when defined(enableEncryption):
  let aesKey = expandAESKey("0".repeat(32))
const zeroIV = "0".repeat(32)

{.passl: "-lb2".}

proc blake2b(output: cstring, input: cstring, key: cstring, outlen: int,
             inlen: int, keylen: int): cint {.importc.}

proc blakeHash(data: string): string =
  result = newString(32)
  let err = blake2b(result, data, "", result.len, data.len, 0)
  assert err == 0

proc flush(volume: VolumeHandler): Future[void] {.async.} =
  if volume.flushing:
    return

  volume.flushing = true
  defer:
    volume.flushing = false

  if volume.waitingBlocks.len == 0:
    return

  if volume.lastHash == nil:
    # lastHash not yet fetched
    let lastHash = await coreState.dbConnections[0].hget("volume-hash", $(volume.volumeId))
    volume.lastHash = lastHash
    # if volume is still nil, this volume is empty

  let agg = TlogAggregation(blocks: volume.waitingBlocks,
                            volumeId: volume.volumeId,
                            prev: volume.lastHash)
  let packer = newPacker(initialBufferSize=17000 * volume.waitingBlocks.len)
  packPointer(packer, 0, agg)

  let aggDataCompressed = snappy.compress(packer.buffer)
  let aggHash = $(blakeHash(aggDataCompressed))

  const
    totalChunks = 10
    maxLost = 3

  var aggDataEncrypted = aggDataCompressed
  when defined(enableEncryption):
    aggDataEncrypted.setLen(aggDataEncrypted.len + 32 - (aggDataEncrypted.len mod 32)) # padding
    aggDataEncrypted = aesKey.encrypt(aggDataEncrypted)

  var chunks = erasure_code.encodeString(aggDataCompressed, totalChunks, maxLost)
  var waitFor: seq[Future[int64]] = @[]

  when defined(timeRedis):
    let start = epochTime()

  for i in 0..<totalChunks: # store in parallel
    waitFor.add coreState.dbConnections[i].hset("volume-data", aggHash, chunks[i])
  # echo "waiting..."
  for fut in waitFor: discard (await fut)
  # echo "saved (", waitFor.len, " chunks, ", volume.waitingBlocks.len, " blocks)"

  when defined(timeRedis):
    echo "saving data to ARDB took ", (epochTime() - start) * 1000, " ms"

  discard (await coreState.dbConnections[0].hset("volume-hash", $(volume.volumeId), aggHash))
  volume.lastHash = aggHash

  for view in volume.waitingSegments:
    deallocShared(view.data)

  volume.waitingSegments = @[]
  volume.waitingBlocks = @[]

proc handleMsgProc(segments: seq[ByteView]): auto =
  return proc() =
           let blockMsg = capnp.newUnpacker(segments).unpackPointer(0, TlogBlock)
           if blockMsg.volumeId notin coreState.volumes:
             echo "cache miss: ", blockMsg.volumeId
             coreState.volumes[blockMsg.volumeId] = VolumeHandler(waitingBlocks: @[], volumeId: blockMsg.volumeId,
                                                                  waitingSegments: @[])

           let vol = coreState.volumes[blockMsg.volumeId]
           vol.waitingBlocks.add(blockMsg)
           vol.waitingSegments &= segments
           if vol.waitingBlocks.len > 100:
             vol.flush.ignore

proc handleClient(conn: TcpConnection) {.async.} =
  while true:
    let segmentsR = tryAwait readMultisegment(conn.input)
    #let segmentsR = tryAwait conn.input.read(16472) # 1.05 s
    if segmentsR.isError and segmentsR.error.getOriginal == JustClose: # eof?
      break

    var segments = segmentsR.get

    let blockMsg = capnp.newUnpacker(segments).unpackPointer(0, TlogBlock)
    when defined(disableThreading):
      handleMsgProc(segments)()
    else:
      # blockMsg.volumeId.int mod threadLoopCount() == threadLoopId()
      runOnThread((blockMsg.volumeId.hash mod threadLoopCount()), handleMsgProc(segments))

    let resp = TlogResponse(status: 0, sequences: @[#[blockMsg.sequence]#])
    let data = packPointer(resp)
    await conn.output.write($(data.len + 8) & "\r\n")
    await writeMultisegment(conn.output, data)

proc connectToRedis() {.async.} =
  for i in 1..10:
    let conn = await redis.connect("localhost", 11000 + i, reconnect=true)
    coreState.dbConnections.add(conn)

proc flushAll() {.async.} =
  for vol in coreState.volumes.values:
    await vol.flush

proc periodicFlusher() {.async.} =
  while true:
    await flushAll()
    await asyncSleep(1000)

proc loopMain() {.async.} =
  echo "init"
  coreState = CoreState(dbConnections: @[], hasher: newSha512Pool(),
                        volumes: initTable[string, VolumeHandler]())
  await connectToRedis()
  let server = await createTcpServer(11211, reusePort=true)

  periodicFlusher().ignore

  asyncFor conn in server.incomingConnections:
    echo "incoming connection ", conn.getPeerAddr
    handleClient(conn).ignore

proc main*() =
  when defined(disableThreading):
    loopMain().runMain
  else:
    startMultiloop(mainProc=loopMain, threadCount=4)

when isMainModule:
  main()

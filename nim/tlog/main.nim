import reactor, capnp, collections/pprint, collections, reactor/redis, times, isa/erasure_code, isa/hash, securehash, reactor/threading, os, snappy
import tlog/schema, tlog/util

type
  CoreState = ref object
    dbConnections: seq[RedisClient]

  VolumeHandler = ref object
    volumeId: uint32
    lastHash: string
    waitingBlocks: seq[TlogBlock]

let hasher = newSha512Pool()
var volumes = initTable[uint32, VolumeHandler]()
var coreState {.threadvar.}: CoreState

proc flush(volume: VolumeHandler): Future[void] {.async.} =
  if volume.lastHash == nil:
    # lastHash not yet fetched
    let lastHash = await coreState.dbConnections[0].hget("volume-hash", $(volume.volumeId))
    volume.lastHash = lastHash
    # if volume is still nil, this volume is empty

  let agg = TlogAggregation(blocks: volume.waitingBlocks,
                            volumeId: volume.volumeId,
                            prev: volume.lastHash)
  let aggData = packPointer(agg)
  let aggHash = $(secureHash(aggData)) # TODO: use SHA256

  const
    totalChunks = 10
    maxLost = 3

  let aggDataCompressed = snappy.compress(aggData)
  let chunks = encodeString(aggDataCompressed, totalChunks, maxLost)
  var waitFor: seq[Future[int64]] = @[]
  for i in 0..<totalChunks: # store in parallel
    waitFor.add coreState.dbConnections[i].hset("volume-data", aggHash, chunks[i])
  echo "waiting..."
  for fut in waitFor: discard (await fut)
  echo "saved"

  discard (await coreState.dbConnections[0].hset("volume-hash", $(volume.volumeId), aggHash))

proc handleMsgProc(blockMsg: TlogBlock): auto =
  return proc() =
           if blockMsg.volumeId notin volumes:
             echo "cache miss: ", blockMsg.volumeId
             volumes[blockMsg.volumeId] = VolumeHandler(waitingBlocks: @[], volumeId: blockMsg.volumeId)

           let vol = volumes[blockMsg.volumeId]
           vol.waitingBlocks.add(blockMsg)
           vol.flush.ignore

proc handleClient(conn: TcpConnection) {.async.} =
  while true:
    let segmentsR = tryAwait readMultisegment(conn.input)
    if segmentsR.isError and segmentsR.error.getOriginal == JustClose: # eof?
      break

    let segments = await segmentsR
    let blockMsg = capnp.newUnpacker(segments).unpackPointer(0, TlogBlock)
    echo "recv ", blockMsg.pprint
    runOnThread((blockMsg.volumeId.int mod threadLoopCount()), handleMsgProc(blockMsg))

proc connectToRedis() {.async.} =
  for i in 1..10:
    let conn = await redis.connect("localhost", 11000 + i, reconnect=true)
    coreState.dbConnections.add(conn)

proc loopMain() {.async.} =
  echo "init"
  coreState = CoreState(dbConnections: @[])
  await connectToRedis()
  let server = await createTcpServer(11211, reusePort=true)
  asyncFor conn in server.incomingConnections:
    echo "incoming connection ", conn.getPeerAddr
    handleClient(conn).ignore

proc main*() =
  startMultiloop(mainProc=loopMain)

when isMainModule:
  main()

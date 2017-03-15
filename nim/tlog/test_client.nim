import tlog/schema, reactor, capnp, strutils

proc main*() {.async.} =
  let conn = await connectTcp("localhost", 11211)
  let data = packPointer(TlogBlock(
    volumeId: 0,
    sequence: 1,
    lba: 0,
    size: 0,
    data: " ".repeat(1024),
    timestamp: 0))

  await conn.output.writeItem(uint32(0), littleEndian)
  await conn.output.writeItem(uint32(data.len div 8), littleEndian)
  await conn.output.write(data)

  await asyncSleep(1000)

when isMainModule:
  main().runMain

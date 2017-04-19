import tlog/schema, tlog/util, reactor, capnp, strutils, collections

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

  let respData = await conn.input.readMultisegment
  let resp = newUnpacker(respData).unpackPointer(0, TlogResponse)
  echo resp.pprint

when isMainModule:
  main().runMain

import reactor, tlog/util, strutils, capnp, tlog/schema, strutils, os

proc main() {.async.} =
  let conn = await connectTcp("localhost", 11211)
  let data = " ".repeat(4096)
  let dataMsg = packPointer(TlogBlock(
    volumeId: 0,
    sequence: 1,
    lba: 0,
    size: 0,
    data: data,
    timestamp: 0))

  conn.input.forEachChunk(proc(s: seq[byte]) = discard).ignore

  for i in 0..parseInt(paramStr(1)):
    await writeMultisegment(conn.output, dataMsg)

when isMainModule:
  main().runMain

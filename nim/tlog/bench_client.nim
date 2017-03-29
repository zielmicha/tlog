import reactor, tlog/util, strutils, capnp, tlog/schema

proc main() {.async.} =
  let conn = await connectTcp("localhost", 11211)
  let data = " ".repeat(1024)
  let dataMsg = packPointer(TlogBlock(
    volumeId: 0,
    sequence: 1,
    lba: 0,
    size: 0,
    data: data,
    timestamp: 0))

  conn.input.forEachChunk(proc(s: seq[byte]) = discard).ignore

  while true:
    await writeMultisegment(conn.output, dataMsg)

when isMainModule:
  main().runMain

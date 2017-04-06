import reactor, tlog/util, strutils, capnp, tlog/schema, strutils, os

proc main() {.async.} =
  let conn = await connectTcp("localhost", 11211)
  let data = " ".repeat(1024 * 16)
  var dataMsg = packPointer(TlogBlock(
    volumeId: 2,
    sequence: 1,
    lba: 1,
    size: data.len.uint32,
    data: data,
    timestamp: 111))

  dataMsg.setLen(16472 - 8)

  echo dataMsg.len
  let count = parseInt(paramStr(1))

  proc reader() {.async.} =
    for i in 0..<count:
      let cnt = parseInt((await conn.input.readLine()).strip)
      let data = await conn.input.read(cnt)

  let rd = reader()

  for i in 0..<count:
    await writeMultisegment(conn.output, dataMsg)

  await rd

when isMainModule:
  main().runMain

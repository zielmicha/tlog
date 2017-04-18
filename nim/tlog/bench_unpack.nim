import capnp, tlog/schema, strutils, os

proc main() =
  let data = " ".repeat(1024 * 16)
  var dataMsg = packPointer(TlogBlock(
    volumeId: 2,
    sequence: 1,
    lba: 1,
    size: data.len.uint32,
    data: data,
    timestamp: 111))

  for i in 0..<20000:
    # capnp.newUnpackerFlat 0.148s
    # unpackPointer(0, TlogBlock) ~1s
    let r = capnp.newUnpackerFlat(dataMsg).unpackPointer(0, TlogBlock)

when isMainModule:
  main()

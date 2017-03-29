import reactor, capnp, collections

proc readMultisegment*(input: ByteInput): Future[seq[string]] {.async.} =
  # TODO: move this to capnp.nim
  # TODO(perf): read without copy (unserialize straight from receive buffer)
  let count = (await input.readItem(uint32, littleEndian)).int
  if count > 128 or count < 0:
    raise newException(ValueError, "message too big")

  let lengthsBuf = await input.read((count + 1 + (count mod 2)) * 4)
  var segments: seq[string] = newSeq[string](count + 1)
  var totalLength: int64 = 0

  for i in 0..count:
    let length = unpack(lengthsBuf, i * 4, uint32) * 8
    totalLength += length.int64
    if totalLength > 128 * 1024 * 1024:
      raise newException(ValueError, "message too big")
    let data = await input.read(length.int)
    segments[i] = data

  segments.shallow # pass this seq by reference
  return segments

proc writeMultisegment*(output: ByteOutput, data: string) {.async.} =
  await output.writeItem(uint32(0), littleEndian)
  await output.writeItem(uint32(data.len div 8), littleEndian)
  await output.write(data)

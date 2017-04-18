import reactor, capnp, collections

proc receiveInto*[T](self: Input[T], target: View[T]) {.async.} =
  var offset = 0
  while offset < target.len:
    await self.waitForData
    offset += self.receiveSomeInto(target.slice(offset))

proc readMultisegment*(input: ByteInput): Future[seq[ByteView]] {.async.} =
  # TODO: move this to capnp.nim
  # TODO(perf): read without copy (unserialize straight from receive buffer?)
  let count = (await input.readItem(uint32, littleEndian)).int
  if count > 128 or count < 0:
    raise newException(ValueError, "message too big")

  let lengthsBuf = await input.read((count + 1 + (count mod 2)) * 4)
  var segments = newSeq[ByteView](count + 1)
  var totalLength: int64 = 0

  for i in 0..count:
    let length = unpack(lengthsBuf, i * 4, uint32) * 8
    totalLength += length.int64
    if totalLength > 128 * 1024 * 1024:
      raise newException(ValueError, "message too big")
    segments[i] = ByteView(data: allocShared(length.int), size: length.int)
    #var data = await input.read(length.int) # TODO: use readInto
    #data.stringView.copyTo(segments[i])
    await input.receiveInto(segments[i])

  segments.shallow # pass this seq by reference
  return segments

proc writeMultisegment*(output: ByteOutput, data: string) {.async.} =
  await output.writeItem(uint32(0), littleEndian)
  await output.writeItem(uint32(data.len div 8), littleEndian)
  await output.write(data)

import capnp, capnp/gensupport, collections/iface

# file: ../schema/tlog_schema.capnp

type
  TlogBlock* = ref object
    volumeId*: uint32
    sequence*: uint64
    lba*: uint64
    size*: uint32
    crc32*: uint32
    data*: string
    timestamp*: uint64

  TlogAggregation* = ref object
    name*: string
    size*: uint64
    timestamp*: uint64
    volumeId*: uint32
    blocks*: seq[TlogBlock]
    prev*: string



makeStructCoders(TlogBlock, [
  (volumeId, 0, 0, true),
  (sequence, 8, 0, true),
  (lba, 16, 0, true),
  (size, 4, 0, true),
  (crc32, 24, 0, true),
  (timestamp, 32, 0, true)
  ], [
  (data, 0, PointerFlag.none, true)
  ], [])

makeStructCoders(TlogAggregation, [
  (size, 0, 0, true),
  (timestamp, 8, 0, true),
  (volumeId, 16, 0, true)
  ], [
  (name, 0, PointerFlag.text, true),
  (blocks, 1, PointerFlag.none, true),
  (prev, 2, PointerFlag.none, true)
  ], [])



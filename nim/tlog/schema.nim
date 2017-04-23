import capnp, capnp/gensupport, collections/iface, collections/views

# file: ../schema/tlog_schema.capnp

type
  TlogResponse* = ref object
    status*: int8
    sequences*: seq[uint64]

  TlogBlock* = ref object
    volumeId*: string
    sequence*: uint64
    lba*: uint64
    size*: uint32
    crc32*: uint32
    data*: ByteView
    timestamp*: uint64

  TlogAggregation* = ref object
    name*: string
    size*: uint64
    timestamp*: uint64
    volumeId*: string
    blocks*: seq[TlogBlock]
    prev*: string



makeStructCoders(TlogResponse, [
  (status, 0, 0, true)
  ], [
  (sequences, 0, PointerFlag.none, true)
  ], [])

makeStructCoders(TlogBlock, [
  (sequence, 0, 0, true),
  (lba, 8, 0, true),
  (size, 16, 0, true),
  (crc32, 20, 0, true),
  (timestamp, 24, 0, true)
  ], [
  (volumeId, 0, PointerFlag.text, true),
  (data, 1, PointerFlag.none, true)
  ], [])

makeStructCoders(TlogAggregation, [
  (size, 0, 0, true),
  (timestamp, 8, 0, true)
  ], [
  (name, 0, PointerFlag.text, true),
  (volumeId, 1, PointerFlag.text, true),
  (blocks, 2, PointerFlag.none, true),
  (prev, 3, PointerFlag.none, true)
  ], [])



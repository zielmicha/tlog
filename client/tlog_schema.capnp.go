package client

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type TlogBlock struct{ capnp.Struct }

// TlogBlock_TypeID is the unique identifier for the type TlogBlock.
const TlogBlock_TypeID = 0x8cf178de3c82d431

func NewTlogBlock(s *capnp.Segment) (TlogBlock, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 1})
	return TlogBlock{st}, err
}

func NewRootTlogBlock(s *capnp.Segment) (TlogBlock, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 1})
	return TlogBlock{st}, err
}

func ReadRootTlogBlock(msg *capnp.Message) (TlogBlock, error) {
	root, err := msg.RootPtr()
	return TlogBlock{root.Struct()}, err
}

func (s TlogBlock) String() string {
	str, _ := text.Marshal(0x8cf178de3c82d431, s.Struct)
	return str
}

func (s TlogBlock) VolumeId() uint32 {
	return s.Struct.Uint32(0)
}

func (s TlogBlock) SetVolumeId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s TlogBlock) Sequence() uint32 {
	return s.Struct.Uint32(4)
}

func (s TlogBlock) SetSequence(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s TlogBlock) Lba() uint64 {
	return s.Struct.Uint64(8)
}

func (s TlogBlock) SetLba(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s TlogBlock) Size() uint32 {
	return s.Struct.Uint32(16)
}

func (s TlogBlock) SetSize(v uint32) {
	s.Struct.SetUint32(16, v)
}

func (s TlogBlock) Crc32() uint32 {
	return s.Struct.Uint32(20)
}

func (s TlogBlock) SetCrc32(v uint32) {
	s.Struct.SetUint32(20, v)
}

func (s TlogBlock) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s TlogBlock) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TlogBlock) SetData(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s TlogBlock) Timestamp() uint64 {
	return s.Struct.Uint64(24)
}

func (s TlogBlock) SetTimestamp(v uint64) {
	s.Struct.SetUint64(24, v)
}

// TlogBlock_List is a list of TlogBlock.
type TlogBlock_List struct{ capnp.List }

// NewTlogBlock creates a new list of TlogBlock.
func NewTlogBlock_List(s *capnp.Segment, sz int32) (TlogBlock_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 1}, sz)
	return TlogBlock_List{l}, err
}

func (s TlogBlock_List) At(i int) TlogBlock { return TlogBlock{s.List.Struct(i)} }

func (s TlogBlock_List) Set(i int, v TlogBlock) error { return s.List.SetStruct(i, v.Struct) }

// TlogBlock_Promise is a wrapper for a TlogBlock promised by a client call.
type TlogBlock_Promise struct{ *capnp.Pipeline }

func (p TlogBlock_Promise) Struct() (TlogBlock, error) {
	s, err := p.Pipeline.Struct()
	return TlogBlock{s}, err
}

type TlogAggregation struct{ capnp.Struct }

// TlogAggregation_TypeID is the unique identifier for the type TlogAggregation.
const TlogAggregation_TypeID = 0xe46ab5b4b619e094

func NewTlogAggregation(s *capnp.Segment) (TlogAggregation, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return TlogAggregation{st}, err
}

func NewRootTlogAggregation(s *capnp.Segment) (TlogAggregation, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return TlogAggregation{st}, err
}

func ReadRootTlogAggregation(msg *capnp.Message) (TlogAggregation, error) {
	root, err := msg.RootPtr()
	return TlogAggregation{root.Struct()}, err
}

func (s TlogAggregation) String() string {
	str, _ := text.Marshal(0xe46ab5b4b619e094, s.Struct)
	return str
}

func (s TlogAggregation) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s TlogAggregation) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TlogAggregation) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s TlogAggregation) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s TlogAggregation) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s TlogAggregation) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s TlogAggregation) Timestamp() uint64 {
	return s.Struct.Uint64(8)
}

func (s TlogAggregation) SetTimestamp(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s TlogAggregation) VolumeId() uint32 {
	return s.Struct.Uint32(16)
}

func (s TlogAggregation) SetVolumeId(v uint32) {
	s.Struct.SetUint32(16, v)
}

func (s TlogAggregation) Blocks() (TlogBlock_List, error) {
	p, err := s.Struct.Ptr(1)
	return TlogBlock_List{List: p.List()}, err
}

func (s TlogAggregation) HasBlocks() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s TlogAggregation) SetBlocks(v TlogBlock_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewBlocks sets the blocks field to a newly
// allocated TlogBlock_List, preferring placement in s's segment.
func (s TlogAggregation) NewBlocks(n int32) (TlogBlock_List, error) {
	l, err := NewTlogBlock_List(s.Struct.Segment(), n)
	if err != nil {
		return TlogBlock_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s TlogAggregation) Prev() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s TlogAggregation) HasPrev() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s TlogAggregation) SetPrev(v []byte) error {
	return s.Struct.SetData(2, v)
}

// TlogAggregation_List is a list of TlogAggregation.
type TlogAggregation_List struct{ capnp.List }

// NewTlogAggregation creates a new list of TlogAggregation.
func NewTlogAggregation_List(s *capnp.Segment, sz int32) (TlogAggregation_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	return TlogAggregation_List{l}, err
}

func (s TlogAggregation_List) At(i int) TlogAggregation { return TlogAggregation{s.List.Struct(i)} }

func (s TlogAggregation_List) Set(i int, v TlogAggregation) error {
	return s.List.SetStruct(i, v.Struct)
}

// TlogAggregation_Promise is a wrapper for a TlogAggregation promised by a client call.
type TlogAggregation_Promise struct{ *capnp.Pipeline }

func (p TlogAggregation_Promise) Struct() (TlogAggregation, error) {
	s, err := p.Pipeline.Struct()
	return TlogAggregation{s}, err
}

const schema_f4533cbae6e08506 = "x\xdal\xd2?h\x14A\x18\x05\xf0\xf7\xbe\xd9\xbdK" +
	" \xf1\x9c\xb8`\x04!`\x19PL\xd2\x85\x80\x7f:" +
	"\xad\xf2\xa9`)\x9b\xcdr\x9e\xd9\xbb=\xef6A," +
	"\x05;\xc5&\x85)\x02\x11,\x14\xd2\x08\x17\x0bQH" +
	"a B\x04\x85\x13\xac\xe4\x10\xb1\x0f\xd8X\xc8\xca\x10" +
	"\xf7rB\xba\x997\xdf\xc0\xbc\x1fs~\x82\x17e\xca" +
	"?)\x80\x8e\xfb\xa5|\xaa\xfb`\xee\xdb\xbd\xfdG\xd0" +
	"1zy\xe9a\xef\xe7\x9b\xb9\xeb\xbf\xe0\xb3\x0c\xd8\xfd" +
	"\xef\xf6O\x19\x98\xfa\xfd\x84`\xbe\xda;\xf5\xba\xb3u" +
	"\xe7\x87\x1b5\x03\xa3\xa6\x0c\xcc<\x961\x9eX\x17\xb7" +
	"\\\x93\x9b\xc4\xd9<K\xd2\xea\xadvt[\xe2zx" +
	".\x0a\x9b\x8d\xe6\xec\x8d$\xad^N\xcai\xb44O" +
	"\xeai\xe3\x01\x1e\x01\xbbu\x15\xd0\x8e\xa1n\x0b-\x19" +
	"\xd0\x85\xef\\\xf8\xd6Pw\x85V\x18P\x00\xbbs\x06" +
	"\xd0mC\xdd\x13Z\xe3\x054\x80\xfd0\x09\xe8{C" +
	"\xfd$\xb4\x9e\x1f\xd0\x03\xec\xc7i@w\x0d\xb5+\xa4" +
	"\x1f\xd0\x07\xecg7\xb8g\xa8_\x85\xb6d\x02\x96\x00" +
	"\xfb\xe5\x1a\xa0]C\xed\x09\xf3\x954Y\xae\xc7W\x16" +
	"\x01p\x08\xc2!0o\xc7w\x97\xe3F\x14\x0fd\xe5" +
	"d!\xe40\x84\xc3`\xa5]\xbb\x1f\x17\x07\x13Q+" +
	"\x9a\x99.v\x95\xc50\x0b9\x0a\xe1(\x98g\xb5z" +
	"\xdc\xce\xc2:\xd8,\xee\xf6\x91\xcc\xffH\x97\xaa\xd5V" +
	"\\\x0d\xb3Z\xca\x86\xa3\x1a\xefS\xad\xb9\x0a\xab\x86\xba" +
	"!,\xa4\xd6]\xf6\xd4P\x9f\x0fH=s\xb56\x0c" +
	"us@\xea\xa53}a\xa8\x1d'\xc5\x03\xa9W\xb3" +
	"\x80n\xfe\x83\xf6\xe5\x80jg\xf2\x10\xba\xd2\x08\xeb1" +
	"G \x1c)\xea\xf6\xdf\x7fD\xa7#\x0c/,$i" +
	"\xb4\xd4\xe61p\xde\x90\xc7\x0f\xbf\x1d\xe8\xc2J\xb3\x15" +
	"\xaf\x14P\x7f\x03\x00\x00\xff\xff\xc3\xf9\x92\x86"

func init() {
	schemas.Register(schema_f4533cbae6e08506,
		0x8cf178de3c82d431,
		0xe46ab5b4b619e094)
}

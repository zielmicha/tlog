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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 1})
	return TlogBlock{st}, err
}

func NewRootTlogBlock(s *capnp.Segment) (TlogBlock, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 1})
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

func (s TlogBlock) Sequence() uint64 {
	return s.Struct.Uint64(8)
}

func (s TlogBlock) SetSequence(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s TlogBlock) Lba() uint64 {
	return s.Struct.Uint64(16)
}

func (s TlogBlock) SetLba(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s TlogBlock) Size() uint32 {
	return s.Struct.Uint32(4)
}

func (s TlogBlock) SetSize(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s TlogBlock) Crc32() uint32 {
	return s.Struct.Uint32(24)
}

func (s TlogBlock) SetCrc32(v uint32) {
	s.Struct.SetUint32(24, v)
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
	return s.Struct.Uint64(32)
}

func (s TlogBlock) SetTimestamp(v uint64) {
	s.Struct.SetUint64(32, v)
}

// TlogBlock_List is a list of TlogBlock.
type TlogBlock_List struct{ capnp.List }

// NewTlogBlock creates a new list of TlogBlock.
func NewTlogBlock_List(s *capnp.Segment, sz int32) (TlogBlock_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 40, PointerCount: 1}, sz)
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

const schema_f4533cbae6e08506 = "x\xdal\xd2?h\x13o\x1c\x06\xf0\xe7\xf9\xbewI" +
	"\x0b\xed/\xbf\xb7\x1eXA(8\x16\x14\xd3n\xa5\xe0" +
	"\x9fM\xa7~Up\x94\xeb\xf5\x88\xb1\x97\\L\xaeE" +
	"\x1c\x057\xc5\xa5\x83\x1d\x0a\x15\x1c\x14\xba\x08\xa9\x83(" +
	"t\xb0P\xa1\x82B\x04'\x09\"\xee\x05\x17\x079y" +
	"\x89\x97*t\xfb\xde\xc3sp\xcf\x87;;\xc5\xf3R" +
	"\xf5\x8f\x0b\xa0\x93~)\xaf\xf6\xee\xcd\x7f\xb9s\xf0\x00" +
	":A?/\xdd\xef\x7f\x7f5\x7f\xf5\x07|\x96\x01{" +
	"\xf0\xd5\xfe*\x03\xd5\x9f\x8f\x08\xe6k\xfd\x13/\xbb\xdb" +
	"\xb7\xbe\xb9\xaa\xf9\xabj\xca\xc0\xecC\x99\xe0\xb1\x0dq" +
	"\xe7\xba\\'N\xe7Y\x92\xd6nt\xa2\x9b\x127\xc2" +
	"3Q\xd8j\xb6\xe6\xae%i\xedbRN\xa3\xe5\x05" +
	"RO\x1a\x0f\xf0\x08\xd8\xed\xcb\x80v\x0duGh\xc9" +
	"\x80.|\xe3\xc2\xd7\x86\xba'\xb4\"\x01\x05\xb0\xbb\xa7" +
	"\x00\xdd1\xd4}\xa15\x0ch\x00\xfbn\x1a\xd0\xb7\x86" +
	"\xfaAh\xbdR@\x0f\xb0\xefg\x00\xdd3\xd4\x9e\x90" +
	"~@\x1f\xb0\x1f]q\xdfP?\x0bm\xc9\x0bX\x02" +
	"\xec\xa7+\x80\xf6\x0c\xb5/\xccW\xd3d\xa5\x11_Z" +
	"\x02\xc0\x11\x08G\xc0\xbc\x13\xdf^\x89\x9bQ\xec\xb2Q" +
	"\x08G\xc1r\xb2\x18\x16w\xa5S\xbf\x1b\x17\xe5\xa9\xa8" +
	"\x1d\xcd\xce\x14O\x95\xa50\x0b9\x0e\xe18\x98g\xf5" +
	"F\xdc\xc9\xc2\x06\xd8*\xde\x1d\"\x99\x7f\x91.\xd4j" +
	"\xed\xb8\x16f\xf5\x94MG59\xa4Zw\x13\xd6\x0c" +
	"uSXHm\xb8\xec\xb1\xa1>uR\x1cH=q" +
	"\xb36\x0du\xcbIy\x03\xa9\xe7\xce\xf4\x99\xa1v\x9d" +
	"\x14\x07R/\xe6\x00\xdd\xfa\x03\xed\xcb\x80jw\xfa\x10" +
	"\xba\xd2\x0c\x1b1\xc7 \x1c+\xe6\x0e\xbf\xff\x88MG" +
	"\x18\x9e[L\xd2h\xb9\xc3\xff\xc0\x05C\xfe\x7f\xf8\xdb" +
	"\x81.\xac\xb4\xda\xf1j\x01\xf5;\x00\x00\xff\xff\xc9\x06" +
	"\x92\x89"

func init() {
	schemas.Register(schema_f4533cbae6e08506,
		0x8cf178de3c82d431,
		0xe46ab5b4b619e094)
}

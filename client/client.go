package client

import (
	"bytes"
	"errors"
	"fmt"
	"net"

	"zombiezen.com/go/capnproto2"
)

var (
	ErrInvalidDataLen = errors.New("data length must be 16384 bytes")
)

type Client struct {
	conn net.Conn
}

func New(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Send(volID, seq, crc32 uint32, lba, timestamp uint64, data []byte) error {
	if len(data) != 1024*16 {
		return ErrInvalidDataLen
	}

	b, err := buildCapnp(volID, seq, crc32, lba, timestamp, data)
	if err != nil {
		return err
	}
	_, err = c.conn.Write(b)
	return err
}

func buildCapnp(volID, seq, crc uint32, lba, timestamp uint64, data []byte) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.MultiSegment(nil))
	if err != nil {
		return nil, fmt.Errorf("build capnp:%v", err)
	}

	block, err := NewRootTlogBlock(seg)
	if err != nil {
		return nil, fmt.Errorf("create block:%v", err)
	}

	block.SetVolumeId(volID)
	block.SetSequence(seq)
	block.SetLba(lba)
	block.SetCrc32(crc)
	block.SetTimestamp(timestamp)
	block.SetSize(uint32(len(data)))
	block.SetData(data)

	buf := new(bytes.Buffer)

	err = capnp.NewEncoder(buf).Encode(msg)
	return buf.Bytes(), err
}

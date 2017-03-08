package client

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"

	"zombiezen.com/go/capnproto2"
)

var (
	ErrInvalidDataLen = errors.New("data length must be 16384 bytes")
)

// Client defines a Tlog Client.
// This client is not thread/goroutine safe
type Client struct {
	addr string
	conn *net.TCPConn
	bw   *bufio.Writer
}

// New creates a new tlog client
func New(addr string) (*Client, error) {
	conn, err := createConn(addr)
	if err != nil {
		return nil, err
	}

	return &Client{
		addr: addr,
		conn: conn,
		bw:   bufio.NewWriter(conn),
	}, nil
}

/*
func (c *Client) resetConn() error {
	conn, err := createConn(c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	c.bw = bufio.NewWriter(conn)
}
*/

func createConn(addr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	conn.SetKeepAlive(true)
	return conn, nil
}

// Send sends the transaction tlog to server.
// It returns error in these cases:
// - broken network
// - failed to send all tlog
// in case of errors, client is not in valid state,
// shouldn't be used anymore
func (c *Client) Send(volID uint32, seq uint64, crc32 uint32,
	lba, timestamp uint64, data []byte) error {

	if len(data) != 1024*16 {
		return ErrInvalidDataLen
	}

	b, err := buildCapnp(volID, seq, crc32, lba, timestamp, data)
	if err != nil {
		return err
	}

	_, err = c.sendAll(b)
	return err
}

func (c *Client) sendAll(b []byte) (int, error) {
	nWrite := 0
	for nWrite < len(b) {
		n, err := c.bw.Write(b[nWrite:])
		if err != nil && !isNetTempErr(err) {
			return nWrite, err
		}
		nWrite += n
	}
	return nWrite, c.bw.Flush()
}

func isNetTempErr(err error) bool {
	if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
		return true
	}
	return false
}

func buildCapnp(volID uint32, seq uint64, crc uint32, lba, timestamp uint64, data []byte) ([]byte, error) {
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

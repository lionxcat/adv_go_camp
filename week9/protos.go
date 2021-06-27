package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
)

// copied from "https://github.com/Terry-Mao/goim/blob/master/api/protocol/protocol.go"
// proto struct sizes and offsets
const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

// proto message struct
type protoMsg struct {
	pack_len   int32
	header_len int16
	ver        int16
	op         int32
	seq        int32
	body       []byte
}

// newProtoMsg gen new proto message for tests
func newProtoMsg(v int16, o, s int32, b []byte) *protoMsg {
	p := protoMsg{}
	p.header_len = _rawHeaderSize
	p.ver = v
	p.op = o
	p.seq = s
	body_len := len(b)
	if body_len > 0 {
		p.body = make([]byte, body_len)
		copy(p.body, b)
	}
	p.pack_len = int32(p.header_len) + int32(body_len)
	return &p
}

func (msg *protoMsg) String() string {
	return fmt.Sprintf("ID: %d, v%d, Op: %d, message:%s", msg.seq, msg.ver, msg.op, string(msg.body))
}

// writeSock write encoded proto message to socket
func (msg *protoMsg) writeSock(conn net.Conn) (err error) {
	// deal with head fields
	buf := make([]byte, _rawHeaderSize)
	putInt32(buf[_packOffset:], msg.pack_len)
	putInt16(buf[_headerOffset:], int16(_rawHeaderSize))
	putInt16(buf[_verOffset:], int16(msg.ver))
	putInt32(buf[_opOffset:], msg.op)
	putInt32(buf[_seqOffset:], msg.seq)
	_, err = conn.Write(buf)
	if err != nil {
		return err
	}
	// write body
	if msg.body != nil {
		_, err = conn.Write(msg.body)
	}
	return err
}

// readSock read and decode proto message from socket
func readSock(reader *bufio.Reader) (msg protoMsg, err error) {
	// read fixed length head
	buf := make([]byte, _rawHeaderSize)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return
	}
	msg.pack_len = fromInt32(buf[_packOffset:_headerOffset])
	msg.header_len = fromInt16(buf[_headerOffset:_verOffset])
	msg.ver = fromInt16(buf[_verOffset:_opOffset])
	msg.op = fromInt32(buf[_opOffset:_seqOffset])
	msg.seq = fromInt32(buf[_seqOffset:])
	if msg.header_len != _rawHeaderSize {
		err = errors.New("default server codec header length error")
	}
	// read body bytes
	if body_len := msg.pack_len - int32(msg.header_len); body_len > 0 {
		msg.body = make([]byte, body_len)
		_, err = io.ReadFull(reader, msg.body)
	}
	return
}

// deal with bigendian encodes and decodes
// copied from https://github.com/Terry-Mao/goim/blob/master/pkg/encoding/binary/endian.go

func putInt16(buf []byte, v int16) {
	_ = buf[1]
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
}

func fromInt16(buf []byte) int16 { return int16(buf[1]) | int16(buf[0])<<8 }

func putInt32(buf []byte, v int32) {
	_ = buf[3]
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
}

func fromInt32(buf []byte) int32 {
	return int32(buf[3]) | int32(buf[2])<<8 | int32(buf[1])<<16 | int32(buf[0])<<24
}

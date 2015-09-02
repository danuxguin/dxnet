package network

import (
	// "errors"
	"lanstonetech.com/common"
)

const (
	MAX_PACKAGE_LEN    = 512 * 1024
	MESSAGE_CONTINUE   = 1 //消息处理继续
	MESSAGE_BREAK      = 2 //消息处理中断
	MESSAGE_DISCONNECT = 3 //消息处理完成，断开当前连接
	MESSAGE_ERROR      = 4 //消息处理错误，断开当前连接
)

type MsgHeader struct {
	MsgID     uint32
	MsgVer    uint32
	MsgLen    uint32
	MsgCpsLen uint32
}

type Message struct {
	*MsgHeader
	Data []byte
}

func NewMessage() *Message {
	return new(Message)
}

func ParseHeader(head []byte) *MsgHeader {
	var header MsgHeader
	pos := 0

	header.MsgID = common.ReadUint32(head[pos : pos+4])
	pos += 4
	header.MsgVer = common.ReadUint32(head[pos : pos+4])
	pos += 4
	header.MsgLen = common.ReadUint32(head[pos : pos+4])
	pos += 4
	header.MsgCpsLen = common.ReadUint32(head[pos : pos+4])
	pos += 4

	return &header
}

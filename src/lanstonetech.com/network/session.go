package network

import (
	// "io"
	"lanstonetech.com/common"
	// "lanstonetech.com/packet"
	// "fmt"
	"net"
)

type Session struct {
	conn *net.TCPConn
}

//new session
func NewSession(conn *net.TCPConn) *Session {
	return &Session{
		conn: conn,
	}
}

//read from session
func (this *Session) RecvMSG() (*Message, error) {
	Data := make([]byte, MAX_PACKAGE_LEN)
	_, err := this.conn.Read(Data[0:common.PACKET_HEAD_LEN])
	if err != nil {
		return nil, err
	}

	message := NewMessage()
	header := ParseHeader(Data[0:common.PACKET_HEAD_LEN])

	message.Data = make([]byte, header.MsgLen)
	_, err = this.conn.Read(Data[0:header.MsgLen])
	if err != nil {
		return nil, err
	}

	message.MsgHeader = header
	message.Data = Data[0:header.MsgLen]

	//解压缩
	//解密

	return message, nil
}

//send data to session
func (this *Session) SendMSG(msg *Message) error {
	//加密
	//压缩
	_, err := this.conn.Write(msg.Data)
	return err
}

func (this *Session) Close() {
	defer this.conn.Close()
}

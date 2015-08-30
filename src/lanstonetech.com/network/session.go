package network

import (
	"io"
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
func (this *Session) RecvMSG() ([]byte, error) {
	// var msg *Message
	// msg.Data = make([]byte, MAX_PACKAGE_LEN)
	Data := make([]byte, MAX_PACKAGE_LEN)

	_, err := io.ReadFull(this.conn, Data)
	if err != nil {
		return nil, err
	}

	//解压缩
	//解密

	return Data, nil
}

//send data to session
func (this *Session) SendMSG(msg *Message) error {
	//加密
	//压缩
	_, err := this.conn.Write(msg.Data)
	return err
}

func (this *Session) Close() {

}

package server

import (
	// "fmt"
	"lanstonetech.com/common"
	"lanstonetech.com/common/logger"
	"lanstonetech.com/network"
	"lanstonetech.com/packet"
	"net"
	"time"
)

type server struct {
	*network.Dispatcher
}

var Server server

func init() {
	Server.Dispatcher = network.NewDispatcher()
}

func InitPacketHandler() {
	Server.Dispatcher.AddHandler(1001, TestHandler)
}

func OnRun(cmd []string) {

	defer logger.CatchException()

	logger.Infof("OnRun")

	InitPacketHandler()

	//执行循环
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("123.57.81.38"), 10001, ""})
	if err != nil {
		panic(err)
	}

	for {
		conn, e := listen.AcceptTCP()
		if e != nil {
			logger.Infof("[Main] => AcceptTCP failed! err=%v", e)
			continue
		}

		//协程处理
		go Server.handlerConnection(conn)
	}
}

//处理连接请求
func (this *server) handlerConnection(conn *net.TCPConn) {

	defer logger.CatchException()

	logger.Infof("New connection coming ... IP=%s ", conn.RemoteAddr())

	conn.SetNoDelay(true)                                        //无延迟
	conn.SetKeepAlive(true)                                      //保持激活
	conn.SetReadBuffer(64 * 1024)                                //设置读缓冲区大小
	conn.SetWriteBuffer(64 * 1024)                               //设置写缓冲区大小
	conn.SetReadDeadline(time.Now().Add(30000000 * time.Second)) //设置读超时

	session := network.NewSession(conn)
	defer session.Close()

	for {
		msg, err := session.RecvMSG()
		if err != nil {
			logger.Infof("RecvMsgs IP=%s err=%v", conn.RemoteAddr(), err.Error())
			return
		}

		if len(msg) == 0 {
			continue
		}

		ret := this.ProcessMessages(session, msg)
		if ret == false {
			return
		}
	}
}

func (this *server) ProcessMessages(session *network.Session, msg []byte) bool {
	header := this.ParseHeader(msg[:16])
	message, err := network.NewMessage(msg[16:32])

	if err != nil {
		logger.Infof("Err=%v", err)
		return false
	}

	ret := this.Dispatcher.Handle(header.MsgID, session, message)
	switch ret {
	case 0:
		return true
	case 1:
		return false
	}

	return true
}

func (this *server) ParseHeader(head []byte) packet.MsgHeader {
	var header packet.MsgHeader
	pos := 0

	header.MsgID = common.ReadUint32(head[pos : pos+4])
	pos += 4
	header.MsgVer = common.ReadUint32(head[pos : pos+4])
	pos += 4
	header.MsgLen = common.ReadUint32(head[pos : pos+4])
	pos += 4
	header.MsgCpsLen = common.ReadUint32(head[pos : pos+4])
	pos += 4

	return header
}

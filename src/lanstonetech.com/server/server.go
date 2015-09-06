package server

import (
	// "fmt"
	"lanstonetech.com/common/logger"
	"lanstonetech.com/network"
	"lanstonetech.com/system/config"
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
	logger.Infof("ip = %v port = %v", config.SERVER_IP, config.SERVER_PORT)
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(config.SERVER_IP), config.SERVER_PORT, ""})
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

		ret := this.ProcessMessage(session, msg)
		if ret == false {
			return
		}
	}
}

func (this *server) ProcessMessage(session *network.Session, message *network.Message) bool {
	ret := this.Dispatcher.Handle(session, message)
	switch ret {
	case network.MESSAGE_CONTINUE:
		return true
	case network.MESSAGE_BREAK:
		return true
	case network.MESSAGE_DISCONNECT:
		return false
	case network.MESSAGE_ERROR:
		return false
	}

	return true
}

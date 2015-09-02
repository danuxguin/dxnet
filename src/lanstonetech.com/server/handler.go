package server

import (
	"lanstonetech.com/common/logger"
	"lanstonetech.com/network"
	// "net"
)

func ProcessMessage(session *network.Session, message *network.Message) {

}

func TestHandler(session *network.Session, message *network.Message) int {
	logger.Infof("TestHandler======================")
	session.SendMSG(message)

	return network.MESSAGE_DISCONNECT
}

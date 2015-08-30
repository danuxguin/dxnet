package server

import (
	"lanstonetech.com/network"
	// "net"
)

func ProcessMessage(session *network.Session, message *network.Message) {

}

func TestHandler(session *network.Session, message *network.Message) int {
	session.SendMSG(message)

	return 1
}

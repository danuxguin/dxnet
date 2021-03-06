package network

import (
	"fmt"
	// "lanstonetech.com/common"
	"sync"
)

type PacketHandler func(session *Session, message *Message) int

type Dispatcher struct {
	rwlock     sync.RWMutex
	handlerMap map[uint32]PacketHandler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlerMap: make(map[uint32]PacketHandler),
	}
}

func (this *Dispatcher) AddHandler(id uint32, handler PacketHandler) {
	this.rwlock.Lock()
	defer this.rwlock.Unlock()
	this.handlerMap[id] = handler
}

func (this *Dispatcher) DelHandler(id uint32, handler PacketHandler) {
	this.rwlock.Lock()
	defer this.rwlock.Unlock()
	delete(this.handlerMap, id)
}

func (this *Dispatcher) Handle(session *Session, message *Message) int {
	this.rwlock.RLock()
	defer this.rwlock.RUnlock()

	for k, v := range this.handlerMap {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	h, ok := this.handlerMap[message.MsgHeader.MsgID]
	if ok {
		return h(session, message)
	} else {
		fmt.Println("NOT FOUND")
	}

	return 0
}

package network

import (
	"fmt"
	"sync"
)

type PacketHandler func(session *Session, message *Message) int

type dispatcher struct {
	rwlock     sync.RWMutex
	handlerMap map[uint32]PacketHandler
}

var Dispatcher dispatcher

func NewDispatcher() *dispatcher {
	return &dispatcher{
		handlerMap: make(map[uint32]PacketHandler),
	}
}

func (this *dispatcher) AddHandler(id uint32, handler PacketHandler) {
	this.rwlock.Lock()
	defer this.rwlock.Unlock()
	this.handlerMap[id] = handler
}

func (this *dispatcher) DelHandler(id uint32, handler PacketHandler) {
	this.rwlock.Lock()
	defer this.rwlock.Unlock()
	delete(this.handlerMap, id)
}

func (this *dispatcher) Handle(id uint32, session *Session, message *Message) int {
	this.rwlock.RLock()
	defer this.rwlock.RUnlock()
	h, ok := this.handlerMap[id]
	if ok {
		return h(session, message)
	} else {
		fmt.Println("NOT FOUND")
	}

	return 0
}

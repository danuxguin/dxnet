package network

import (
	"errors"
)

const (
	MAX_PACKAGE_LEN = 512 * 1024
)

type Message struct {
	Data []byte
}

func NewMessage(message []byte) (*Message, error) {
	if len(message) > MAX_PACKAGE_LEN {
		return nil, errors.New("Message is too big")
	}
	msg := new(Message)
	msg.Data = make([]byte, len(message))
	msg.Data = message

	return msg, nil
}

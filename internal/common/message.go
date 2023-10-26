package common

import "github.com/godbus/dbus/v5"

type Message struct {
	Name    string
	Path    string
	Sender  string
	Variant []any
	Headers map[dbus.HeaderField]dbus.Variant
}

func newMessage(message *dbus.Message) Message {
	if message == nil {
		return Message{}
	}
	return Message{
		Name:    message.Type.String(),
		Variant: message.Body,
		Headers: message.Headers,
	}
}

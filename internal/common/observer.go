package common

import "github.com/godbus/dbus/v5"

type Observer interface {
	GetSignals() []Signal
	GetListeners() []Listener
	Initialise(connection *dbus.Conn)
	Refresh(connection *dbus.Conn)

	AttachSignals(func(Signal))
	DetachSignals(func(Signal))

	Broadcast(*Message)
}

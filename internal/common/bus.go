package common

import (
	"github.com/godbus/dbus/v5"
)

type bus interface {
	AddSignals() error
	RemoveSignals()
	Close()
	Listen()
}

type Bus struct {
	bus
	Connection *dbus.Conn
	// Signals    []Signal
	// Listeners  []Listener
	Observers []Observer
}

func (b *Bus) RegisterObserver(observer Observer) {
	b.Observers = append(b.Observers, observer)
}

func (b *Bus) RefreshObservers() {
	for _, observer := range b.Observers {
		observer.Refresh(b.Connection)
	}
}

func (b *Bus) AttachSignals() {
	for _, observer := range b.Observers {
		observer.AttachSignals(func(signal Signal) {
			call := b.Connection.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, signal.String())
			if call.Err != nil {
				panic(call.Err)
			}
		})
	}
}

func (b *Bus) DetachSignals() {
	for _, observer := range b.Observers {
		observer.DetachSignals(func(signal Signal) {
			call := b.Connection.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, signal.String())
			if call.Err != nil {
				panic(call.Err)
			}
		})
	}
}

func (b *Bus) Broadcast(message *Message) {
	for _, observer := range b.Observers {
		observer.Broadcast(message)
	}
}

func (b *Bus) DetachSignalsAndClose() {
	// defer closing the connection, since
	// remove signals could panic
	defer b.Connection.Close()
	b.DetachSignals()
}

func (b *Bus) makeChannel() chan *dbus.Message {
	c := make(chan *dbus.Message, 10)
	b.Connection.Eavesdrop(c)
	return c
}

func (b *Bus) Listen() {
	b.AttachSignals()
	defer b.DetachSignalsAndClose()

	eavesdrop := b.makeChannel()
	defer close(eavesdrop)

	for m := range eavesdrop {
		ch := newMessage(m)
		b.Broadcast(&ch)
	}
}

func NewSessionBus() Bus {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	return Bus{Connection: conn}
}

func NewSystemBus() Bus {
	conn, err := dbus.SystemBus()
	if err != nil {
		panic(err)
	}
	return Bus{Connection: conn}
}

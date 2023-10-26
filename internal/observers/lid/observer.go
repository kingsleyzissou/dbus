package lid

import (
	"fmt"
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

type Lid struct {
	Closed bool `json:"closed"`

	// we don't want these to be serialized
	signals   []common.Signal
	listeners []common.Listener
}

func (l Lid) notify() {
	enabled := fmt.Sprintf("%v", l.Closed)
	common.Notify("lid", enabled)
}

func (l Lid) GetListeners() []common.Listener {
	return l.listeners
}

func (l Lid) GetSignals() []common.Signal {
	return l.signals
}

func (l *Lid) Refresh(connection *dbus.Conn) {
	l.Closed = isClosed(connection, callObjects[closed])
	l.notify()
}

func (l *Lid) Initialise(connection *dbus.Conn) {
	l.signals = signals
	l.listeners = []common.Listener{
		l.isClosedListener,
	}
	l.Refresh(connection)
}

func (l *Lid) AttachSignals(attach func(common.Signal)) {
	for _, signal := range l.signals {
		attach(signal)
	}
}

func (l *Lid) DetachSignals(detach func(common.Signal)) {
	for _, signal := range l.signals {
		detach(signal)
	}
}

func (l *Lid) Broadcast(message *common.Message) {
	for _, listener := range l.listeners {
		listener(message)
	}
}

func AttachToBus(bus *common.Bus) {
	l := Lid{}
	l.Initialise(bus.Connection)
	bus.RegisterObserver(&l)
}

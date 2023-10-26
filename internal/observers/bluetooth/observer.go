package bluetooth

import (
	"fmt"
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

type Bluetooth struct {
	Enabled bool `json:"enabled"`

	// we don't want these to be serialized
	signals   []common.Signal
	listeners []common.Listener
}

func (b Bluetooth) notify() {
	enabled := fmt.Sprintf("%v", b.Enabled)
	common.Notify("bluetooth", enabled)
}

func (b Bluetooth) GetListeners() []common.Listener {
	return b.listeners
}

func (b Bluetooth) GetSignals() []common.Signal {
	return b.signals
}

func (b *Bluetooth) Refresh(connection *dbus.Conn) {
	b.Enabled = isEnabled(connection, callObjects[enabled])
	b.notify()
}

func (b *Bluetooth) Initialise(connection *dbus.Conn) {
	b.signals = signals
	b.listeners = []common.Listener{
		b.isEnabledListener,
	}
	b.Refresh(connection)
}

func (b *Bluetooth) AttachSignals(attach func(common.Signal)) {
	for _, signal := range b.signals {
		attach(signal)
	}
}

func (b *Bluetooth) DetachSignals(detach func(common.Signal)) {
	for _, signal := range b.signals {
		detach(signal)
	}
}

func (b *Bluetooth) Broadcast(message *common.Message) {
	for _, listener := range b.listeners {
		listener(message)
	}
}

func AttachToBus(bus *common.Bus) {
	b := Bluetooth{}
	b.Initialise(bus.Connection)
	bus.RegisterObserver(&b)
}

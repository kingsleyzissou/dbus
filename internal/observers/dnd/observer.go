package dnd

import (
	"fmt"
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

type DoNotDisturb struct {
	Enabled bool `json:"enabled"`

	// we don't want these to be serialized
	signals   []common.Signal
	listeners []common.Listener
}

func (d DoNotDisturb) notify() {
	enabled := fmt.Sprintf("%v", d.Enabled)
	common.Notify("screencast", enabled)
}

func (d DoNotDisturb) GetListeners() []common.Listener {
	return d.listeners
}

func (d DoNotDisturb) GetSignals() []common.Signal {
	return d.signals
}

func (d *DoNotDisturb) Refresh(connection *dbus.Conn) {
	d.Enabled = isEnabled(connection, callObjects[enabled])
	d.notify()
}

func (d *DoNotDisturb) Initialise(connection *dbus.Conn) {
	d.Refresh(connection)
}

func (d *DoNotDisturb) AttachSignals(attach func(common.Signal)) {
	for _, signal := range d.signals {
		attach(signal)
	}
}

func (d *DoNotDisturb) DetachSignals(detach func(common.Signal)) {
	for _, signal := range d.signals {
		detach(signal)
	}
}

func (d *DoNotDisturb) Broadcast(message *common.Message) {
	for _, listener := range d.listeners {
		listener(message)
	}
}

func AttachToBus(bus *common.Bus) {
	d := DoNotDisturb{}
	d.signals = signals
	d.listeners = []common.Listener{
		d.isEnabledListener,
	}

	d.Initialise(bus.Connection)

	bus.RegisterObserver(&d)
}

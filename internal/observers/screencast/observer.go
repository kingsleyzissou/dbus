package screencast

import (
	"fmt"
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

type Screencast struct {
	Sharing bool `json:"sharing"`

	// we don't want these to be serialized
	signals   []common.Signal
	listeners []common.Listener
}

func (s Screencast) notify() {
	sharing := fmt.Sprintf("%v", s.Sharing)
	common.Notify("screencast", sharing)
}

func (s Screencast) GetListeners() []common.Listener {
	return s.listeners
}

func (s Screencast) GetSignals() []common.Signal {
	return s.signals
}

func (s *Screencast) Refresh(connection *dbus.Conn) {
	s.Sharing = isScreenSharing(connection, callObjects[sharing])
	s.notify()
}

func (s *Screencast) Initialise(connection *dbus.Conn) {
	s.signals = signals
	s.listeners = []common.Listener{
		s.isSharingListener,
	}
	s.Refresh(connection)
}

func (s *Screencast) AttachSignals(attach func(common.Signal)) {
	for _, signal := range s.signals {
		attach(signal)
	}
}

func (s *Screencast) DetachSignals(detach func(common.Signal)) {
	for _, signal := range s.signals {
		detach(signal)
	}
}

func (s *Screencast) Broadcast(message *common.Message) {
	for _, listener := range s.listeners {
		listener(message)
	}
}

func AttachToBus(bus *common.Bus) {
	s := Screencast{}
	s.Initialise(bus.Connection)
	bus.RegisterObserver(&s)
}

package power

import (
	"encoding/json"
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

type Power struct {
	Type         string  `json:"type"`
	Percentage   float64 `json:"percentage"`
	Charging     bool    `json:"charging"`
	UntilEmpty   int64   `json:"until_empty,omitempty"`
	UntilFull    int64   `json:"until_full,omitempty"`
	WarningLevel string  `json:"warning_level,omitempty"`

	// we don't want these to be serialized
	signals   []common.Signal
	listeners []common.Listener
}

func (p Power) notify() {
	jsonPower, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	common.Notify("power", string(jsonPower))
}

func (p Power) GetListeners() []common.Listener {
	return p.listeners
}

func (p Power) GetSignals() []common.Signal {
	return p.signals
}

func (p *Power) Refresh(connection *dbus.Conn) {
	p.Type = getKind(connection, callObjects[kind])
	p.Percentage = getLevel(connection, callObjects[percentage])
	p.Charging = isCharging(connection, callObjects[charging])
	p.UntilEmpty = timeUntilEmpty(connection, callObjects[untilEmpty])
	p.UntilFull = timeUntilFull(connection, callObjects[untilFull])
	p.WarningLevel = getWarningLevel(connection, callObjects[warningLevel])
	p.notify()
}

func (p *Power) Initialise(connection *dbus.Conn) {
	p.signals = signals
	p.listeners = []common.Listener{
		p.typeChangeListener,
		p.percentageChangeListener,
		p.chargingChangeListener,
		p.untilEmptyChangeListener,
		p.untilFullChangeListener,
		p.warningLevelChangeListener,
	}
	p.Refresh(connection)
}

func (p *Power) AttachSignals(attach func(common.Signal)) {
	for _, signal := range p.signals {
		attach(signal)
	}
}

func (p *Power) DetachSignals(detach func(common.Signal)) {
	for _, signal := range p.signals {
		detach(signal)
	}
}

func (p *Power) Broadcast(message *common.Message) {
	for _, listener := range p.listeners {
		listener(message)
	}
}

func AttachToBus(bus *common.Bus) {
	l := Power{}
	l.Initialise(bus.Connection)
	bus.RegisterObserver(&l)
}

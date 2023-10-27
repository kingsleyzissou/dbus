package network

import (
	"encoding/json"
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

type Network struct {
	Enabled      bool   `json:"enabled"`
	VPN          bool   `json:"vpn"`
	Connectivity string `json:"connectivity,omitempty"`
	SSID         string `json:"ssid,omitempty"`
	Device       string `json:"device,omitempty"`

	// we don't want these to be serialized
	signals   []common.Signal
	listeners []common.Listener
}

func (n Network) notify() {
	jsonNetwork, err := json.Marshal(n)
	if err != nil {
		panic(err)
	}
	common.Notify("network", string(jsonNetwork))
}

func (n Network) GetListeners() []common.Listener {
	return n.listeners
}

func (n Network) GetSignals() []common.Signal {
	return n.signals
}

func (n *Network) Refresh(connection *dbus.Conn) {
	n.Enabled = isEnabled(connection, callObjects[enabled])

	if n.Enabled {
		n.VPN = isVPNActive(connection, callObjects[vpnActive])
		n.Connectivity = getConnectivity(connection, callObjects[connected])
		n.Device = getDeviceType(connection, callObjects[deviceType])
	}

	n.notify()
}

func (n *Network) Initialise(connection *dbus.Conn) {
	n.signals = signals
	n.listeners = []common.Listener{
		n.isEnabledListener,
		n.stateChangeListener,
		n.ssidChangeListener,
		n.vpnChangeListener,
		n.typeChangeListener,
	}
	n.Refresh(connection)
}

func (n *Network) AttachSignals(attach func(common.Signal)) {
	for _, signal := range n.signals {
		attach(signal)
	}
}

func (n *Network) DetachSignals(detach func(common.Signal)) {
	for _, signal := range n.signals {
		detach(signal)
	}
}

func (n *Network) Broadcast(message *common.Message) {
	for _, listener := range n.listeners {
		listener(message)
	}
}

func AttachToBus(bus *common.Bus) {
	n := Network{}
	n.Initialise(bus.Connection)
	bus.RegisterObserver(&n)
}

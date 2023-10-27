package network

import (
	"kingsley/dbus/internal/common"
)

func (n *Network) isEnabledListener(m *common.Message) {
	variant := parseVariant(common.PropertiesChanged, m.Variant)
	if val, ok := variant[enabled.String()]; ok {
		enabled := val.Value().(bool)
		n.Enabled = enabled
		n.notify()
	}
}

func (n *Network) stateChangeListener(m *common.Message) {
	variant := parseVariant(common.PropertiesChanged, m.Variant)
	if val, ok := variant[connected.String()]; ok {
		n.Connectivity = parseConnectivityState(val)
		n.notify()
	}
}

func (n *Network) ssidChangeListener(m *common.Message) {
	variant := parseVariant(common.PropertiesChanged, m.Variant)
	if val, ok := variant[ssid.String()]; ok {
		n.SSID = val.Value().(string)
		n.notify()
	}
}

func (n *Network) vpnChangeListener(m *common.Message) {
	variant := parseVariant(vpnStateChanged, m.Variant)
	if val, ok := variant[vpnStateChanged.String()]; ok {
		enabled := parseVPNState(val)
		n.VPN = enabled
		n.notify()
	}
}

func (n *Network) typeChangeListener(m *common.Message) {
	variant := parseVariant(common.PropertiesChanged, m.Variant)
	if val, ok := variant[deviceType.String()]; ok {
		n.Device = parseDeviceType(val)
		n.notify()
	}
}

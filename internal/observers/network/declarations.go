package network

import (
	"kingsley/dbus/internal/common"
)

const (
	sender = "org.freedesktop.NetworkManager"
	path   = "/org/freedesktop/NetworkManager"
	member = "org.freedesktop.NetworkManager"

	enabled           common.PropertyKey = "WirelessEnabled"
	connected         common.PropertyKey = "Connectivity"
	deviceType        common.PropertyKey = "PrimaryConnectionType"
	vpnActive         common.PropertyKey = "VpnState"
	vpnStateChanged   common.PropertyKey = "VpnStateChanged"
	activeConnections common.PropertyKey = "ActiveConnections"
	ssid              common.PropertyKey = "Id"
)

var callObjects = common.CallObjectMap{
	enabled: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  enabled,
	},
	connected: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  connected,
	},
	deviceType: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  deviceType,
	},
	ssid: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  ssid,
	},
	vpnActive: {
		Sender:    sender,
		Path:      "/org/freedesktop/NetworkManager/DnsManager",
		Interface: "org.freedesktop.NetworkManager.DnsManager",
		Property:  "Configuration",
	},
}

var signals = []common.Signal{
	{
		"type":   "signal",
		"sender": sender,
		"path":   path,
		"member": common.PropertiesChanged,
	},
	{
		"type":      "signal",
		"interface": "org.freedesktop.NetworkManager.VPN.Connection",
		"member":    "VpnStateChanged",
	},
}

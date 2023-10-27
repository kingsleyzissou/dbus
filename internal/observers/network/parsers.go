package network

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

const (
	NM_CONNECTIVITY_UNKNOWN uint32 = 1
	NM_CONNECTIVITY_NONE    uint32 = 2
	NM_CONNECTIVITY_PORTAL  uint32 = 3
	NM_CONNECTIVITY_LIMITED uint32 = 4
	NM_CONNECTIVITY_FULL    uint32 = 5

	NM_DEVICE_TYPE_ETHERNET string = "802-3-ethernet"
	NM_DEVICE_TYPE_WIFI     string = "802-11-wireless"
	NM_DEVICE_TYPE_UNKNOWN  string = "unknown"
	NM_DEVICE_TYPE_GENERIC  string = "generic"

	IS_CONNECTED    uint32 = 5
	IS_DISCONNECTED uint32 = 7
)

var connectivityMap = map[uint32]string{
	NM_CONNECTIVITY_UNKNOWN: "unknown",
	NM_CONNECTIVITY_NONE:    "none",
	NM_CONNECTIVITY_PORTAL:  "portal",
	NM_CONNECTIVITY_LIMITED: "limited",
	NM_CONNECTIVITY_FULL:    "full",
}

func parseVariant(key common.PropertyKey, variant []interface{}) map[string]dbus.Variant {
	return common.ParseVariant(key.String(), variant)
}

func parseConnectivityState(variant dbus.Variant) string {
	state, ok := connectivityMap[variant.Value().(uint32)]
	if !ok {
		return "unknown"
	}
	return state
}

var deviceTypeMap = map[string]string{
	NM_DEVICE_TYPE_UNKNOWN:  "unknown",
	NM_DEVICE_TYPE_GENERIC:  "generic",
	NM_DEVICE_TYPE_ETHERNET: "ethernet",
	NM_DEVICE_TYPE_WIFI:     "wifi",
}

func parseDeviceType(variant dbus.Variant) string {
	state, ok := deviceTypeMap[variant.Value().(string)]
	if !ok {
		return "unknown"
	}
	return state
}

func parseDNSConfig(variant dbus.Variant) []map[string]dbus.Variant {
	return variant.Value().([]map[string]dbus.Variant)
}

func parseObjectPath(variant dbus.Variant) string {
	if path, ok := variant.Value().(dbus.ObjectPath); ok && path.IsValid() {
		return string(path)
	}
	return ""
}

func parseObjectPaths(variant dbus.Variant) []string {
	connections := variant.Value().([]dbus.ObjectPath)
	str := []string{}
	for _, connection := range connections {
		if connection.IsValid() {
			str = append(str, string(connection))
		}
	}
	return str
}

func parseVPNState(variant dbus.Variant) bool {
	state, ok := variant.Value().(uint32)
	if !ok {
		return false
	}
	if state == IS_DISCONNECTED {
		return false
	}
	return state == IS_CONNECTED
}

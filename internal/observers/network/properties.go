package network

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

func isEnabled(conn *dbus.Conn, c common.CallObject) bool {
	property, err := c.GetProperty(conn)
	if err != nil {
		return false
	}
	return property.Value().(bool)
}

func isVPNActive(conn *dbus.Conn, c common.CallObject) bool {
	property, err := c.GetProperty(conn)
	if err != nil {
		return false
	}
	maps := property.Value().([]map[string]dbus.Variant)
	for _, m := range maps {
		for k, v := range m {
			if k == "vpn" && v.Value().(bool) {
				return true
			}
		}
	}
	return false
}

func getConnectivity(conn *dbus.Conn, c common.CallObject) string {
	property, err := c.GetProperty(conn)
	if err != nil {
		return ""
	}
	return connectivityMap[property.Value().(uint32)]
}

func getDeviceType(conn *dbus.Conn, c common.CallObject) string {
	property, err := c.GetProperty(conn)
	if err != nil {
		return ""
	}
	return property.Value().(string)
}

package power

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

func getKind(conn *dbus.Conn, c common.CallObject) string {
	property, err := c.GetProperty(conn)
	if err != nil {
		return ""
	}
	val, ok := typeMap[property.Value().(uint32)]
	if !ok {
		return "Unknown"
	}
	return val
}

func getLevel(conn *dbus.Conn, c common.CallObject) float64 {
	property, err := c.GetProperty(conn)
	if err != nil {
		return 0
	}
	return property.Value().(float64)
}

func isCharging(conn *dbus.Conn, c common.CallObject) bool {
	property, err := c.GetProperty(conn)
	if err != nil {
		return false
	}
	return property.Value().(bool)
}

func timeUntilEmpty(conn *dbus.Conn, c common.CallObject) int64 {
	property, err := c.GetProperty(conn)
	if err != nil {
		return 0
	}
	return property.Value().(int64)
}

func timeUntilFull(conn *dbus.Conn, c common.CallObject) int64 {
	property, err := c.GetProperty(conn)
	if err != nil {
		return 0
	}
	return property.Value().(int64)
}

func getWarningLevel(conn *dbus.Conn, c common.CallObject) string {
	property, err := c.GetProperty(conn)
	if err != nil {
		return ""
	}
	val, ok := warningLevelMap[property.Value().(uint32)]
	if !ok {
		return "Unknown"
	}
	return val
}

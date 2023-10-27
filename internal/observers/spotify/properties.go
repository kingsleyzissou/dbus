package spotify

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

func isRunning(conn *dbus.Conn, c common.CallObject) bool {
	call := c.CallMethod(conn)
	if call.Err != nil {
		return false
	}
	return call.Body[0].(bool)
}

func isPlaying(conn *dbus.Conn, c common.CallObject) bool {
	property, err := c.GetProperty(conn)
	if err != nil {
		return false
	}
	return parsePlaybackStatus(property)
}

func getMetadata(conn *dbus.Conn, c common.CallObject) Metadata {
	property, err := c.GetProperty(conn)
	if err != nil {
		return Metadata{}
	}
	metadata := parseMetadata(property)
	return *metadata
}

package bluetooth

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

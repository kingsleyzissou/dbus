package screencast

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

func isScreenSharing(conn *dbus.Conn, c common.CallObject) bool {
	call := c.CallMethod(conn)
	if call.Err != nil {
		return false
	}
	return len(call.Body) > 0
}

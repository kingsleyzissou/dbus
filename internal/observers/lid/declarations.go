package lid

import "kingsley/dbus/internal/common"

const (
	sender = "org.freedesktop.UPower"
	path   = "/org/freedesktop/UPower"
	member = "org.freedesktop.UPower"

	closed common.PropertyKey = "LidIsClosed"
)

var callObjects = common.CallObjectMap{
	closed: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  closed,
	},
}

var signals = []common.Signal{
	{
		"type":   "signal",
		"sender": sender,
		"path":   path,
		"member": common.PropertiesChanged,
	},
}

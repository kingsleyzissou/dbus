package bluetooth

import "kingsley/dbus/internal/common"

const (
	sender = "org.bluez"
	path   = "/org/bluez/hci0"
	member = "org.bluez.Adapter1"

	enabled common.PropertyKey = "Powered"
)

var callObjects = common.CallObjectMap{
	enabled: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  enabled,
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

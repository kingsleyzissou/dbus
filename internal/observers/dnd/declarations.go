package dnd

import "kingsley/dbus/internal/common"

const (
	sender = "org.freedesktop.Notifications"
	path   = "/org/freedesktop/Notifications"
	member = "org.dunstproject.cmd0"

	enabled common.PropertyKey = "paused"
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

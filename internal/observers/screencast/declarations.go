package screencast

import "kingsley/dbus/internal/common"

const (
	sender = "org.freedesktop.impl.portal.desktop.hyprland"
	path   = "/org/freedesktop/portal/desktop/session"
	member = "org.freedesktop.DBus.Introspectable.Introspect"

	screencastInterface = "org.freedesktop.portal.ScreenCast"
	screensharingOpened = "OpenPipeWireRemote"

	sessionInterface    = "org.freedesktop.impl.portal.Session"
	screensharingClosed = "Close"

	sharing common.PropertyKey = "Sharing"
)

var callObjects = common.CallObjectMap{
	sharing: {
		Sender:    sender,
		Path:      path,
		Interface: member,
	},
}

var signals = []common.Signal{
	{
		"eavesdrop": "true",
		"type":      "method_call",
		"interface": screencastInterface,
		"member":    screensharingOpened,
	},
	{
		"eavesdrop": "true",
		"type":      "method_call",
		"interface": sessionInterface,
		"member":    screensharingClosed,
	},
}

package spotify

import (
	"kingsley/dbus/internal/common"
)

const (
	sender = "org.mpris.MediaPlayer2.spotify"
	path   = "/org/mpris/MediaPlayer2"
	member = "org.mpris.MediaPlayer2.Player"

	running  common.PropertyKey = "Running"
	playing  common.PropertyKey = "PlaybackStatus"
	metadata common.PropertyKey = "Metadata"
)

var callObjects = common.CallObjectMap{
	running: {
		Sender:    common.FreedesktopSender,
		Path:      common.FreedesktopPath,
		Interface: common.FreedesktopInterface,
		Property:  common.NameHasOwner,
		OtherArgs: []interface{}{sender},
	},
	playing: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  playing,
	},
	metadata: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  metadata,
	},
}

// Signals
var signals = []common.Signal{
	{
		"type":   "signal",
		"sender": sender,
		"path":   path,
		"member": common.PropertiesChanged,
	},
	{
		"type":   "signal",
		"sender": common.FreedesktopSender,
		"path":   common.FreedesktopPath,
		"member": common.NameOwnerChanged,
		"arg0":   sender,
	},
}

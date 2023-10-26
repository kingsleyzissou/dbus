package power

import "kingsley/dbus/internal/common"

const (
	sender = "org.freedesktop.UPower"
	path   = "/org/freedesktop/UPower/devices/battery_BAT0"
	member = "org.freedesktop.UPower.Device"

	kind         common.PropertyKey = "Type"
	percentage   common.PropertyKey = "Percentage"
	charging     common.PropertyKey = "PowerSupply"
	untilEmpty   common.PropertyKey = "TimeToEmpty"
	untilFull    common.PropertyKey = "TimeToFull"
	warningLevel common.PropertyKey = "WarningLevel"
)

var callObjects = common.CallObjectMap{
	kind: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  kind,
	},
	percentage: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  percentage,
	},
	charging: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  charging,
	},
	untilEmpty: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  untilEmpty,
	},
	untilFull: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  untilFull,
	},
	warningLevel: {
		Sender:    sender,
		Path:      path,
		Interface: member,
		Property:  warningLevel,
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

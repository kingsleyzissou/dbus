package power

import (
	"github.com/godbus/dbus/v5"
)

var typeMap = map[uint32]string{
	0: "Unknown",
	1: "Line Power",
	2: "Battery",
	3: "Ups",
	4: "Monitor",
	5: "Mouse",
	6: "Keyboard",
	7: "Pda",
	8: "Phone",
}

func parseType(variant dbus.Variant) string {
	kind, ok := typeMap[variant.Value().(uint32)]
	if !ok {
		return "Unknown"
	}
	return kind
}

var warningLevelMap = map[uint32]string{
	0: "Unknown",
	1: "None",
	3: "Low",
	4: "Critical",
	5: "Action",
}

func parseWarningLevel(variant dbus.Variant) string {
	level, ok := warningLevelMap[variant.Value().(uint32)]
	if !ok {
		return "Unknown"
	}
	return level
}

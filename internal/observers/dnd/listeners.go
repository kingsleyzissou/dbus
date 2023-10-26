package dnd

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

func parseVariant(variant []interface{}) map[string]dbus.Variant {
	return common.ParseVariant("PropertiesChanged", variant)
}

func (d *DoNotDisturb) isEnabledListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[enabled.String()]; ok {
		d.Enabled = val.Value().(bool)
		d.notify()
	}
}

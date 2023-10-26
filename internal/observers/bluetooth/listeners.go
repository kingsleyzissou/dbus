package bluetooth

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

func parseVariant(variant []interface{}) map[string]dbus.Variant {
	return common.ParseVariant(common.PropertiesChanged, variant)
}

func (b *Bluetooth) isEnabledListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[enabled.String()]; ok {
		b.Enabled = val.Value().(bool)
		b.notify()
	}
}

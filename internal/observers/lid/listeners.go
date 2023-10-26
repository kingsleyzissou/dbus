package lid

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

func parseVariant(variant []interface{}) map[string]dbus.Variant {
	return common.ParseVariant(common.PropertiesChanged, variant)
}

func (l *Lid) isClosedListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[closed.String()]; ok {
		l.Closed = val.Value().(bool)
		l.notify()
	}
}

package power

import (
	"kingsley/dbus/internal/common"

	"github.com/godbus/dbus/v5"
)

func parseVariant(variant []interface{}) map[string]dbus.Variant {
	return common.ParseVariant(common.PropertiesChanged, variant)
}

func (p *Power) typeChangeListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[kind.String()]; ok {
		kind := parseType(val)
		p.Type = kind
		p.notify()
	}
}

func (p *Power) percentageChangeListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[percentage.String()]; ok {
		percentage := val.Value().(float64)
		p.Percentage = percentage
		p.notify()
	}
}

func (p *Power) chargingChangeListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[charging.String()]; ok {
		charging := val.Value().(bool)
		p.Charging = charging
		p.notify()
	}
}

func (p *Power) untilEmptyChangeListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[untilEmpty.String()]; ok {
		untilEmpty := val.Value().(int64)
		p.UntilEmpty = untilEmpty
		p.notify()
	}
}

func (p *Power) untilFullChangeListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[untilFull.String()]; ok {
		untilFull := val.Value().(int64)
		p.UntilFull = untilFull
		p.notify()
	}
}

func (p *Power) warningLevelChangeListener(m *common.Message) {
	variant := parseVariant(m.Variant)
	if val, ok := variant[warningLevel.String()]; ok {
		warningLevel := parseWarningLevel(val)
		p.WarningLevel = warningLevel
		p.notify()
	}
}

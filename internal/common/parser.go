package common

import "github.com/godbus/dbus/v5"

func ParseVariant(defaultKey string, variant []any) map[string]dbus.Variant {
	if len(variant) == 0 {
		return nil
	}

	if variant, ok := variant[0].(uint32); ok {
		sub := map[string]dbus.Variant{
			defaultKey: dbus.MakeVariant(variant),
		}
		return sub
	}

	if val, ok := variant[1].(map[string]dbus.Variant); ok {
		return val
	}

	return nil
}

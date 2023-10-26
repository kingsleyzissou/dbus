package common

import (
	"fmt"
	"strings"

	"github.com/godbus/dbus/v5"
)

type Signal map[string]string

func (m Signal) String() string {
	s := []string{}
	for k, v := range m {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(s, ",")
}

type CallObject struct {
	Sender    string
	Path      dbus.ObjectPath
	Interface string
	Property  PropertyKey
	OtherArgs []interface{}
}

type PropertyKey string

func (p PropertyKey) String() string {
	return string(p)
}

type CallObjectMap map[PropertyKey]CallObject

func (c CallObject) CallMethod(conn *dbus.Conn) *dbus.Call {
	obj := conn.Object(c.Sender, c.Path)
	call := obj.Call(fmt.Sprintf("%s.%s", c.Interface, c.Property), 0, c.OtherArgs...)
	return call
}

func (c CallObject) GetProperty(conn *dbus.Conn) (dbus.Variant, error) {
	obj := conn.Object(c.Sender, c.Path)
	property, err := obj.GetProperty(fmt.Sprintf("%s.%s", c.Interface, c.Property))
	return property, err
}

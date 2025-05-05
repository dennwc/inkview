package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"

import (
	"errors"
	"fmt"
)

// HwAddress returns device MAC address.
func HwAddress() string {
	return C.GoString(C.GetHwAddress())
}

// Connections returns all available network connections.
// Name can be used as an argument to Connect.
func Connections() []string {
	list := C.EnumConnections()
	return strArr(list)
}

func WirelessNetworks() []string {
	list := C.EnumWirelessNetworks()
	return strArr(list)
}

type NetError struct {
	Code int
	Text string
}

func (e NetError) Error() string {
	if e.Text != "" {
		return e.Text
	}
	return fmt.Sprintf("unknown net error: %d", e.Code)
}

func netError(e C.int) error {
	if e == 0 {
		return nil
	}
	str := C.GoString(C.NetError(e))
	return NetError{Code: int(e), Text: str}
}

func Connect(name string) error {
	cname, free := cString(name)
	defer free()
	e := C.NetConnect(cname)
	return netError(e)
}

func Disconnect() error {
	e := C.NetDisconnect()
	return netError(e)
}

func OpenNetworkInfo() {
	C.OpenNetworkInfo()
}

var (
	ErrNoConnections = errors.New("no connections available")
)

// KeepNetwork will connect a default network interface on the device and will keep it enabled.
// Returned function can be called to disconnect an interface.
func KeepNetwork() (func(), error) {
	conns := Connections()
	if len(conns) == 0 {
		return nil, ErrNoConnections
	}
	var last error
	for _, c := range conns {
		last = Connect(c)
		if last == nil {
			return func() {
				_ = Disconnect()
			}, nil
		}
	}
	return nil, last
}

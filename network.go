package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"

import (
	"fmt"
)

func HwAddress() string {
	return C.GoString(C.GetHwAddress())
}

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
	e := C.NetConnect(C.CString(name))
	return netError(e)
}

func Disconnect() error {
	e := C.NetDisconnect()
	return netError(e)
}

func OpenNetworkInfo() {
	C.OpenNetworkInfo()
}

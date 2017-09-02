package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"

import (
	"time"
	"unsafe"
)

func cbool(v bool) C.int {
	if v {
		return 1
	}
	return 0
}

func incPtr(ptr unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(C.size_t(0)))
}

func strArr(list **C.char) (out []string) {
	if list == nil {
		return
	}
	for list != nil {
		s := C.GoString(*list)
		if len(s) == 0 {
			break
		}
		out = append(out, s)
		list = (**C.char)(incPtr(unsafe.Pointer(list)))
	}
	return
}

func SetSleepMode(on bool) {
	C.iv_sleepmode(cbool(on))
}

func SleepMode() bool {
	return C.GetSleepmode() != 0
}

func BatteryPower() int {
	return int(C.GetBatteryPower())
}

func Temperature() int {
	return int(C.GetTemperature())
}

func IsPressed(key Key) bool {
	return C.IsKeyPressed(C.int(key)) != 0
}

func IsCharging() bool {
	return C.IsCharging() != 0
}

func IsUSBconnected() bool {
	return C.IsUSBconnected() != 0
}

func IsSDinserted() bool {
	return C.IsSDinserted() != 0
}

func DeviceModel() string {
	return C.GoString(C.GetDeviceModel())
}

func HardwareType() string {
	return C.GoString(C.GetHardwareType())
}

func SoftwareVersion() string {
	return C.GoString(C.GetSoftwareVersion())
}

func SerialNumber() string {
	return C.GoString(C.GetSerialNumber())
}

func DeviceKey() string {
	return C.GoString(C.GetDeviceKey())
}

func Sleep(dt time.Duration, deep bool) {
	ms := int(dt / time.Millisecond)
	if ms == 0 {
		ms = 1
	}
	_ = C.GoSleep(C.int(ms), cbool(deep))
}

func SetAutoPowerOff(on bool) {
	C.SetAutoPowerOff(cbool(on))
}

func PowerOff() {
	C.PowerOff()
}

func OpenMainMenu() {
	C.OpenMainMenu()
}

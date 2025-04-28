package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"

import (
	"bufio"
	"os"
	"strings"
	"time"
	"unsafe"
)

var defaultKeyboardLang = "en"

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

func GetConfig() (map[string]string, error) {
	file, err := os.Open(C.GLOBALCONFIGFILE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	configs := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if parts := strings.SplitN(line, "=", 2); len(parts) == 2 {
			configs[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return configs, nil
}

// open the book in the default reader. If the .app file, then run the application
func OpenBook(path string) {
	cPath, free := cString(path)
	defer free()
	C.OpenBook(cPath, (*C.char)(nil), C.int(0))
}

func PageSnapshot() {
	C.PageSnapshot()
}

func cString(str string) (*C.char, func()) {
	cstr := C.CString(str)
	return cstr, func() {
		C.free(unsafe.Pointer(cstr))
	}
}

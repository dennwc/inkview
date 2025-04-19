package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"

import (
	"fmt"
	"image"
	"time"
	"unsafe"
)

var DefaultDelay = time.Second

type Icon int

const (
	Info     = Icon(C.ICON_INFORMATION)
	Question = Icon(C.ICON_QUESTION)
	Warning  = Icon(C.ICON_WARNING)
	Error    = Icon(C.ICON_ERROR)
)

func SetMessageDelay(time time.Duration) {
	DefaultDelay = time
}

func Message(icon Icon, title, text string, dt time.Duration) {
	if dt == 0 {
		dt = DefaultDelay
	}
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	ctxt := C.CString(text)
	defer C.free(unsafe.Pointer(ctxt))
	C.Message(C.int(icon), ctitle, ctxt, C.int(dt/time.Millisecond))
}

func messagef(icon Icon, title, def, format string, args ...interface{}) {
	if title == "" {
		title = def
	}
	Message(icon, title, fmt.Sprintf(format, args...), 0)
}

func Infof(title, format string, args ...interface{}) {
	messagef(Info, title, "Info", format, args...)
}

func Questionf(title, format string, args ...interface{}) {
	messagef(Question, title, "Question", format, args...)
}

func Warningf(title, format string, args ...interface{}) {
	messagef(Warning, title, "Warning", format, args...)
}

func Errorf(title, format string, args ...interface{}) {
	messagef(Error, title, "Error", format, args...)
}

func ShowHourglass() {
	C.ShowHourglass()
}

func ShowHourglassAt(p image.Point) {
	C.ShowHourglassAt(C.int(p.X), C.int(p.Y))
}

func HideHourglass() {
	C.HideHourglass()
}

func DisableExitHourglass() {
	C.DisableExitHourglass()
}

func DrawTopPanel() {
	emptyStr := C.CString("")
	defer C.free(unsafe.Pointer(emptyStr))
	C.DrawPanel(nil, emptyStr, emptyStr, -1)
}

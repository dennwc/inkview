package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview

extern void c_rotate_handler(int direction);
extern void c_dialog_handler(int button);

*/
import "C"

import (
	"fmt"
	"image"
	"time"
	"unsafe"
)

var DefaultDelay = time.Second

type RotateBoxHandler func(Orientation)

var userRotateBoxHandler RotateBoxHandler

// return 1 for left button, 2 for right button. 1 for progressbar cancel button
type DialogHandler func(button int)

var userDialogHandler DialogHandler

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

func OpenRotateBox() {
	var rotateHandler C.iv_rotatehandler
	rotateHandler = (C.iv_rotatehandler)(C.c_rotate_handler)
	C.OpenRotateBox(rotateHandler)
}

func SetRotateBoxHandler(handler RotateBoxHandler) {
	userRotateBoxHandler = handler
}

//export goRotateHandler
func goRotateHandler(d C.int) {
	userRotateBoxHandler(Orientation(d))
}

func Dialog(icon Icon, title, text, button1, button2 string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbutton1 := C.CString(button1)
	defer C.free(unsafe.Pointer(cbutton1))
	cbutton2 := C.CString(button2)
	defer C.free(unsafe.Pointer(cbutton2))

	var dialogHandler C.iv_dialoghandler
	dialogHandler = (C.iv_dialoghandler)(C.c_dialog_handler)

	C.Dialog(C.int(icon), ctitle, ctext, cbutton1, cbutton2, dialogHandler)
}

//export goDialogHandler
func goDialogHandler(button C.int) {
	userDialogHandler(int(button))
}

// Use dialog handler for callback
func OpenProgressbar(icon Icon, title, text string, percent int) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	var dialogHandler C.iv_dialoghandler
	dialogHandler = (C.iv_dialoghandler)(C.c_dialog_handler)
	C.OpenProgressbar(C.int(icon), ctitle, ctext, C.int(percent), dialogHandler)
}

func UpdateProgressbar(text string, percent int) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.UpdateProgressbar(ctext, C.int(percent))
}

func CloseProgressbar() {
	C.CloseProgressbar()
}

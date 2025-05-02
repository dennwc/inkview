package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview

extern void c_rotate_handler(int direction);
extern void c_dialog_handler(int button);
extern void c_timeedit_handler(long time);

*/
import "C"

import (
	"fmt"
	"image"
	"time"
)

var DefaultDelay = time.Second

type RotateBoxHandler func(Orientation)

var userRotateBoxHandler RotateBoxHandler

// return 1 for left button, 2 for right button. 1 for progressbar cancel button
type DialogHandler func(button int)

var userDialogHandler DialogHandler

type TimeEditHandler func(time time.Time)

var userTimeEditHandler TimeEditHandler

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
	ctitle, free := cString(title)
	defer free()
	ctxt, free2 := cString(text)
	defer free2()
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
	emptyStr, free := cString("")
	defer free()
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
	ctitle, free1 := cString(title)
	defer free1()
	ctext, free2 := cString(text)
	defer free2()
	cbutton1, free3 := cString(button1)
	defer free3()
	cbutton2, free4 := cString(button2)
	defer free4()

	var dialogHandler C.iv_dialoghandler
	dialogHandler = (C.iv_dialoghandler)(C.c_dialog_handler)

	C.Dialog(C.int(icon), ctitle, ctext, cbutton1, cbutton2, dialogHandler)
}

func SetDialogHandler(handler DialogHandler) {
	userDialogHandler = handler
}

//export goDialogHandler
func goDialogHandler(button C.int) {
	userDialogHandler(int(button))
}

// Use dialog handler for callback
func OpenProgressbar(icon Icon, title, text string, percent int) {
	ctitle, free := cString(title)
	defer free()
	ctext, free2 := cString(text)
	defer free2()
	var dialogHandler C.iv_dialoghandler
	dialogHandler = (C.iv_dialoghandler)(C.c_dialog_handler)
	C.OpenProgressbar(C.int(icon), ctitle, ctext, C.int(percent), dialogHandler)
}

func UpdateProgressbar(text string, percent int) {
	ctext, free := cString(text)
	defer free()
	C.UpdateProgressbar(ctext, C.int(percent))
}

func CloseProgressbar() {
	C.CloseProgressbar()
}

func SetTimeEditHandler(handler TimeEditHandler) {
	userTimeEditHandler = handler
}

//export goTimeEditHandler
func goTimeEditHandler(t C.long) {
	userTimeEditHandler(time.Unix(int64(t), 0))
}

func OpenTimeEdit(title string, p image.Point, initime time.Time) {
	ctitle, free := cString(title)
	defer free()
	var timeEditHandler C.iv_timeedithandler
	timeEditHandler = (C.iv_timeedithandler)(C.c_timeedit_handler)
	C.OpenTimeEdit(ctitle, C.int(p.X), C.int(p.Y), C.long(initime.Unix()), timeEditHandler)
}

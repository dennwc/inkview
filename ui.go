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
	// void Message(int icon, const char *title, const char *text, int timeout);
	C.Message(C.int(icon), C.CString(title), C.CString(text), C.int(dt/time.Millisecond))
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
	// Taken from the original sudoku.app, a pattern was identified that in applications this code is responsible for rendering the top bar.
	C.DrawPanel(nil, C.CString(""), C.CString(""), -1)
}

package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"

import (
	"image"
)

func ScreenSize() image.Point {
	return image.Point{
		X: int(C.ScreenWidth()),
		Y: int(C.ScreenHeight()),
	}
}

func Screen() image.Rectangle {
	return image.Rectangle{
		Max: ScreenSize(),
	}
}

// Repaint puts Draw event into app's events queue. Eventually Draw method will be called on app object.
//
// Usage: Call Repaint to make app (eventually) redraw itself on the screen.
func Repaint() {
	C.Repaint()
}

// FullUpdate sends content of the whole screen buffer to display driver. Display depth is set to 2 bpp (usually) or 4
// bpp if necessary. Function isn't synchronous i.e. it returns faster, than display is redrawn.
// Update is performed for active app (task) only, if display isn't locked and NO_DISPLAY flag in
// ivstate.uiflags isn't set.
//
// Usage: Tradeoff between quality and speed. Recommended for text and common UI elements. Not
// recommended if quality of picture (image) is required, in such case use FullUpdateHQ().
func FullUpdate() {
	C.FullUpdate()
}

// SoftUpdate is an alternative to FullUpdate. It's effect is (almost) PartialUpdate for the whole screen.
func SoftUpdate() {
	C.SoftUpdate()
}

// PartialUpdate sends content of the given rectangle in screen buffer to display driver. Function is smart and tries to
// perform the most suitable update possible: black and white update is performed if all pixels in given rectangle
// are black and white. Otherwise grayscale update is performed. If whole screen is specified, then grayscale update is performed.
func PartialUpdate(r image.Rectangle) {
	sz := r.Size()
	C.PartialUpdate(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y))
}

func SetOrientation(orientation Orientation) {
	C.SetOrientation(C.int(orientation))
}

// Original pb apps prefer to use setDefaultOrientation on init action (It's an undocumented function, found by reverse engineering)
func SetDefaultOrientation(orientation Orientation) {
	C.SetOrientation(C.int(orientation))
}

func GetOrientation() Orientation {
	return Orientation(C.GetOrientation())
}

func SetGlobalOrientation(orientation Orientation) {
	C.SetGlobalOrientation(C.int(orientation))
}

func GetGlobalOrientation() Orientation {
	return Orientation(C.GetOrientation())
}

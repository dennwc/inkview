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

func FullUpdate() {
	C.FullUpdate()
}

func SoftUpdate() {
	C.SoftUpdate()
}

func PartialUpdate(r image.Rectangle) {
	sz := r.Size()
	C.PartialUpdate(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y))
}

func PartialUpdateBW(r image.Rectangle) {
	sz := r.Size()
	C.PartialUpdateBW(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y))
}

func DynamicUpdateBW(r image.Rectangle) {
	sz := r.Size()
	C.DynamicUpdateBW(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y))
}

func FineUpdate() {
	C.FineUpdate()
}

func DynamicUpdate() {
	C.DynamicUpdate()
}

func FineUpdateSupported() bool {
	return C.FineUpdateSupported() != 0
}

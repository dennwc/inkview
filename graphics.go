package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"

import (
	"image"
	"image/color"
)

func Pad(r image.Rectangle, n int) image.Rectangle {
	dp := image.Pt(n, n)
	r.Min = r.Min.Add(dp)
	r.Max = r.Max.Sub(dp)
	return r
}

var (
	Black     = color.Black
	White     = color.White
	DarkGray  = color.Gray{0x55}
	LightGray = color.Gray{0xaa}
)

func colorToInt(cl color.Color) int {
	r, g, b, _ := cl.RGBA()
	// 0x00RRGGBB
	return (int(b>>8) & 0xff) + int((g>>8)&0xff)<<8 + int((r>>8)&0xff)<<16
}

// ClearScreen fills current canvas with white color.
func ClearScreen() {
	C.ClearScreen()
}

func SetClip(r image.Rectangle) {
	sz := r.Size()
	C.SetClip(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y))
}

func DrawPixel(p image.Point, cl color.Color) {
	C.DrawPixel(C.int(p.X), C.int(p.Y), C.int(colorToInt(cl)))
}

func DrawLine(p1, p2 image.Point, cl color.Color) {
	C.DrawLine(C.int(p1.X), C.int(p1.Y), C.int(p2.X), C.int(p2.Y), C.int(colorToInt(cl)))
}

func DrawRect(r image.Rectangle, cl color.Color) {
	sz := r.Size()
	C.DrawRect(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y), C.int(colorToInt(cl)))
}

func FillArea(r image.Rectangle, cl color.Color) {
	sz := r.Size()
	C.FillArea(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y), C.int(colorToInt(cl)))
}

func InvertArea(r image.Rectangle) {
	sz := r.Size()
	C.InvertArea(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y))
}

func InvertAreaBW(r image.Rectangle) {
	sz := r.Size()
	C.InvertAreaBW(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y))
}

func DimArea(r image.Rectangle, cl color.Color) {
	sz := r.Size()
	C.DimArea(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y), C.int(colorToInt(cl)))
}

func DrawSelection(r image.Rectangle, cl color.Color) {
	sz := r.Size()
	C.DrawSelection(C.int(r.Min.X), C.int(r.Min.Y), C.int(sz.X), C.int(sz.Y), C.int(colorToInt(cl)))
}

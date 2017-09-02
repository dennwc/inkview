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

const (
	DefaultFont           = string(C.DEFAULTFONT)
	DefaultFontBold       = string(C.DEFAULTFONTB)
	DefaultFontItalic     = string(C.DEFAULTFONTI)
	DefaultFontBoldItalic = string(C.DEFAULTFONTBI)
	DefaultFontMono       = string(C.DEFAULTFONTM)
)

func OpenFont(name string, size int, aa bool) *Font {
	p := C.OpenFont(C.CString(name), C.int(size), cbool(aa))
	if p == nil {
		return nil
	}
	return &Font{p: p}
}

type Font struct {
	p *C.ifont
}

func (f *Font) SetActive(cl color.Color) {
	if f != nil && f.p != nil {
		C.SetFont(f.p, C.int(colorToInt(cl)))
	}
}
func (f *Font) Close() {
	if f == nil || f.p == nil {
		return
	}
	C.CloseFont(f.p)
	f.p = nil
}

func DrawString(p image.Point, s string) {
	C.DrawString(C.int(p.X), C.int(p.Y), C.CString(s))
}

func DrawStringR(p image.Point, s string) {
	C.DrawStringR(C.int(p.X), C.int(p.Y), C.CString(s))
}

func CharWidth(c rune) int {
	return int(C.CharWidth(C.ushort(c)))
}

func StringWidth(s string) int {
	return int(C.StringWidth(C.CString(s)))
}

func SetTextStrength(n int) {
	C.SetTextStrength(C.int(n))
}

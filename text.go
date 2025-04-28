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
	"image/color"
	"unsafe"
)

const (
	DefaultFont           = string(C.DEFAULTFONT)
	DefaultFontBold       = string(C.DEFAULTFONTB)
	DefaultFontItalic     = string(C.DEFAULTFONTI)
	DefaultFontBoldItalic = string(C.DEFAULTFONTBI)
	DefaultFontMono       = string(C.DEFAULTFONTM)
)

func OpenFont(name string, size int, aa bool) *Font {
	cname, free := cString(name)
	defer free()
	p := C.OpenFont(cname, C.int(size), cbool(aa))
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
	cs, free := cString(s)
	defer free()
	C.DrawString(C.int(p.X), C.int(p.Y), cs)
}

func DrawStringR(p image.Point, s string) {
	cs, free := cString(s)
	defer free()
	C.DrawStringR(C.int(p.X), C.int(p.Y), cs)
}

func CharWidth(c rune) int {
	return int(C.CharWidth(C.ushort(c)))
}

func StringWidth(s string) int {
	cs, free := cString(s)
	defer free()
	return int(C.StringWidth(cs))
}

func SetTextStrength(n int) {
	C.SetTextStrength(C.int(n))
}

func GetCurrentLang() string {
	configs, err := GetConfig()
	if err == nil {
		lang, ok := configs["language"]
		if ok {
			return fmt.Sprintf("%v", lang)
		}
	}
	return "en"
}

// Probably changes the language the app should run in, translations depend on it
func LoadLanguage(lang string) {
	cLang, free := cString(lang)
	defer free()
	C.LoadLanguage(cLang)
}

// Add translation text that will later be used in getLangText
func AddTranslation(label, trans string) {
	cLabel, free := cString(label)
	defer free()
	cTrans, free2 := cString(trans)
	defer free2()
	C.AddTranslation(cLabel, cTrans)
}

// Get text with translation, translation variables can be found only in original pocketbook apps
func GetLangText(s string) string {
	cS, free := cString(s)
	defer free()
	cText := C.GetLangText(cS)
	defer C.free(unsafe.Pointer(cText))
	return C.GoString(cText)
}

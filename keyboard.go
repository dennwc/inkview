package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview

extern void c_keyboard_handler(char *);
*/
import "C"
import (
	"unsafe"
)

//export goKeyboardHandler
func goKeyboardHandler(text *C.char) {

	userKeyboardHandler(C.GoString(text))
	keyboardBuffer = []byte{}
}

var keyboardBuffer []byte

type KeyboardHandler func(string)

var userKeyboardHandler KeyboardHandler

func SetKeyboardHandler(kh KeyboardHandler) {
	userKeyboardHandler = kh
}

// Open default keyboard
func OpenKeyboard(placeholder string, buflen int) {

	if buflen <= 0 {
		buflen = 1024
	}

	keyboardBuffer := make([]byte, buflen)

	ctitle, free := cString(placeholder)
	defer free()

	cbuffer := (*C.char)(unsafe.Pointer(&keyboardBuffer[0]))

	var chandler C.iv_keyboardhandler
	chandler = (C.iv_keyboardhandler)(C.c_keyboard_handler)

	C.OpenKeyboard(ctitle, cbuffer, C.int(buflen), C.int(0), chandler)
}

// Open keyboard from .kbd file
func OpenCustomKeyboard(keyboardFileName, placeholder string, buflen int) {

	if buflen <= 0 {
		buflen = 1024
	}

	keyboardBuffer := make([]byte, buflen)

	ctitle, free := cString(placeholder)
	defer free()

	cbuffer := (*C.char)(unsafe.Pointer(&keyboardBuffer[0]))

	var chandler C.iv_keyboardhandler
	chandler = (C.iv_keyboardhandler)(C.c_keyboard_handler)

	cfileName, free2 := cString(keyboardFileName)
	defer free2()

	C.OpenCustomKeyboard(cfileName, ctitle, cbuffer, C.int(buflen), C.int(0), chandler)
}

// Probably changes the keybaord language
func LoadKeyboard() {
	keyboardLang, free := cString(defaultKeyboardLang)
	defer free()
	C.LoadKeyboard(keyboardLang)
}

// Close keyboard layout
func CloseKeyboard() {
	C.CloseKeyboard()
}

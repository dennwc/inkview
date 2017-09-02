package ink

/*
#include "inkview.h"

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"
import "image"

type Event interface {
	isEvent()
}

type UnknownEvent struct {
	Type   int
	P1, P2 int
}

func (UnknownEvent) isEvent() {}

type InitEvent struct{}

func (InitEvent) isEvent() {}

type ExitEvent struct{}

func (ExitEvent) isEvent() {}

type KeyEvent struct {
	Key   Key
	State KeyState
}

func (KeyEvent) isEvent() {}

type PointerEvent struct {
	image.Point
	State PointerState
}

func (PointerEvent) isEvent() {}

type TouchEvent struct {
	image.Point
	State TouchState
}

func (TouchEvent) isEvent() {}

type KeyState int

const (
	KeyStateDown    = KeyState(C.EVT_KEYDOWN)
	KeyStatePress   = KeyState(C.EVT_KEYPRESS)
	KeyStateUp      = KeyState(C.EVT_KEYUP)
	KeyStateRelease = KeyState(C.EVT_KEYRELEASE)
	KeyStateRepeat  = KeyState(C.EVT_KEYREPEAT)
)

type PointerState int

const (
	PointerUp   = PointerState(C.EVT_POINTERUP)
	PointerDown = PointerState(C.EVT_POINTERDOWN)
	PointerMove = PointerState(C.EVT_POINTERMOVE)
	PointerLong = PointerState(C.EVT_POINTERLONG)
	PointerHold = PointerState(C.EVT_POINTERHOLD)
)

type TouchState int

const (
	TouchUp   = TouchState(C.EVT_TOUCHUP)
	TouchDown = TouchState(C.EVT_TOUCHDOWN)
	TouchMove = TouchState(C.EVT_TOUCHMOVE)
)

// Key is a key code for buttons.
type Key int

const (
	KeyBack   = Key(C.KEY_BACK)
	KeyDelete = Key(C.KEY_DELETE)
	KeyOk     = Key(C.KEY_OK)
	KeyUp     = Key(C.KEY_UP)
	KeyDown   = Key(C.KEY_DOWN)
	KeyLeft   = Key(C.KEY_LEFT)
	KeyRight  = Key(C.KEY_RIGHT)
	KeyMinus  = Key(C.KEY_MINUS)
	KeyPlus   = Key(C.KEY_PLUS)
	KeyMenu   = Key(C.KEY_MENU)
	KeyMusic  = Key(C.KEY_MUSIC)
	KeyPower  = Key(C.KEY_POWER)
	KeyPrev   = Key(C.KEY_PREV)
	KeyNext   = Key(C.KEY_NEXT)
	KeyPrev2  = Key(C.KEY_PREV2)
	KeyNext2  = Key(C.KEY_NEXT2)
)

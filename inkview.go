package ink

/*
#include "inkview.h"

extern int main_handler(int t, int p1, int p2);

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"
import "image"

type Handler func(e Event)

var mainHandler Handler

// Run starts main event loop. It should be called before calling any other function.
func Run(h Handler) {
	if h == nil {
		panic("no handler")
	}
	mainHandler = h
	C.InkViewMain(C.iv_handler(C.main_handler))
}

//export goMainHandler
func goMainHandler(typ int, p1, p2 int) int {
	var e Event = UnknownEvent{Type: typ, P1: p1, P2: p2}
	switch typ {
	case 38, 39:
		return 0
	case C.EVT_INIT:
		e = InitEvent{}
	case C.EVT_EXIT:
		e = ExitEvent{}
	default:
		switch {
		case typ >= C.EVT_KEYDOWN && typ <= C.EVT_KEYREPEAT:
			e = KeyEvent{Key: Key(p1), State: KeyState(typ)}
		case typ >= C.EVT_POINTERUP && typ <= C.EVT_POINTERHOLD:
			e = PointerEvent{Point: image.Pt(p1, p2), State: PointerState(typ)}
		case typ >= C.EVT_TOUCHUP && typ <= C.EVT_TOUCHMOVE:
			e = TouchEvent{Point: image.Pt(p1, p2), State: TouchState(typ)}
		}
	}
	mainHandler(e)
	return 0
}

func OpenScreen() {
	C.OpenScreen()
}

func OpenScreenExt() {
	C.OpenScreenExt()
}

func Close() {
	C.CloseApp()
}

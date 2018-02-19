package ink

/*
#include "inkview.h"

extern int main_handler(int t, int p1, int p2);

#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"
import (
	"image"
	"sync"
)

var (
	mainMu  sync.Mutex
	mainApp App

	errMu   sync.Mutex
	mainErr error
)

func SetErr(err error) {
	errMu.Lock()
	mainErr = err
	errMu.Unlock()
}

// Run starts main event loop. It should be called before calling any other function.
func Run(app App) error {
	if app == nil {
		panic("no app")
	}
	mainMu.Lock()
	defer mainMu.Unlock()
	SetErr(nil)

	mainApp = app
	C.InkViewMain(C.iv_handler(C.main_handler))

	errMu.Lock()
	err := mainErr
	errMu.Unlock()
	return err
}

func handleEvent(typ int, p1, p2 int) bool {
	switch typ {
	case C.EVT_INIT:
		if err := mainApp.Init(); err != nil {
			mainErr = err
			Exit()
		}
		return true
	case C.EVT_EXIT:
		if err := mainApp.Close(); err != nil {
			mainErr = err
		}
		return true
	case C.EVT_SHOW:
		mainApp.Draw()
		return true
	//case C.EVT_HIDE:
	//case C.EVT_FOREGROUND:
	//	return mainApp.Show()
	//case C.EVT_BACKGROUND:
	//	return mainApp.Hide()
	case C.EVT_ORIENTATION:
		return mainApp.Orientation(Orientation(p1))
	default:
		switch {
		case typ >= C.EVT_KEYDOWN && typ <= C.EVT_KEYREPEAT:
			return mainApp.Key(KeyEvent{
				Key:   Key(p1),
				State: KeyState(typ),
			})
		case typ >= C.EVT_POINTERUP && typ <= C.EVT_POINTERHOLD:
			return mainApp.Pointer(PointerEvent{
				Point: image.Pt(p1, p2),
				State: PointerState(typ),
			})
		case typ >= C.EVT_TOUCHUP && typ <= C.EVT_TOUCHMOVE:
			return mainApp.Touch(TouchEvent{
				Point: image.Pt(p1, p2),
				State: TouchState(typ),
			})
		}
	}
	return false
}

//export goMainHandler
func goMainHandler(typ int, p1, p2 int) int {
	if handleEvent(typ, p1, p2) {
		return 1
	}
	return 0
}

//func OpenScreen() {
//	C.OpenScreen()
//}

//func OpenScreenExt() {
//	C.OpenScreenExt()
//}

// Exit can be called to exit an application event loop.
func Exit() {
	C.CloseApp()
}

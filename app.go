package ink

type App interface {
	// Init is called when application is started.
	Init() error
	// Close is called before exiting an application.
	Close() error

	// Draw is called each time an application view should be updated.
	// Can be queued by Repaint.
	Draw()

	//// Show is called when application becomes active.
	//// Delivered on application start and when switching from another app.
	//Show() bool
	//// Hide is called when application becomes inactive (switching to another app).
	//Hide() bool

	// Key is called on each key-related event.
	Key(e KeyEvent) bool
	// Pointer is called on each pointer-related event.
	Pointer(e PointerEvent) bool
	// Touch is called on each touch-related event.
	Touch(e TouchEvent) bool
	// Orientation is called each time an orientation of device changes.
	Orientation(o Orientation) bool
}

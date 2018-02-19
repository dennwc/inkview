package ink

import (
	"context"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

var DefaultFontHeight = 14

type RunFunc func(ctx context.Context, w io.Writer) error

func newLogWriter(log *Log, update func()) *logWriter {
	return &logWriter{log: log, update: update}
}

type logWriter struct {
	log    *Log
	update func()
}

func (w *logWriter) Write(p []byte) (int, error) {
	defer w.update()
	return w.log.Write(p)
}
func (w *logWriter) draw() {
	w.log.Draw()
}
func (w *logWriter) close() {
	w.log.Close()
}

func newCliApp(fnc RunFunc, c RunConfig) *cliApp {
	return &cliApp{cli: fnc, conf: c}
}

type cliApp struct {
	redraws int32 // atomic

	conf RunConfig
	cli  RunFunc

	wg   sync.WaitGroup
	err  error
	stop func()

	log *logWriter

	stopNet func()

	rmu     sync.Mutex
	running bool
}

func (app *cliApp) setRunning(v bool) {
	app.rmu.Lock()
	app.running = v
	app.rmu.Unlock()
}

func (app *cliApp) isRunning() bool {
	app.rmu.Lock()
	v := app.running
	app.rmu.Unlock()
	return v
}

func (app *cliApp) redraw() {
	// allow only one repaint in queue
	if atomic.CompareAndSwapInt32(&app.redraws, 0, 1) {
		Repaint()
	}
}
func (app *cliApp) draw() {
	ClearScreen()
	app.log.draw()
	FullUpdate()
	atomic.StoreInt32(&app.redraws, 0)
}

func (app *cliApp) println(args ...interface{}) {
	fmt.Fprintln(app.log, args...)
}

func (app *cliApp) Init() error {
	ClearScreen()
	l := NewLog(Pad(Screen(), 10), DefaultFontHeight)
	app.log = newLogWriter(l, app.redraw)

	if app.conf.Certs {
		now := time.Now()
		app.println("reading certs...")
		app.draw()
		if err := InitCerts(); err != nil {
			app.println("error reading certs:", err)
		} else {
			app.println("loaded certs in", time.Since(now))
		}
		app.draw()
	}

	if app.conf.Network {
		var err error
		app.stopNet, err = KeepNetwork()
		if err != nil {
			app.println("cannot connect to the network:", err)
		} else {
			app.println("network connected")
		}
		app.draw()
	}

	app.setRunning(true)

	ctx, cancel := context.WithCancel(context.Background())
	app.stop = cancel
	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		err := app.cli(ctx, app.log)
		if app.stopNet != nil {
			app.stopNet()
		}
		app.err = err
		if err != nil {
			app.println("error:", err)
		}
		app.println("<press any key to exit>")
		app.redraw()
	}()
	return nil
}

func (app *cliApp) stopCli() error {
	if !app.isRunning() {
		return app.err
	}
	app.stop()
	app.wg.Wait()
	app.setRunning(false)
	return app.err
}

func (app *cliApp) Close() error {
	err := app.stopCli()
	app.log.close()
	if app.stopNet != nil {
		app.stopNet()
	}
	return err
}

func (app *cliApp) Draw() {
	app.draw()
}

func (*cliApp) Show() bool {
	return false
}

func (*cliApp) Hide() bool {
	return false
}

func (app *cliApp) Key(e KeyEvent) bool {
	if app.isRunning() || (e.Key == KeyPrev && e.State == KeyStateDown) {
		Exit()
		return true
	}
	return false
}

func (app *cliApp) Pointer(e PointerEvent) bool {
	if app.isRunning() {
		Exit()
		return true
	}
	if e.State == PointerDown {
		app.redraw()
	}
	return true
}

func (*cliApp) Touch(e TouchEvent) bool {
	return false
}

func (*cliApp) Orientation(o Orientation) bool {
	return false
}

type RunConfig struct {
	Certs   bool // initialize certificate pool
	Network bool // keep networking enabled while app is running
}

// RunCLI starts a command-line application that can write to device display.
// Context will be cancelled when application is closed.
// Provided callback can use any SDK functions.
func RunCLI(fnc RunFunc, c *RunConfig) error {
	if c == nil {
		c = &RunConfig{}
	}
	return Run(newCliApp(fnc, *c))
}

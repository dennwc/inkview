package main

import (
	"github.com/dennwc/inkview"
)

func main() {
	var (
		log *ink.Log
	)
	ink.Run(func(e ink.Event) {
		switch e := e.(type) {
		case ink.InitEvent:
			ink.ClearScreen()
			log = ink.NewLog(ink.Pad(ink.Screen(), 10), 14)

			log.Println("Device:")
			log.Println(ink.DeviceModel())
			log.Println(ink.SoftwareVersion())
			log.Println(ink.HwAddress())
			log.Println()

			for _, name := range ink.Connections() {
				log.Printf("conn: %q", name)
			}
			log.Draw()

			ink.FullUpdate()
		case ink.ExitEvent:
			log.Close()
		case ink.PointerEvent:
			if e.State == ink.PointerDown {
				log.Printf("click: %v", e.Point)
				log.Draw()
				ink.SoftUpdate()
			}
		case ink.KeyEvent:
			if e.Key == ink.KeyPrev && e.State == ink.KeyStateDown {
				ink.Close()
			}
		}
	})
}

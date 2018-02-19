package main

import (
	"context"
	"fmt"
	"io"

	"github.com/dennwc/inkview"
)

func main() {
	ink.RunCLI(func(ctx context.Context, w io.Writer) error {
		fmt.Fprintln(w, "Device:")
		fmt.Fprintln(w, ink.DeviceModel())
		fmt.Fprintln(w, ink.SoftwareVersion())
		fmt.Fprintln(w, ink.HwAddress())
		fmt.Fprintln(w)

		for _, name := range ink.Connections() {
			fmt.Fprintf(w, "conn: %q", name)
		}
		return nil
	}, nil)
}

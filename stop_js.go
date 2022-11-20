//go:build js && wasm

package main

import (
	"os"
	"syscall/js"
)

func init() {
	js.Global().Set("stopWasm", js.FuncOf(func(this js.Value, args []js.Value) any {
		os.Exit(1)

		return nil
	}))
}

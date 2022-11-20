//go:build js && wasm

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"syscall/js"
)

// BASIC Version 3 Reference Manual (1-8, 7-9)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n23/mode/2up
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n81/mode/2up
func input(i any) {
	fmt.Print("? ")

	var r string

	wait := make(chan any)
	js.Global().Call("wasmPrompt").Call("then", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) > 0 {
			r = args[0].String()
		}

		wait <- nil

		return nil
	}))
	<-wait

	r = strings.Replace(r, "\n", "", -1)
	r = strings.Replace(r, "\r", "", -1)

	e := reflect.ValueOf(i).Elem()
	switch i.(type) {
	case *string:
		e.SetString(strings.ToUpper(strings.Trim(r, "\"")))
	case *int:
		o, err := strconv.ParseInt(r, 10, 64)
		if err != nil {
			fmt.Println(" ILLEGAL DATA, RETYPE INPUT")
			input(i)
			return
		}

		e.SetInt(o)
	case *float64:
		f, err := strconv.ParseFloat(r, 64)
		if err != nil {
			fmt.Println(" ILLEGAL DATA, RETYPE INPUT")
			input(i)
			return
		}

		e.SetFloat(f)
	}
}

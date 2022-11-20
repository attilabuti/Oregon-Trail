//go:build !js || !wasm

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// BASIC Version 3 Reference Manual (1-8, 7-9)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n23/mode/2up
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n81/mode/2up
func input(i any) {
	fmt.Print("? ")

	r, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		r = ""
	} else {
		r = strings.Trim(regexp.MustCompile(`\r?\n`).ReplaceAllString(r, ""), " ")
	}

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

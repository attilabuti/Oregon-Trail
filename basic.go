package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// BASIC Version 3 Reference Manual (5-2)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n53/mode/2up
// INT()
// math.Floor()

// BASIC Version 3 Reference Manual (5-1)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n53/mode/2up
func rnd(r int) float64 {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond)

	return rand.Float64()
}

// BASIC Version 3 Reference Manual (7-13)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n85/mode/2up
func tab(i int) string {
	if i > 1 {
		return strings.Repeat(" ", i-1)
	}

	return ""
}

// BASIC Version 3 Reference Manual (5-4)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n55/mode/2up
func clk(i int) float64 {
	now := time.Now()
	return float64(now.Hour()) + ((float64(1) / float64(3600)) * float64(((now.Minute() * 60) + now.Second())))
}

// BASIC Version 3 Reference Manual (4-2)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n45/mode/2up
func on_goto(v any, s ...func()) {
	t := reflect.TypeOf(v).Kind()
	if t != reflect.Int && t != reflect.Float64 || len(s) == 0 {
		return
	}

	var val int
	if t == reflect.Float64 {
		val = int(math.Round(v.(float64)))
	} else {
		val = v.(int)
	}

	if val <= 0 || val > len(s) {
		fmt.Println("ON EXPRESSION OUT OF RANGE")
		return
	}

	for i, fn := range s {
		if val == (i + 1) {
			fn()
			return
		}
	}
}

// BASIC Version 3 Reference Manual (7-12)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n83/mode/2up
func format_numeric(n any) string {
	t := reflect.TypeOf(n).Kind()
	if t != reflect.Int && t != reflect.Float64 {
		return ""
	}

	var num string

	if t == reflect.Float64 {
		prec := 2
		if math.Mod(n.(float64), 1.0) == 0 {
			prec = 0
		}

		num = strconv.FormatFloat(n.(float64), 'f', prec, 64)
	} else if t == reflect.Int {
		num = strconv.Itoa(n.(int))
	}

	num = strings.Replace(num, "-0.", "-.", -1)
	num = strings.Replace(num, "0.", ".", -1)

	if num[0:1] == "-" {
		num = fmt.Sprintf("%s ", num)
	} else {
		num = fmt.Sprintf(" %s ", num)
	}

	return num
}

// BASIC Version 3 Reference Manual (7-13)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n85/mode/2up
func print_zoning(s ...any) string {
	out := ""
	for _, e := range s {
		zone := 15
		v := ""
		if t := reflect.TypeOf(e).Kind(); t == reflect.Int || t == reflect.Float64 {
			v = format_numeric(e)
		} else {
			v = fmt.Sprintf("%v ", e)
		}

		if len(v) > 15 {
			if len(v)%15 > 0 {
				zone = ((len(v) / 15) + 1) * 15
			} else {
				zone = (len(v) / 15) * 15
			}
		}

		space := zone - len(v)
		if space < 0 {
			space = 0
		}

		out += v + strings.Repeat(" ", space)
	}

	return out
}

var dataTable []float64

// BASIC Version 3 Reference Manual (7-23)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n95/mode/2up
func data(i ...float64) {
	dataTable = i
}

var readPosition = 0

// BASIC Version 3 Reference Manual (7-23)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n95/mode/2up
func read(i *float64) {
	if len(dataTable)-1 >= readPosition {
		*i = dataTable[readPosition]
		readPosition++
	} else {
		fmt.Println("END OF DATA")
	}
}

// BASIC Version 3 Reference Manual (7-4)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n75/mode/2up
func restore() {
	readPosition = 0
}

// BASIC Version 3 Reference Manual (3-4)
// https://archive.org/details/bitsavers_cdccyberlaBASICOct80_41102953/page/n43/mode/2up
func stop() {
	os.Exit(0)
}

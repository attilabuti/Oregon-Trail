package main

import (
	"fmt"
	"math"
)

// ***SETTING DATE***
//
// 1230
func settingDate() {
	if M >= 2040 {
		finalTurn() // 5430
		return
	}

	// 1250
	D3 = D3 + 1
	fmt.Print("\n")
	fmt.Print("MONDAY ")

	// 1310
	_d := map[int]string{
		1:  "APRIL 12",     // 1310
		2:  "APRIL 26",     // 1330
		3:  "MAY 10",       // 1350
		4:  "MAY 24",       // 1370
		5:  "JUNE 7",       // 1390
		6:  "JUNE 21",      // 1410
		7:  "JULY 5",       // 1430
		8:  "JULY 19",      // 1450
		9:  "AUGUST 2",     // 1470
		10: "AUGUST 16",    // 1490
		11: "AUGUST 31",    // 1510
		12: "SEPTEMBER 13", // 1530
		13: "SEPTEMBER 27", // 1550
		14: "OCTOBER 11",   // 1570
		15: "OCTOBER 25",   // 1590
		16: "NOVEMBER 8",   // 1610
		17: "NOVEMBER 22",  // 1630
		18: "DECEMBER 6",   // 1650
		19: "DECEMBER 20",  // 1670
	}

	_s := int(math.Round(D3))

	if _s == 20 {
		// 1690
		fmt.Println("YOU HAVE BEEN ON THE TRAIL TOO LONG  ------")
		fmt.Println("YOUR FAMILY DIES IN THE FIRST BLIZZARD OF WINTER")

		dying() // 5170

		return
	} else {
		fmt.Print(_d[_s] + " ")
	}

	// 1720
	fmt.Println("1847")
	fmt.Print("\n")

	beginTurn() // 1750
}

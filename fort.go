package main

import (
	"fmt"
	"math"
)

// ***STOPPING AT FORT***
//
// 2290
func fort() {
	fmt.Println("ENTER WHAT YOU WISH TO SPEND ON THE FOLLOWING")

	// 2300
	fmt.Print("FOOD ")
	fortSubRoutine()
	F = F + 2/3.0*P

	// 2420
	fmt.Print("AMMUNITION ")
	fortSubRoutine()
	B = math.Floor(B + 2/3.0*P*50)

	// 2450
	fmt.Print("CLOTHING ")
	fortSubRoutine()
	C = C + 2/3.0*P

	// 2480
	fmt.Print("MISCELLANEOUS SUPPLIES")
	fortSubRoutine()
	M1 = M1 + 2/3.0*P

	M = M - 45

	// 2520
	beforeEating() // 2720
}

// 2330
func fortSubRoutine() {
	input(&P)
	if P < 0 {
		return
	}

	T = T - P
	if T >= 0 {
		return
	}

	fmt.Println("YOU DON'T HAVE THAT MUCH--KEEP YOUR SPENDING DOWN")
	fmt.Println("YOU MISS YOUR CHANCE TO SPEND ON THAT ITEM")

	T = T + P
	P = 0
}

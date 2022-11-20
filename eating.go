package main

import (
	"fmt"
	"math"
)

// 2720
func beforeEating() {
	if F >= 13 {
		eating() // 2750
	} else {
		// ***DYING***
		// 5060
		fmt.Println("YOU RAN OUT OF FOOD AND STARVED TO DEATH")
		dying() // 5170
	}
}

// ***EATING***
//
// 2750
func eating() {
	for {
		fmt.Println("DO YOU WANT TO EAT (1) POORLY  (2) MODERATELY")
		fmt.Print("OR (3) WELL ")
		input(&E)

		// 2780
		if E >= 1 && E <= 3 {
			break
		}
	}

	// 2800
	E = math.Floor(E)
	F = F - 8 - 5*E

	// 2820
	if F >= 0 {
		// 2860
		M = M + 200 + (A-220)/5.0 + 10*rnd(-1)
		L1, C1 = 0, 0

		ridersAttack() // 2890
	} else {
		F = F + 8 + 5*E
		fmt.Println("YOU CAN'T EAT THAT WELL")

		eating() // 2750
	}
}

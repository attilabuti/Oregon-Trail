package main

import "fmt"

// ***HUNTING***
//
// 2540
func hunting() {
	if B <= 39 {
		fmt.Println("TOUGH---YOU NEED MORE BULLETS TO GO HUNTING")
		fortHuntContinue() // 2080

		return
	}

	// 2570
	M = M - 45
	shootingSubRoutine() // 6140

	if B1 <= 1 {
		// **BELLS IN LINE 2660**
		// 2660
		fmt.Println("RI\aGHT BETWEE\aN THE EYE\aS---YOU GOT A\a BIG ONE!!\a!!")
		// fmt.Println("RIGHT BETWEEN THE EYES---YOU GOT A BIG ONE!!!!")
		fmt.Println("FULL BELLIES TONIGHT!")

		F = F + 52 + rnd(-1)*6
		B = B - 10 - rnd(-1)*4
		// 2720
	} else {
		if 100*rnd(-1) < 13*B1 {
			// 2710
			fmt.Println("YOU MISSED---AND YOUR DINNER GOT AWAY.....")
			// 2720
		} else {
			// 2610
			F = F + 48 - 2*B1
			fmt.Println("NICE SHOT--RIGHT ON TARGET--GOOD EATIN' TONIGHT!!")
			B = B - 10 - 3*B1
			// 2720
		}
	}

	beforeEating() // 2720
}

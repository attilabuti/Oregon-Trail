package main

import (
	"fmt"
	"math"
)

// ***SHOOTING SUB-ROUTINE***
// THE METHOD OF TIMING THE SHOOTING (LINES 6210-6240)
// WILL VARY FROM SYSTEM TO SYSTEM.  FOR EXAMPLE, H-P
// USERS WILL PROBABLY PREFER TO USE THE 'ENTER' STATEMENT.
// IF TIMING ON THE USER'S SYSTEM IS HIGHLY SUSCEPTIBLE
// TO SYSTEM RESPONSE TIME, THE FORMULA IN LINE 6240 CAN
// BE TAILORED TO ACCOMODATE THIS BY EITHER INCREASING
// OR DECREASING THE 'SHOOTING' TIME RECORDED BY THE SYSTEM.
//
// 6140
func shootingSubRoutine() {
	SS = map[int]string{
		1: "BANG",
		2: "BLAM",
		3: "POW",
		4: "WHAM",
	}

	// 6190
	S6 = math.Floor(rnd(-1)*4 + 1)
	fmt.Printf("TYPE %s\n", SS[int(S6)])

	// 6210
	B3 = clk(0)
	input(&CS)
	B1 = clk(0)

	// 6240
	B1 = ((B1 - B3) * 3600) - (D9 - 1)

	fmt.Print("\n")

	// 6255
	if B1 < 0 {
		B1 = 0
	}

	if CS != SS[int(S6)] {
		B1 = 9
	}
}

// ***ILLNESS SUB-ROUTINE***
//
// 6300
func illnessSubRoutine() {
	if 100*rnd(-1) < 10+35*(E-1) {
		// 6370
		fmt.Println("MILD ILLNESS---MEDICINE USED")
		M = M - 5
		M1 = M1 - 2
	} else {
		if 100*rnd(-1) < 100-(40/math.Pow(4, (E-1))) {
			// 6410
			fmt.Println("BAD ILLNESS---MEDICINE USED")
			M = M - 5
			M1 = M1 - 5
		} else {
			fmt.Println("SERIOUS ILLNESS---")
			fmt.Println("YOU MUST STOP FOR MEDICAL ATTENTION")
			M1 = M1 - 10
			S4 = 1
		}
	}

	// 6440
	if M1 < 0 {
		// ***DYING***
		// 5110
		fmt.Println("YOU RAN OUT OF MEDICAL SUPPLIES")
		diedOf()
	} else {
		if L1 == 1 {
			mountainsSettingDate() // 4940
		} else {
			mountains() // 4710
		}
	}
}

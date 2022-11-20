package main

import (
	"fmt"
	"math"
)

// ***MOUNTAINS***
//
// 4710
func mountains() {
	if M <= 950 {
		settingDate() // 1230
		return
	}

	if rnd(-1)*10 <= 9-(math.Pow((M/100-15), 2)+72)/(math.Pow((M/100-15), 2)+12) {
		// 4730
		fmt.Println("RUGGED MOUNTAINS")

		if rnd(-1) > .1 { // 4740
			if rnd(-1) > .11 { // 4780
				// 4840
				fmt.Println("THE GOING GETS SLOW")
				M = M - 45 - rnd(-1)/.02
			} else {
				fmt.Println("WAGON DAMAGED!---LOSE TIME AND SUPPLIES")
				M1 = M1 - 5
				B = B - 200
				M = M - 20 - 30*rnd(-1)
			}
		} else {
			// 4750
			fmt.Println("YOU GOT LOST---LOSE VALUABLE TIME TRYING TO FIND TRAIL!")
			M = M - 60
		}
	}

	// 4860
	if F1 != 1 {
		F1 = 1
		if rnd(-1) < .8 {
			blizzard()
			return
		} else {
			fmt.Println("YOU MADE IT SAFELY THROUGH SOUTH PASS--NO SNOW")
		}
	}

	// 4900
	if M < 1700 {
		mountainsSettingDate() // 4940
	} else {
		if F2 == 1 {
			mountainsSettingDate() // 4940
		} else {
			F2 = 1
			if rnd(-1) < .7 {
				blizzard()
			} else {
				mountainsSettingDate() // 4940
			}
		}
	}
}

// 4940
func mountainsSettingDate() {
	if M <= 950 {
		M9 = 1
	}

	settingDate()
}

// 4970
func blizzard() {
	fmt.Println("BLIZZARD IN MOUNTAIN PASS--TIME AND SUPPLIES LOST")

	L1 = 1
	F = F - 25
	M1 = M1 - 10
	B = B - 300
	M = M - 30 - 40*rnd(-1)

	// 5030
	if C < 18+2*rnd(-1) {
		illnessSubRoutine() // 6300
	} else {
		mountainsSettingDate() // 4940
	}
}

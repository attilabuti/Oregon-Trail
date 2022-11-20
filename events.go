package main

import "fmt"

// ***SELECTION OF EVENTS***
//
// 3550
func selectionOfEvents() {
	D1 = 0
	restore()
	R1 = 100 * rnd(-1)

	for {
		// 3580
		D1 = D1 + 1
		if D1 == 16 {
			evHelpfulIndians() // 4670
			return
		} else {
			read(&D)

			if R1 <= D {
				break
			}
		}
	}

	// 3620
	data(6, 11, 13, 15, 17, 22, 32, 35, 37, 42, 44, 54, 64, 69, 95)

	if D1 > 10 {
		// 3650
		on_goto(D1-10,
			evPoisonousSnake,    // 4220
			evWagonSwamped,      // 4290
			evWildAnimalsAttack, // 4340
			evHailStorm,         // 4560
			evIllness,           // 4610
			evHelpfulIndians,    // 4670
		)
	} else {
		// 3640
		on_goto(D1,
			evWagonBreaksDown, // 3660
			evOxInjuresLeg,    // 3700
			evBadLuck,         // 3740
			evOxWandersOff,    // 3790
			evYourSonGetsLost, // 3820
			evUnsafeWater,     // 3850
			evHeavyRains,      // 3880
			evBanditsAttack,   // 3960
			evFireInWagon,     // 4130
			evHeavyFog,        // 4190
		)
	}
}

// 3660
func evWagonBreaksDown() {
	fmt.Println("WAGON BREAKS DOWN--LOSE TIME AND SUPPLIES FIXING IT")
	M = M - 15 - 5*rnd(-1)
	M1 = M1 - 8
	mountains() // 4710
}

// 3700
func evOxInjuresLeg() {
	fmt.Println("OX INJURES LEG---SLOWS YOU DOWN REST OF TRIP")
	M = M - 25
	A = A - 20
	mountains() // 4710
}

// 3740
func evBadLuck() {
	fmt.Println("BAD LUCK---YOUR DAUGHTER BROKE HER ARM")
	fmt.Println("YOU HAD TO STOP AND USE SUPPLIES TO MAKE A SLING")
	M = M - 5 - 4*rnd(-1)
	M1 = M1 - 2 - 3*rnd(-1)
	mountains() // 4710
}

// 3790
func evOxWandersOff() {
	fmt.Println("OX WANDERS OFF---SPEND TIME LOOKING FOR IT")
	M = M - 17
	mountains() // 4710
}

// 3820
func evYourSonGetsLost() {
	fmt.Println("YOUR SON GETS LOST---SPEND HALF THE DAY LOOKING FOR HIM")
	M = M - 10
	mountains() // 4710
}

// 3850
func evUnsafeWater() {
	fmt.Println("UNSAFE WATER--LOSE TIME LOOKING FOR CLEAN SPRING")
	M = M - 10*rnd(-1) - 2
	mountains() // 4710
}

// 3880
func evHeavyRains() {
	if M > 950 {
		// 4490
		fmt.Print("COLD WEATHER---BRRRRRRR!---YOU ")
		if C <= 22+4*rnd(-1) {
			fmt.Print("DON'T ")
			C1 = 1
		}

		fmt.Println("HAVE ENOUGH CLOTHING TO KEEP YOU WARM")

		if C1 == 0 { // 4710
			mountains()
		} else {
			illnessSubRoutine() // 6300
		}
	} else {
		// 3890
		fmt.Println("HEAVY RAINS---TIME AND SUPPLIES LOST")
		F = F - 10
		B = B - 500
		M1 = M1 - 15
		M = M - 10*rnd(-1) - 5

		mountains() // 4710
	}
}

// 3960
func evBanditsAttack() {
	fmt.Println("BANDITS ATTACK")
	shootingSubRoutine() // 6140
	B = B - 20*B1

	if B >= 0 { // 3980
		if B1 <= 1 { // 4030
			// 4100
			fmt.Println("QUICKEST DRAW OUTSIDE OF DODGE CITY!!!")
			fmt.Println("YOU GOT 'EM!")

			mountains() // 4710

			return
		}
	} else {
		fmt.Println("YOU RAN OUT OF BULLETS---THEY GET LOTS OF CASH")
		T = T / 3.0
	}

	// 4040
	fmt.Println("YOU GOT SHOT IN THE LEG AND THEY TOOK ONE OF YOUR OXEN")
	K8 = 1
	fmt.Println("BETTER HAVE A DOC LOOK AT YOUR WOUND")
	M1 = M1 - 5
	A = A - 20

	mountains() // 4710
}

// 4130
func evFireInWagon() {
	fmt.Println("THERE WAS A FIRE IN YOUR WAGON--FOOD AND SUPPLIES DAMAGED")
	F = F - 40
	B = B - 400
	M1 = M1 - rnd(-1)*8 - 3
	M = M - 15
	mountains() // 4710
}

// 4190
func evHeavyFog() {
	fmt.Println("LOSE YOUR WAY IN HEAVY FOG---TIME IS LOST")
	M = M - 10 - 5*rnd(-1)
	mountains() // 4710
}

// 4220
func evPoisonousSnake() {
	fmt.Println("YOU KILLED A POISONOUS SNAKE AFTER IT BIT YOU")
	B = B - 10
	M1 = M1 - 5

	if M1 >= 0 {
		// 4280
		mountains() /// 4710
	} else {
		// 4260
		fmt.Println("YOU DIE OF SNAKEBITE SINCE YOU HAVE NO MEDICINE")
		dying() // 5170
	}
}

// 4290
func evWagonSwamped() {
	fmt.Println("WAGON GETS SWAMPED FORDING RIVER--LOSE FOOD AND CLOTHES")
	F = F - 30
	C = C - 20
	M = M - 20 - 20*rnd(-1)
	mountains() /// 4710
}

// 4340
func evWildAnimalsAttack() {
	fmt.Println("WILD ANIMALS ATTACK!")
	shootingSubRoutine() // 6140

	if B > 39 { // 4360
		if B1 > 2 { // 4410
			fmt.Println("SLOW ON THE DRAW---THEY GOT AT YOUR FOOD AND CLOTHES")
		} else {
			fmt.Println("NICE SHOOTIN' PARDNER---THEY DIDN'T GET MUCH")
		}

		// 4450
		B = B - 20*B1
		C = C - B1*4
		F = F - B1*8

		mountains() // 4710
	} else {
		// 4370
		fmt.Println("YOU WERE TOO LOW ON BULLETS--")
		fmt.Println("THE WOLVES OVERPOWERED YOU")
		K8 = 1

		diedOf() // 5120
	}
}

// 4560
func evHailStorm() {
	fmt.Println("HAIL STORM---SUPPLIES DAMAGED")

	M = M - 5 - rnd(-1)*10
	B = B - 200
	M1 = M1 - 4 - rnd(-1)*3

	mountains() // 4710
}

// 4610
func evIllness() {
	if E == 1 {
		illnessSubRoutine() // 6300
	} else {
		if E == 3 {
			// 4650
			if rnd(-1) < .5 {
				illnessSubRoutine() // 6300
			} else {
				mountains() // 4710
			}
		} else {
			// 4630
			if rnd(-1) > .25 {
				illnessSubRoutine() // 6300
			} else {
				mountains() // 4710
			}
		}
	}
}

// 4670
func evHelpfulIndians() {
	fmt.Println("HELPFUL INDIANS SHOW YOU WHERE TO FIND MORE FOOD")
	F = F + 14
	mountains() // 4710
}

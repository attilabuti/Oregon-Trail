package main

import (
	"fmt"
	"math"
)

// ***RIDERS ATTACK***
//
// 2890
func ridersAttack() {
	if rnd(-1)*10 > (math.Pow((M/100-4), 2)+72)/(math.Pow((M/100-4), 2)+12)-1 {
		selectionOfEvents() // 3550
	} else {
		// 2900
		fmt.Print("RIDERS AHEAD.  THEY ")
		S5 = 0

		if rnd(-1) >= .8 {
			fmt.Print("DON'T ")
			S5 = 1
		}

		// 2950
		fmt.Println("LOOK HOSTILE")
		fmt.Println("TACTICS")

		// 2970
		tactics()

		// 3040
		if S5 == 1 { // DON'T LOOK HOSTILE
			// 3330
			ridersFriendly()
		} else if T1 > 1 { // LOOK HOSTILE - (2) ATTACK / (3) CONTINUE / (4) CIRCLE WAGONS
			// 3110
			ridersHostile()
		} else { // LOOK HOSTILE - (1) RUN
			// 3060
			M = M + 20
			M1 = M1 - 15
			B = B - 150
			A = A - 40

			afterAttack() // 3470
		}
	}
}

// 2970
func tactics() {
	for {
		fmt.Println("(1) RUN  (2) ATTACK  (3) CONTINUE  (4) CIRCLE WAGONS")

		if rnd(-1) <= .2 {
			S5 = 1 - S5
		}

		// 3000
		input(&T1)

		if T1 >= 1 && T1 <= 4 {
			break
		}
	}

	T1 = math.Floor(T1)
}

// 3330
func ridersFriendly() {
	switch T1 {
	case 1: // 3340 - (1) RUN
		M = M + 15
		A = A - 10
	case 2: // 3380 - (2) ATTACK
		M = M - 5
		B = B - 100
	case 4: // 3430 - (4) CIRCLE WAGONS
		M = M - 20
	}

	afterAttack() // 3470
}

// 3110
func ridersHostile() {
	switch T1 {
	case 2: // 3120
		shootingSubRoutine() // 6140
		B = B - B1*40 - 80
		afterShooting() // 3140
	case 3: // 3250
		if rnd(-1) > .8 {
			// 3450
			fmt.Println("THEY DID NOT ATTACK")
			selectionOfEvents() // 3550
		} else {
			// 3260
			B = B - 150
			M1 = M1 - 15
			afterAttack() // 3470
		}
	case 4: // 3290
		shootingSubRoutine() // 6140
		B = B - B1*30 - 80
		M = M - 25
		afterShooting() // 3140
	}
}

// 3140
func afterShooting() {
	if B1 > 1 {
		if B1 <= 4 { // 3170
			// 3220
			fmt.Println("KINDA SLOW WITH YOUR COLT .45")
		} else {
			// 3180
			fmt.Println("LOUSY SHOT---YOU GOT KNIFED")
			K8 = 1
			fmt.Println("YOU HAVE TO SEE OL' DOC BLANCHARD")
		}
	} else {
		// 3150
		fmt.Println("NICE SHOOTING---YOU DROVE THEM OFF")
	}

	afterAttack() // 3470
}

// 3470
func afterAttack() {
	if S5 == 0 {
		// 3500
		fmt.Println("RIDERS WERE HOSTILE--CHECK FOR LOSSES")

		if B >= 0 {
			selectionOfEvents() // 3550
		} else {
			fmt.Println("YOU RAN OUT OF BULLETS AND GOT MASSACRED BY THE RIDERS")
			dying()
		}
	} else {
		fmt.Println("RIDERS WERE FRIENDLY, BUT CHECK FOR POSSIBLE LOSSES")
		selectionOfEvents() // 3550
	}
}

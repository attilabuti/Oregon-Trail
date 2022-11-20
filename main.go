package main

import (
	"fmt"
	"math"
)

/**/
// PROGRAM NAME - OREGON        VERSION:01/01/78
// ORIGINAL PROGRAMMING BY BILL HEINEMANN - 1971
// SUPPORT RESEARCH AND MATERIALS BY DON RAVITSCH,
//      MINNESOTA EDUCATIONAL COMPUTING CONSORTIUM STAFF
// CDC CYBER 70/73-26     BASIC 3.1
// DOCUMENTATION BOOKLET 'OREGON' AVAILABLE FROM
//    MECC SUPPORT SERVICES
//    2520 BROADWAY DRIVE
//    ST. PAUL, MN  55113
//
//  *FOR THE MEANING OF THE VARIABLES USED, LIST LINES 6470-6790*
//
// 160
func main() {
	fmt.Print("DO YOU NEED INSTRUCTIONS  (YES/NO) ")
	// RANDOMIZE REMOVED
	input(&CS)

	if CS != "NO" {
		instructions() // 210
	}

	rifle()            // 690
	initialPurchases() // 810

	// 1190
	fmt.Print("\nMONDAY MARCH 29 1847\n\n")

	beginTurn() // 1750
}

// ***BEGINNING EACH TURN***
//
// 1750
func beginTurn() {
	if F < 0 {
		F = 0
	}

	if B < 0 {
		B = 0
	}

	if C < 0 {
		C = 0
	}

	if M1 < 0 {
		M1 = 0
	}

	if F < 13 {
		fmt.Println("YOU'D BETTER DO SOME HUNTING OR BUY FOOD AND SOON!!!!")
	}

	// 1850
	F = math.Floor(F)
	B = math.Floor(B)
	C = math.Floor(C)
	M1 = math.Floor(M1)
	T = math.Floor(T)
	M = math.Floor(M)
	M2 = M

	// 1920
	if S4 == 1 || K8 == 1 {
		// 1950
		T = T - 20
		if T < 0 {
			// ***DYING***
			// 5080
			T = 0
			fmt.Println("YOU CAN'T AFFORD A DOCTOR")
			diedOf() // 5120

			return
		} else {
			fmt.Println("DOCTOR'S BILL IS $20")
			K8, S4 = 0, 0
		}
	}

	// 1990
	if M9 == 1 {
		// 2020
		fmt.Println("TOTAL MILEAGE IS 950")
		M9 = 0
	} else {
		fmt.Printf("TOTAL MILEAGE IS%v\n", format_numeric(M))
	}

	// 2040
	fmt.Println(print_zoning("FOOD", "BULLETS", "CLOTHING", "MISC. SUPP.", "CASH"))
	fmt.Println(print_zoning(F, B, C, M1, T))

	if X1 == -1 {
		// 2170
		huntContinue()

		// 2270
		on_goto(X, fort, hunting, beforeEating)
	} else {
		// 2070
		X1 = X1 * (-1)
		fortHuntContinue()
	}
}

// 2170
func huntContinue() {
	for {
		fmt.Println("DO YOU WANT TO (1) HUNT, OR (2) CONTINUE")
		input(&X)

		if X != 1 {
			X = 2
		}
		X = X + 1

		if X == 3 || B > 39 {
			// 2260
			X1 = X1 * (-1)
			return
		} else {
			fmt.Println("TOUGH---YOU NEED MORE BULLETS TO GO HUNTING")
		}
	}
}

// 2080
func fortHuntContinue() {
	fmt.Print("DO YOU WANT TO (1) STOP AT THE NEXT FORT, (2) HUNT, ")
	fmt.Println("OR (3) CONTINUE")
	input(&X)

	if X > 2 || X < 1 {
		// 2150
		X = 3
	} else {
		X = math.Floor(X)
	}

	// 2270
	on_goto(X, fort, hunting, beforeEating)
}

// ***FINAL TURN***
//
// 5430
func finalTurn() {
	F9 = (2040 - M2) / (M - M2)
	F = F + (1-F9)*(8+5*E)
	fmt.Print("\n")

	// **BELLS IN LINES 5470,5480**
	fmt.Println("YOU\a FINALLY ARRI\aVED AT ORE\aGON CITY\a") // 5470
	fmt.Println("AFTER\a 2040 LONG MILES\a---HOORAY!!\a!!!")  // 5480
	// fmt.Println("YOU FINALLY ARRIVED AT OREGON CITY")
	// fmt.Println("AFTER 2040 LONG MILES---HOORAY!!!!!")
	fmt.Println("A REAL PIONEER!")
	fmt.Print("\n")

	// 5510
	F9 = math.Floor(F9 * 14)
	D3 = D3*14 + F9
	F9 = F9 + 1

	if F9 >= 8 {
		// 5550
		F9 = F9 - 7
	}

	// 5560
	_d := map[int]string{
		1: "MONDAY",    // 5570
		2: "TUESDAY",   // 5590
		3: "WEDNESDAY", // 5600
		4: "THURSDAY",  // 5630
		5: "FRIDAY",    // 5650
		6: "SATURDAY",  // 5670
		7: "SUNDAY",    // 5690
	}

	fmt.Print(_d[int(math.Round(F9))] + " ")

	// 5700
	_m := ""
	switch {
	case D3 <= 124: // 5710
		D3 = D3 - 93
		_m = "JULY"
	case D3 <= 155: // 5750
		D3 = D3 - 124
		_m = "AUGUST"
	case D3 <= 185: // 5790
		D3 = D3 - 155
		_m = "SEPTEMBER"
	case D3 <= 216: // 5830
		D3 = D3 - 185
		_m = "OCTOBER"
	case D3 <= 246: // 5870
		D3 = D3 - 216
		_m = "NOVEMBER"
	case D3 > 246: // 5900
		D3 = D3 - 246
		_m = "DECEMBER"
	}

	fmt.Printf(_m+" %v 1847", format_numeric(D3))

	// 5920
	fmt.Print("\n")
	fmt.Println(print_zoning("FOOD", "BULLETS", "CLOTHING", "MISC. SUPP.", "CASH"))

	// 5940
	if B < 0 {
		B = 0
	}

	if C < 0 {
		C = 0
	}

	if M1 < 0 {
		M1 = 0
	}

	if T < 0 {
		T = 0
	}

	if F < 0 {
		F = 0
	}

	// 6040
	fmt.Println(print_zoning(math.Floor(F), math.Floor(B), math.Floor(C), math.Floor(M1), math.Floor(T)))
	fmt.Print("\n")

	fmt.Println(tab(11) + "PRESIDENT JAMES K. POLK SENDS YOU HIS")
	fmt.Println(tab(17) + "HEARTIEST CONGRATULATIONS")
	fmt.Print("\n")
	fmt.Println(tab(11) + "AND WISHES YOU A PROSPEROUS LIFE AHEAD")
	fmt.Print("\n")
	fmt.Println(tab(22) + "AT YOUR NEW HOME")

	stop()
}

func init() {
	SS = make(map[int]string)
	data(6, 11, 13, 15, 17, 22, 32, 35, 37, 42, 44, 54, 64, 69, 95)
}

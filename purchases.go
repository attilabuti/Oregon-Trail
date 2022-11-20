package main

import "fmt"

// **INITIAL PURCHASES***
//
// 810
func initialPurchases() {
	X1 = -1
	K8, S4, F1, F2, M, M9, D3 = 0, 0, 0, 0, 0, 0, 0

	// 830
	for {
		fmt.Print("\n\n")

		// 850
		for {
			fmt.Print("HOW MUCH DO YOU WANT TO SPEND ON YOUR OXEN TEAM ")
			input(&A)

			if A >= 200 {
				if A <= 300 {
					break
				} else {
					fmt.Println("TOO MUCH")
				}
			} else {
				fmt.Println("NOT ENOUGH")
			}
		}

		_buy := func(_p *float64, _t string) {
			for {
				fmt.Print(_t)
				input(_p)

				if *_p >= 0 {
					break
				} else {
					fmt.Println("IMPOSSIBLE")
				}
			}
		}

		_buy(&F, "HOW MUCH DO YOU WANT TO SPEND ON FOOD ")                    // 930
		_buy(&B, "HOW MUCH DO YOU WANT TO SPEND ON AMMUNITION ")              // 980
		_buy(&C, "HOW MUCH DO YOU WANT TO SPEND ON CLOTHING ")                // 1030
		_buy(&M1, "HOW MUCH DO YOU WANT TO SPEND ON MISCELLANEOUS SUPPLIES ") // 1080

		// 1130
		T = 700 - A - F - B - C - M1
		if T >= 0 {
			B = 50 * B
			fmt.Printf("AFTER ALL YOUR PURCHASES, YOU NOW HAVE %v DOLLARS LEFT\n", format_numeric(T))
			return
		} else {
			fmt.Println("YOU OVERSPENT--YOU ONLY HAD $700 TO SPEND.  BUY AGAIN")
		}
	}
}

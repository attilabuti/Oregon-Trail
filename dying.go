package main

import "fmt"

// ***DYING***
//
// 5170
func dying() {
	fmt.Print("\n")
	fmt.Println("DUE TO YOUR UNFORTUNATE SITUATION, THERE ARE A FEW")
	fmt.Println("FORMALITIES WE MUST GO THROUGH")
	fmt.Print("\n")

	fmt.Println("WOULD YOU LIKE A MINISTER?")
	input(&CS)

	fmt.Println("WOULD YOU LIKE A FANCY FUNERAL?")
	input(&CS)

	fmt.Println("WOULD YOU LIKE US TO INFORM YOUR NEXT OF KIN?")
	input(&CS)

	if CS == "YES" {
		// 5310
		fmt.Println("THAT WILL BE $4.50 FOR THE TELEGRAPH CHARGE.")
	} else {
		// 5280
		fmt.Println("BUT YOUR AUNT SADIE IN ST. LOUIS IS REALLY WORRIED ABOUT YOU")
	}

	fmt.Print("\n")

	// 5330
	fmt.Println("WE THANK YOU FOR THIS INFORMATION AND WE ARE SORRY YOU")
	fmt.Println("DIDN'T MAKE IT TO THE GREAT TERRITORY OF OREGON")
	fmt.Println("BETTER LUCK NEXT TIME")
	fmt.Print("\n\n")

	// 5380
	fmt.Println(tab(30) + "SINCERELY")
	fmt.Print("\n")
	fmt.Println(tab(17) + "THE OREGON CITY CHAMBER OF COMMERCE")

	stop()
}

// 5120
func diedOf() {
	fmt.Print("YOU DIED OF ")

	if K8 == 1 {
		// 5160
		fmt.Println("INJURIES")
	} else {
		// 5140
		fmt.Println("PNEUMONIA")
	}

	dying()
}

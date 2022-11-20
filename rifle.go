package main

import "fmt"

// RIFLE SKILLS
//
// 690
func rifle() {
	fmt.Print("\n\n")
	fmt.Print(_rifleText)
	input(&D9)

	if D9 > 5 {
		D9 = 0
	}
}

var _rifleText = `HOW GOOD A SHOT ARE YOU WITH YOUR RIFLE?
  (1) ACE MARKSMAN,  (2) GOOD SHOT,  (3) FAIR TO MIDDLIN'
	     (4) NEED MORE PRACTICE,  (5) SHAKY KNEES
ENTER ONE OF THE ABOVE -- THE BETTER YOU CLAIM YOU ARE, THE
FASTER YOU'LL HAVE TO BE WITH YOUR GUN TO BE SUCCESSFUL.
`

package main

import (
	"fmt"
	"math"
	"os"
)

var universe = [9]bool{
	false, false, false,
	false, true, false,
	false, false, false,
}

// κάνε τους constant κάποια στιγμή μέσω function,
// στην Go δεν υποστηρίζονται άμεσα
var universeWin = [9]bool{
	true, true, true,
	true, false, true,
	true, true, true,
}

var universeLose = [9]bool{
	false, false, false,
	false, false, false,
	false, false, false,
}

func pAsciiUniverse() {
	asciiUniverse := [9]string{}
	for i, v := range universe {
		// fmt.Println(i, v)
		switch v {
		case true:
			asciiUniverse[i] = "*"
		case false:
			asciiUniverse[i] = "-"
		}
	}
	fmt.Println(asciiUniverse[0:3])
	fmt.Println(asciiUniverse[3:6])
	fmt.Println(asciiUniverse[6:9])
	fmt.Println()
}

func main() {
	fmt.Println(`
Instructions:
Each star is in a galaxy. When you shoot a star,
everything in its galaxy changes. All stars become
black holes and all black holes become stars.

Numbers represent stars, @ its galaxies, - are black holes:
1@-	@2@	-@3	@--	-@-	--@	---	---	---
@@-	---	-@@	4--	@5@	--6	@@-	---	-@@
---	---	---	@--	-@-	--@	7@-	@8@	-@9

The game starts with the		---
universe in this situation:		-*-
* represent stars, - black holes	---

You win by	***
changing	*-*
to this:	***

You lose	---
if you		---
get this:	---

Current universe and star positions:
`)
	pAsciiUniverse()

	fmt.Println("( 0 exits the game )")

	// main game loop
	for universe != universeLose || universe != universeWin {

		var f float64
		var n int
		ypodoxh := func() {
		enar3h:
			fmt.Println()
			fmt.Print(`Give star position to shoot (1-9): `)
			_, err := fmt.Scan(&f)
			if err != nil {
				fmt.Println("Letters or symbols not accepted")
				goto enar3h
			}
			if f == 0 {
				os.Exit(0)
			}
			if f < 1 || f > 9 || f-math.Ceil(f) != 0 {
				fmt.Println("Only integer numbers between 1-9 are accepted")
				goto enar3h
			}
			n = int(f)
			if universe[n-1] == false {
				fmt.Println("No star in position", n)
				goto enar3h
			}
		}
		ypodoxh()

		fmt.Println()

		switch n {
		case 1:
			universe[0] = !universe[0]
			universe[1] = !universe[1]
			universe[3] = !universe[3]
			universe[4] = !universe[4]
		case 2:
			universe[0] = !universe[0]
			universe[1] = !universe[1]
			universe[2] = !universe[2]
		case 3:
			universe[1] = !universe[1]
			universe[2] = !universe[2]
			universe[4] = !universe[4]
			universe[5] = !universe[5]
		case 4:
			universe[0] = !universe[0]
			universe[3] = !universe[3]
			universe[6] = !universe[6]
		case 5:
			universe[1] = !universe[1]
			universe[3] = !universe[3]
			universe[4] = !universe[4]
			universe[5] = !universe[5]
			universe[7] = !universe[7]
		case 6:
			universe[2] = !universe[2]
			universe[5] = !universe[5]
			universe[8] = !universe[8]
		case 7:
			universe[3] = !universe[3]
			universe[4] = !universe[4]
			universe[6] = !universe[6]
			universe[7] = !universe[7]
		case 8:
			universe[6] = !universe[6]
			universe[7] = !universe[7]
			universe[8] = !universe[8]
		case 9:
			universe[4] = !universe[4]
			universe[5] = !universe[5]
			universe[7] = !universe[7]
			universe[8] = !universe[8]
		}
		pAsciiUniverse()
	}
	if universe == universeWin {
		fmt.Println("You won!!!")
	}
	if universe == universeLose {
		fmt.Println("Sorry, you lost.")
	}
}

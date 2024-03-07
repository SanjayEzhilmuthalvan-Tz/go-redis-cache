// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5 // less is more difficult
	usage    = `Welcome to the Lucky Number Game!

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

Wanna play?

(Provide -v flag to see the picked numbers.)
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var verbose bool

	if args[0] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("please enter a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please enter a positive number.")
		return
	}

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			fmt.Println("congatulations you won")
			return
		}
	}

	fmt.Println("you lost ....Try again?")
}

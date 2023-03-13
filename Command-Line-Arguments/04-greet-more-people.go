package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Greet More People
//
// RESTRICTIONS
//  1. Be sure to match the expected output below
//  2. Store the length of the arguments in a variable
//  3. Store all the arguments in variables as well
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 people!
//  Hello great bilbo !
//  Hello great balbo !
//  Hello great bungo !
//  Nice to meet you all.
// ---------------------------------------------------------

func main() {
	// TYPE YOUR CODE HERE
	names, length := os.Args, len(os.Args)-1
	// BONUS #1:
	// Observe the error if you pass less then 3 arguments.
	if length > 3 {
		return
	}
	// Search on the web how to solve that.
	fmt.Println("There are", length, "people!")
	fmt.Println(" Hello great", names[1])
	fmt.Println(" Hello great", names[2])
	fmt.Println(" Hello great", names[3])
	fmt.Println("Nice to meet you all.")
}

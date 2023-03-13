package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Find the Rectangle's Perimeter
//
//  1. Find the perimeter of a rectangle
//     Its width is  5
//     Its height is 6
//
//  2. Assign the result to the `perimeter` variable
//
//  3. Print the `perimeter` variable
//
// HINT
//  Formula = 2 times the width and height
//
// EXPECTED OUTPUT
//  22
//
// BONUS
//  Find more formulas here and calculate them in new programs
//  https://www.mathsisfun.com/area.html
// ---------------------------------------------------------

func main() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	// USE THE VARIABLES ABOVE WHEN CALCULATING YOUR RESULT
	perimeter = 2 * (width + height)

	// ADD YOUR CODE BELOW
	fmt.Println(perimeter)
}

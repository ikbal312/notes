package main

import "fmt"

/*

Tuple Used IN:
	- Multiple return values
	- Error handling
	- Named return values
	- Multiple assignments
	- Returning multiple values from a function
*/

// // get the power series of integer a and returns tuples of square of a and cube of a
// func powerSeries(a int) (int, int) {

// 	return a * a, a * a * a
// }

// func main() {
// 	var square int
// 	var cube int
// 	square, cube = powerSeries(3)
// 	fmt.Println("Square: ", square, "Cube: ", cube)
// }

// get the power series of integer a and returns tuples of square of a and cube of a
// func powerSeries(a int) (square int, cube int) {
// 	square = a * a
// 	cube = a * a * a
// return				 // naked return  # bad practice
// }

// func main() {
// 	var square int
// 	var cube int
// 	square, cube = powerSeries(4)
// 	fmt.Println("Square: ", square, "Cube: ", cube)
// }

// get the power series of integer a and returns tuples of square of a and cube of a
func powerSeries(a int) (square int, cube int, err error) {
	square = a * a
	cube = a * a * a
	return square, cube, nil
}

func main() {
	var square int
	var cube int
	square, cube, _ = powerSeries(4)
	fmt.Println("Square: ", square, "Cube: ", cube)
}

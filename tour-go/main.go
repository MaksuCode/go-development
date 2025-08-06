package main

import (
	"fmt"
	"math"
)

// ! I stopped at "Type Conversions 13/17 : https://go.dev/tour/basics/13"
// * Variables can be declared in blocks
// * When variables are declared without an initial value, a zero value is given
var (
	i int
	f float64
	b bool
	s string
)

func main() {

	fmt.Println(math.Pi)

	fmt.Printf("Sum is %v\n", add(1, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}

// * A simple function
func add(x, y int) int {
	return x + y
}

// * Function with multiple results
func swap(x, y string) (string, string) {
	return y, x
}

//   - Naked return : This is mostly avoided since it is not readiable. Using explicit return like the example
//     above is a better way
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

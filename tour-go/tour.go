package main

import (
	"fmt"
	"math"
)

// ! I stopped at "Type Conversions 13/17 : https://go.dev/tour/basics/13"
// * Variables can be declared in blocks
// * When variables are declared without an initial value, a zero value is given

func main() {

	fmt.Println(math.Pi)

	// Functions
	fmt.Printf("Sum is %v\n", Add(1, 2))

	x, y := Swap("hello", "world")
	fmt.Println(x, y)

	fmt.Println(Split(17))

	// Zero values
	PrintZeroValues()

	// Type Conversions
	z := Pythagorean(3, 4)
	fmt.Println(z)

	var1 := 1
	var2 := 3.142
	var3 := 0.867 + 0.5i

	fmt.Printf("Types of vars %T , %T and %T\n", var1, var2, var3)

	// Constant
	// Can not be declated with ":=" syntax
	const lang = "Go"
	const truth = true
	fmt.Printf("%v rules? %v\n", lang, truth)

}

/* A simple function */
func Add(x, y int) int {
	return x + y
}

/* Function with multiple results */
func Swap(x, y string) (string, string) {
	return y, x
}

/*
	Named return : This is mostly avoided since it is not readable. Using explicit return like the example

above is a better way
*/
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/* Zero Values of different types */
func PrintZeroValues() {
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

/* Type conversion */
func Pythagorean(x, y int) uint {
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	return z
}

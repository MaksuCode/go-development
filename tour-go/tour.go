package main

import (
	"fmt"
	"math"
)

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

/*
		Simple For loop example
	    the init statement: executed before the first iteration
		the condition expression: evaluated before every iteration
		the post statement: executed at the end of every iteration
*/
func CountFromZero(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += i
	}
	return sum
}

/*
Foor loop with range

	If you do not need to use the init value in the loop, you can simply add a "_"
*/
func CountNumbers(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

/*
Init and Post statements are optional. This now looks like a while loop in another languages.

	In Go we do not have a while loop. For is the only loop
*/
func WhileLoop() int {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func ForeverLoop() {
	for {
		fmt.Println("This will run forever..")
	}
}

package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {

	ints := map[string]int64{
		"first":  12,
		"second": 25,
	}

	floats := map[string]float64{
		"first":  54.22,
		"second": 42.1,
	}

	fmt.Printf("Non Generic Sums : %v and %v\n",
		SumInts(ints),
		SumFloats(floats),
	)

	fmt.Printf("Generic Sums : %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, value := range m {
		s += value
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, value := range m {
		s += value
	}
	return s
}

func SumIntsOrFloats[Key comparable, Value int64 | float64](numbers map[Key]Value) Value {
	var sum Value
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumNumbers[Key comparable, Value Number](numbers map[Key]Value) Value {
	var sum Value
	for _, value := range numbers {
		sum += value
	}
	return sum
}

package main

import "fmt"

func main() {

	ints := map[string]int64{
		"first":  12,
		"second": 25,
	}

	floasts := map[string]float64{
		"first":  54.22,
		"second": 42.1,
	}

	fmt.Printf("Non Generic Sums : %v and %v\n",
		SumInts(ints),
		SumFloats(floasts),
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

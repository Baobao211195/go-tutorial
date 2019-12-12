package main

import (
	"fmt"
	"math"
)

func powf(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	//fmt.Println("value of v", v) // v is undefined, refer to if-short-statement
	return lim
}

func main() {
	fmt.Println(
		powf(3, 2, 10),
		powf(3, 3, 20),
	)
}

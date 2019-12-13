package main

import (
	"fmt"
	"math/rand"
)

// can import separated
import "time"

func main() {
	// Uses default seed of 1, result will be 81:
	fmt.Println(rand.Intn(1000))

	// Seed the random number generator using the current time (nanoseconds since epoch):
	rand.Seed(time.Now().UnixNano())

	// Hard to predict...but is it possible? I know the day, and hour, probably minute...
	fmt.Println(rand.Intn(1000))
}

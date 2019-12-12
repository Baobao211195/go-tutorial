package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// if có thể khai báo một câu lệnh ngắn  trước khi tới vói điều kiện
	// biến được khai báo trong lệnh ngắn chỉ có scope cho tới khi kết thúc if
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}

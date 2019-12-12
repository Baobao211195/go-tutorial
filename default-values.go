package main

import (
	"fmt"
	"math"
)

func main() {
	//var i int
	//var f float64
	//var b bool
	//var s string
	//fmt.Printf("%v %v %v %q\n", i, f, b, s)


	//  use  expression T(v) converts the value v to the type T.
	x, y := 3, 4
	var f  = math.Sqrt(float64(x*x + y*y))
	var z  = uint(f)
	fmt.Println(x, y, z)

	// when khai báo như dưới thì phải sử dụng biến đó
	var (
		k  = 3
		u  = 8
	)
	fmt.Println(k)
	fmt.Println(u)
}

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // không có dấu ()
		sum += i
	}
	fmt.Println(sum)

	// init and post are optional
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// case 2 lollllllllll
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// case 3 for forever
	//for {
	//}

}

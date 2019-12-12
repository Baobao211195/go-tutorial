package go_basic

import "fmt"

// case 1
func add(x int, y int) int {
	return x + y
}

// case 2
func multiple(x, y int) int { // shortened declaration variable
	return x * y
}

// one function can be return any number of result
func swap(x, y string) (string, string) {
	return y, x
}

// đã chỉ rõ cụ thể biến được return về
// là x và y nên trong lệnh return thì không cần thiết chỉ cụ thể
// chỉ nên áp dụng trong những function ngắn và nhỏ
// không nên áp dụng cho nhưng function lớn.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// call function
	fmt.Println(add(2, 3))

	// call truc tiep (giống kiểu lamda trong java hoặc python) anonynous function
	fmt.Println(func(x int, y int) int { return x - y }(5, 1))

	fmt.Println(multiple(2, 3))

	// one function can be return any number of result
	var a, b = swap("oanh", "van")
	fmt.Println(a, b)

	// errors
	//var x = swap("oanh", "vankem")
	//fmt.Println(x)

	//Named return values
	i, j := split(18)
	fmt.Println(i, j)
}

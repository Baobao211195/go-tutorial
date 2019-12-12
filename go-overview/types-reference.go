package go_overview

import "fmt"

func main() {
	// khi khai báo biến mà không chỉ định kiểu
	// thì mặc định là kiểu của biến sẽ là kiểu của giá trị nằm ở bên phải
	var i int
	y := i // j will be an int
	fmt.Println(y)

	v := 42.334 //
	fmt.Printf("v is of type %T\n", v)

}

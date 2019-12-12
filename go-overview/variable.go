package go_overview

import . "fmt"

var c, python, java bool // variables declaration

// khởi tạo biến
// nếu được khởi tạo khi khai báo thì có thể bỏ qua khai báo kiểu của biến.
var i, j = 1, 2

func main() {
	var i int
	Println(i, c, python, java)

	// d = true
	// python = false
	// java = "no!"
	var d, python, java = true, false, "no!"
	Println(i, j, d, python, java)
}

package add

import "fmt"

// Addをaddとしてしまうと外部（ここではmain.go）から参照できない
// 外部から参照する変数や関数は大文字から始める

func Inp() int {
	var a int
	fmt.Scan(&a)

	return a
}

func Add(a, b int) int {
	x := 0
	for i := 0 ; i < a ; i++ { 
		x += b
	}
    return x
}
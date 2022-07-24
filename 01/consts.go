package main

import "fmt"

func main() {
	//常量的声明和变量形式类似

	const pi = 3.14
	const e = 2.71828

	//一次声明多个常量
	const (
		a = 1
		b = "b"
	)

	//多个常量使用相同的值
	const (
		n1 = 100
		n2
		n3
	)

	//n1 n2 n3 都拥有相同的值
	fmt.Println(n1, n2, n3)
}

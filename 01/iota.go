package main

import "fmt"

func main() {
	//iota只能在常量中使用
	//iota 在 const 中每新增一行常量声明将使iota计数一次, 可以理解为 const语句块中的行索引
	//iota在const关键字出现时将被重置为0

	const (
		n1 = iota
		n2
		n3
		n4
	)

	const (
		m1 = iota //0
		m2        //1
		_         //2 被忽略了
		m4        //3
	)

	const (
		i1 = iota //0
		i2 = 100  //100
		i3 = iota //2
		i4        //3
	)

	const (
		j1 = iota //重置为 0
		j2
	)

	const (
		a, b = iota + 1, iota + 2 // 1 , 2
		c, d                      // 2 , 3
		e, f                      // 3 , 4
	)

	fmt.Println(i1, i2, i3, i4)
	fmt.Println(j1, j2)
}

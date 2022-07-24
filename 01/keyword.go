package main

func main() {
	//var 变量名称 变量类型
	var name string
	var age int
	var ok bool

	//批量声明
	var (
		str  string
		i    int
		isOk bool
	)

	//变量声明的时候 会初始化会变量类型的零值
	//整型浮点型默认值为 0
	//字符串默认值为 空字符串
	//布尔值默认为 false
	//切片,函数,指针 默认为 nil

	//声明并且初始化
	var foo int = 10
	//使用 类型推导
	var balance = 10000

	//简短声明 :=
	i := 100

	//匿名变量 _ 或者丢弃m某个值
	//匿名变量不占用空间,不会分配内存,匿名变量不存在重复声明

	//注意
	//函数外的语句必须以 var const func 等关键字开始
	// := 简写不能使用在函数外
	//_多用于占位,表示忽略值
}

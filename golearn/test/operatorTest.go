package test

import "fmt"

func operatorTest() {
	// 算数运算
	a := 100
	b := 200
	fmt.Println(a + b)
	fmt.Println(a / b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(b % a)
	// ++/--是单独的语句不能用到表达式内
	a++
	fmt.Println(a)
	fmt.Println(a == b)
	println(a > b)
	println(a < b)
	println(a >= b)
	println(a <= b)
	println(a != b)
	// 逻辑运算
	c := true
	d := false
	println(c && d)
	println(c || d)
	println(!c)
	e := 4
	f := 8
	// 与
	println(e & f)
	// 或
	println(e | f)
	// 异或
	println(e ^ f)
	// 移位
	println(e << 2)
	println(e >> 2)

	e += 100
	println(e)
}

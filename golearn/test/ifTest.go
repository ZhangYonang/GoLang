package test

import "fmt"

func ifTest(int a) {
	b := 2
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
	// 初始化信息可直接放到if表达式中
	if age := 19; age > 18 {
		println("成年")
	} else {
		println("未成年")
	}
}

func ScanTest() {
	var name string
	var age int
	fmt.Print("请输入姓名年龄:>")
	fmt.Scan(&name, &age)
	fmt.Print("name:" + name + "age:")
	fmt.Println(age)
}
func ifelseTest() {
	var num int
	fmt.Print("请输入一个数字:>")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

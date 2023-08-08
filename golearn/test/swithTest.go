package test

import "fmt"

func switchTest() {
	var a int
	fmt.Scan(a)
	// 单条件
	switch a {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	default:
		fmt.Println("无此等级")
	}
}
func switchTest2() {
	// 多条件
	var key int
	fmt.Print("请输入星期几:>")
	fmt.Scan(&key)
	switch key {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	default:
		fmt.Print("非法输入")
	}
}
func switchTest3() {
	// 判断条件
	var key int
	fmt.Print("请输入分数:>")
	fmt.Scan(&key)
	switch {
	case key >= 80:
		fmt.Println("优秀")
	case key >= 60:
		fmt.Println("良好")
	default:
		fmt.Println("不及格")
	}
	// 想连续执行添加fallthrough go中无需break

}

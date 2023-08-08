package test

import "fmt"

func forTest() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	i := 0
	for ; i <= 10; i++ {
		println(i)
	}
	println(i)
	for i <= 20 {
		println(i)
		i++
	}
}
func forRangeTest() {
	mape := make(map[string]string, 0)
	mape["name"] = "tom"
	mape["age"] = "18"
	var a = [...]int{1, 2, 3}
	for i, v := range a {
		println(i, v)
	}
	for _, v := range a {
		println(v)
	}
	for key, v := range mape {
		println(key, v)
	}
}

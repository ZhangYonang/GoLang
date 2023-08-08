package test

import (
	"bytes"
	"fmt"
	"strings"
)

func StringTest() {
	// 字符串的赋值
	var s string = "赋值方法1:var s string=\"\""
	var s1 = "方法2:var s1=\"\""
	s2 := "方法3: s2:=\"\""
	// 多行字符串
	s4 := `
	LINE1
	LINE2
	LINE3
	`
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s4)
	// 字符串连接
	s3 := "Elysia"
	s5 := "18"
	fmt.Println(s3 + s5)
	// 使用strings中方法连接
	name := "Elysia"
	age := "18"
	join := strings.Join([]string{name, age}, ",")
	fmt.Println(join)
	// buffer
	var buffer bytes.Buffer
	buffer.WriteString("Elysia")
	buffer.WriteString(",")
	buffer.WriteString("18")
	fmt.Println(buffer)

	// 字符串的切片
	s6 := "WuHuqiFei"
	a := 2
	b := 5
	/*表示取s中a位置的字符串*/
	fmt.Println(s6[a])
	/*取s中a到b的字符串*/
	fmt.Println(s6[a:b])
	/*表示从a开始到最后*/
	fmt.Println(s6[a:])
	/*表示从0开始取到b*/
	fmt.Println(s6[:b])

	// 字符串函数
	/*字符串长度*/
	fmt.Println(len(s6))
	/*字符串分割*/
	fmt.Println(strings.Split(s6, ""))
	/*是否包含某字符串*/
	fmt.Println(strings.Contains(s6, "wuhu"))
	/*转换为小写*/
	fmt.Println(strings.ToLower(s6))
	/*转换为大写*/
	fmt.Println(strings.ToUpper(s6))
	/*以。。开头*/
	fmt.Println(strings.HasPrefix(s6, "Wu"))
	/*以。。结尾*/
	fmt.Println(strings.HasSuffix(s6, "ei"))

}

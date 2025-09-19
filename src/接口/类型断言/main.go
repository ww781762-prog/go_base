package main

import "fmt"

func funType(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("数字")
	case string:
		fmt.Println("字符串")
	default:
		fmt.Println("不能识别的类型")

	}

}

func funType2(t interface{}) {
	switch t.(type) {
	case int:
		// 可以直接t 去获取转后的内容
		fmt.Println("数字", t)
	case string:

		fmt.Println("字符串", t)
	default:
		fmt.Println("不能识别的类型", t)
	}

}
func main() {
	funType(1)
	funType("a")
	funType2("a")

}

package main

import "fmt"

func main() {
	var a int
	var b string
	// 中间的分隔符表示，输入的时候以什么为分隔符,回车表示结束
	fmt.Scanf("%d,%s", &a, &b)
	fmt.Println(a, b)
}

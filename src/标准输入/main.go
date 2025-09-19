package main

import "fmt"

func main() {
	var a int
	var b string
	// 中间的，是分隔符，表示输入的时候以什么为分隔符, Scanf的默认分割符为空格
	//默认回车的时候表示结束输入
	fmt.Scanf("%d%s", &a, &b)
	fmt.Scanf("%d,%s", &a, &b)
	// 以回车作为换行符
	fmt.Scanf("%d\n%s", &a, &b)
	fmt.Println(a, b)

	var A int
	var B string
	// 就是以空格或回车作为分割符，然后把输入赋值给A和B
	// 自动跳过空白（空格、换行、tab），依次把输入内容扫描到提供的参数变量里。
	fmt.Scan(&A, &B)
	fmt.Println(A, B)

	var A1 int
	var B1 string
	// 就是以空格作为分割符，然后把输入赋值给A1和B1
	fmt.Scanln(&A1, &B1)
	fmt.Println(A1, B1)
}

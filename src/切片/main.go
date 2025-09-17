package main

import "fmt"

// 切片的扩容，切片的扩容实际是生成更大的数组，而后切片引用新的数组。申请一片更大的内存，然后把原来的内容复制过来，让切片引用这块内存
// 扩容的策略是原来的容量的翻倍
func main1() {
	// 1是切片的长度，2是切片的容量
	// 长度，切片默认会初始化的长度，容量是最多可以使用的长度，超过了容量，会扩容
	a := make([]int, 5, 10)
	fmt.Println(a)

	for i := 0; i < 15; i++ {
		a = append(a, i)
		fmt.Printf("a=%v,address=%p,len=%d,cap=%d\n", a, a, len(a), cap(a))
	}
	b := make([]int, 10, 10)
	fmt.Println(b)
}

// slien copy

func main() {
	var a []int = []int{1, 2}
	var b []int = []int{3, 4, 5, 6, 7}
	copy(a, b)
	fmt.Printf("%v,%v", a, b)
}

package main

import "fmt"

func defer_test() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func defer_test2() {
	var i int = 0
	defer func() {
		fmt.Println("defer i是：", i)
	}()
	i = 100
}

func main() {
	defer_test()
	defer_test2()
}

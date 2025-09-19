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

func defer_test3() int {
	x := 5
	defer func() {
		x = x + 5
		fmt.Println("defer_test3 defer", x)
	}()
	return x
}

func defer_test4() (x int) {
	defer func() {
		x = x + 1
	}()
	return 5
}

func main() {
	//defer_test()
	//defer_test2()
	fmt.Println("defer_test3 函数返回值", defer_test3())
	fmt.Println("defer_test4 函数返回值", defer_test4())

}

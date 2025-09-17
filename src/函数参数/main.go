package main

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func calc(a, b int, op func(a, b int) int) int {
	return op(a, b)
}

func calc1(base int) (func(int) int, func(int) int) {
	add1 := func(i int) int {
		base += i
		return base
	}
	sub1 := func(i int) int {
		base -= i
		return base
	}
	return add1, sub1
}

func testClosure5() {
	for i := 0; i < 5; i++ {
		go func() {
			i++
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}

func main() {
	calc(10, 10, add)
	calc(10, 10, sub)

	a1, c1 := calc1(10)
	fmt.Println(a1(1), c1(2))
	fmt.Println(a1(3), c1(4))
	testClosure5()

}

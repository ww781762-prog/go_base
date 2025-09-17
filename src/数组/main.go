package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var b [3]int = [3]int{0: 10, 1: 10, 2: 10}
	fmt.Printf("%T,%T", a, b)
}

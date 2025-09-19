package main

import "fmt"

func dist(a interface{}) {
	fmt.Printf("%v", a)
}

func main() {

	fmt.Println("a")
	fmt.Println(1)
	var a map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(a)

}

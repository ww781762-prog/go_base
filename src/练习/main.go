package main

import "fmt"

func testReverseString(ts string) {
	var as string
	for i := 0; i < len(ts); i++ {
		as += string(ts[len(ts)-i-1])
	}
	print(as)
}

func testReverseStringV2(ts string) {
	var r []rune = []rune(ts)
	for i := 0; i < len(r)/2; i++ {
		tep := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tep
	}
	fmt.Println(string(r))
}
func QuMo() {
	a := 529 % 10
	print(a)
}
func main() {
	testReverseStringV2("wang你好")
	QuMo()

}

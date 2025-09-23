package main

import "fmt"

func editMap(m map[string]interface{}) {
	m["name"] = "wawa"
}

func main() {
	var m1 = map[string]map[string]interface{}{}
	m1["stu1"] = make(map[string]interface{})
	editMap(m1["stu1"])
	fmt.Println(m1)
}

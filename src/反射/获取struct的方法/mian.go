package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
}

func (p *person) GetName(x string, v int) int {
	fmt.Println("GetName", x)
	return v
}
func main() {
	p := person{Name: "wawawa"}
	v := reflect.ValueOf(&p)
	m := v.MethodByName("GetName")
	// 自己创建Value类型的参数列表数组，如果函数没有参数也需要一个空的
	in := []reflect.Value{
		// reflect.ValueOf("10")
		// 需要把参数转换成 reflect.Value 类型，可以通过 reflect.ValueOf进行转换
		reflect.ValueOf("10"),
		reflect.ValueOf(20),
	}
	d := m.Call(in)
	// 返回值是一个包含Value的列表
	fmt.Println(d[0])

	fmt.Println(v.Type().Elem().Field(0).Tag.Get(""))
}

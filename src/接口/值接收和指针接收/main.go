package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	fmt.Println("Ping")
}
func (*St) Pang() {
	fmt.Println("Pang")
}

func main() {
	// St和*St的方法集不一样
	// St的方法集 Ping
	// *St的方法集 包含St的方法集和*St

	var s1 *St
	s1 = &St{Name: "dadasda"}
	var j Inter = s1
	p := j.(Inter)
	p.Pang()
	// 因为指针类型没有实现这个接口， 接口里面做了一个转换
	// *(&p).Ping()  先取出p里面存储的地址，而后再根据地址取出值
	p.Ping()

	st := St{Name: "xxx"}
	// 指针类型实现了，值类型不能调用
	// a里面存储的是一个值类型的struct对象，那么要调用st.Pang()的时候。要先取st的地址而后调用，(&st).Pang，
	st.Pang()
	// 一个变量存储在接口类型的变量中之后那么这个变量之前的地址就获取不到了，因为进行的是值copy
	//j = st
	// 调用这个会报错，
	j.Pang()

	// &st 就是 *st类型的
	var i Inter = &st
	o := i.(Inter)
	o.Ping()
	o.Pang()

	// 调整满足是ok是true,条件不满足的时候ok是false,
	// o的值为i绑定的值得副本
	if o, ok := i.(Inter); ok {
		o.Ping()
	}

	// 如下会报错因为i没有实现Anter接口。
	//p:=i.(Anter)
	//p.String()

	//如下报错,因为St方法集不包含 Pang
	//s:=i.(St)

	s := i.(*St)
	fmt.Printf("%s", s.Name)
}

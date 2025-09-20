package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type St struct{}

func (St) Ping() {
	fmt.Println("ping")
}

func (*St) Pang() {
	fmt.Println("pang")
}

func main() {
	// 定义一个存放指针的变量，
	// 里面存放的值为nil
	var st *St = nil
	//接口值在 Go 里是一个二元组： (type, value)
	//	type：动态类型（例如 St 或 *St）
	//	value：具体的动态值（例如一个 struct 拷贝，或者一个指针地址）
	var it Inter = st
	//赋值过程是：
	//	1.	在接口 it 里写入动态类型：*St
	//	2.	在接口 it 里写入动态值：*st 的一个 拷贝，也就是nic

	fmt.Printf("%v\n", st)
	fmt.Printf("%p\n", it)

	//因为接口变量有两个字段，一个是实例类型，另一个是指向实例的指针，两个都为空，才为空
	// 下面判断的接口为true，因为*St 和st的方法集不同。
	// 此时 it 的动态类型是 *St（非空），动态值是 nil。
	// 只要动态类型不为 nil，接口值就不为 nil。
	if it != nil {
		// 这个不会报错，因为调用Pang的时候需要传递的是St的指针，St的指针存在的，切不为空的
		// it 里面存储的是一个St实例为空，那么要调用it.Pang()的时候。要先取st的地址而后调用，(&it).Pang，
		it.Pang()
		// *(&it).Pang 如下会报错因为 Ping 调用的是st 对象，st对象是空，所以会报错
		// &it是取 it里面存储的内存地址，取到的值是st的一个指针，Pang是*St，需要的是一个指针 所以可以调用
		// Ping 需要传递一个St的实例，获取St实例的方法是这样 *(&it).Pang
		//it.Ping()
	}
}

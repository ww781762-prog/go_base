package main

import "fmt"

// 计算薪水的 接口
type Employer interface {
	CalcSalary() float32
}

// 人员
type Programer struct {
	name  string
	base  float32
	extra float32
}

func NewProgramer(name string, base float32, extra float32) Programer {
	return Programer{
		name:  name,
		base:  base,
		extra: extra,
	}
}

// 人员1的薪水计算方式
func (p Programer) CalcSalary() float32 {
	return p.base
}

type Sale struct {
	name  string
	base  float32
	extra float32
}

func NewSale(name string, base float32, extra float32) Sale {

	return Sale{
		name:  name,
		base:  base,
		extra: extra,
	}
}

// 人员2的薪水计算方式
func (s Sale) CalcSalary() float32 {
	return s.base + s.extra*s.base
}

func calcAll(emp []Employer) (out float32) {
	for _, e := range emp {
		out += e.CalcSalary()
	}
	return
}

func main() {
	var p1 Programer = NewProgramer("p1", 800, 0)
	var p4 Programer = NewProgramer("p4", 800, 0)

	var s3 Sale = NewSale("s3", 500, 0.3)
	var s4 Sale = NewSale("s4", 500, 0.3)

	var employList []Employer
	employList = append(employList, p1)
	employList = append(employList, p4)
	employList = append(employList, s3)
	employList = append(employList, s4)

	//for _,i:=range employList{
	//	fmt.Println(    i.CalcSalary())
	//}
	fmt.Println("本月支持人力陈本：", calcAll(employList))
}

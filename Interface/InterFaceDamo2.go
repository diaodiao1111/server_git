package Interface

import "fmt"

type Aint interface {
	test()
}

type Bint interface {
	test1()
}

type Cint interface {
	Aint
	Bint
	test2()
}

type Students struct {
}

func (stu *Students) test() {
	fmt.Println(1)
}

func (stu *Students) test1() {
	fmt.Println(2)
}

func (stu *Students) test2() {
	fmt.Println(3)
}

func Run(stu Students) {
	stu.test()
	stu.test2()
	stu.test1()
}

//func main() {
//	// 接口调用的方式
//
//// 1
//
//	var stus Students
//
//	var c Cint = &stus
//
//	c.test2()
//	c.test1()
//	c.test()
//
//// 2
//	var stu Students
//
//	stu.test1()
//	stu.test2()
//	stu.test()
//
//}

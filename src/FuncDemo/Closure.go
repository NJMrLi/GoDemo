package FuncDemo

import "fmt"

//函数闭包问题，下面的代码输出的结果是什么
func ClosureDemo1(){
	var j int = 5
	a := func()(func()) {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}


func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func ClosureDemo2(){
	fmt.Println("----ClosureDemo2----")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func ClosureDemo3(){
	fmt.Println("----ClosureDemo3----")
	var flist []func()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		flist = append(flist, func() {
			fmt.Println(i)
		})
	}

	for _, f := range flist {
		f()
	}
}

//同名的变量i作为内部的局部变量，覆盖了原来循环中的i，此时闭包中的变量不在是共享外循环的i，而是都有各自的内部同名变量i，
func ClosureDemo4(){
	fmt.Println("----ClosureDemo4----")
	var flist []func()
	for i := 0; i < 3; i++ {

		i := i //给i变量重新赋值，
		fmt.Println(i)
		flist = append(flist, func() {
			fmt.Println(i)
		})
	}
	for _, f := range flist {
		f()
	}
}



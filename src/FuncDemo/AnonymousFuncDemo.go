package FuncDemo

import "fmt"

//定义一个函数
func anonyTest1(){
	fmt.Println("anonyTest1")
}

//将改函数赋值给一个变量f，执行f
func AnonyTest(){
	f:= anonyTest1
	f()
}


func AnonyTest2(){
	f:= func() {
		fmt.Println("AnonyTest2")
	}
	f()

	//或者
	func() {
		fmt.Println("AnonyTest2...")
	}()

}

func AnonyTest3(){

	var i=0
	defer func() {
		fmt.Printf("defer func i=%v \n",i)
	}()

	 defer fmt.Printf("defer i=%v \n",i)

	for;i<10; i++{
	}

	fmt.Printf("i=%v \n",i)
}


func Calc(a,b int, op func(int,int)int) int {
	return op(a,b)
}

func add(a,b int) int{
	return a+b
}

func sub(a,b int)int{
	return  a-b
}

func AnonyTest4(){

	var a = 2
	var b = 1

	var x = Calc(a,b,add)
	var y = Calc(a,b,sub)

	fmt.Printf("x=%v, y=%v \n",x,y)
}

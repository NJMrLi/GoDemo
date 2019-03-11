package FuncDemo

import "fmt"

func DeferTest1(){
	defer fmt.Println("我是 defer1")
	defer fmt.Println("我是 defer2")
	fmt.Println("我是DeferTest1")
	fmt.Println("我是DeferTest2")
}


func DeferTest2(){
	i:= 0
	defer  fmt.Printf("defer i=%d\t",i)
	for ;i<=10;i++{
		fmt.Printf("i=%d\t",i)
	}
	fmt.Println()
}

//打印1-100的质数
















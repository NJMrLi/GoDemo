package LoopDemo

import "fmt"

type Foo struct {
	Bar string
}

//测试For...Range问题
func TestForRange(){
	list := []int{
 		1,2,3,
	}
	list2 := make([]*int, len(list))
	for i, value := range list {
		x:= &value
		y:= &list[i]
		fmt.Println("value的地址",x)
		fmt.Println("list[i]",y)
		list2[i] = &value
	}
	fmt.Println("----------------------")
	fmt.Println(list[0], list[1], list[2])
	fmt.Println("----------------------")
	fmt.Println(list2[0], list2[1], list2[2])
}

//测试一下，struct的地址和值类型的地址的问题
func TestPoint(){
	value_temp:= 1
	var value = value_temp
	fmt.Println("--------以下是Int类型---------")
	fmt.Println("value_temp",value_temp)
	fmt.Println("value",value)
	fmt.Println("&value_temp",&value_temp)
	fmt.Println("&value",&value)

	fmt.Println("--------以下是自定义结构体---------")
	value_temp2:= Foo{
		"1111",
		}
	var value2 = value_temp2
	fmt.Println("value_temp2",value_temp2)
	fmt.Println("value2",value2)
	fmt.Println("&value_temp2",&value_temp2)
	fmt.Println("&value2",&value2)
}


func ForTest1(){
	for i:=1;i<=10;i++{
		fmt.Printf("i=%d\t",i)
	}
	fmt.Println()
}

func ForTest2(){
	i:=1
	for  ;i<=10; {
		i=i+2
		fmt.Printf("i=%d\t",i)
	}
	fmt.Println()

	//等价于
	for i<=10 {
		i=i+2
		fmt.Printf("i=%d\t",i)
		fmt.Println()
	}
}

func ForTest3(){
	for x,y:=1,10; x<10 && y>1; x,y = x+1,y-1{
		fmt.Printf("x=%d\t",x)
		fmt.Printf("y=%d\t",y)
		fmt.Println()
	}
	fmt.Println()
}

func ForTest4(){
	count:=1
	for {
		fmt.Printf("Hello\t")
		if(count == 3){
			break
		}
		count++
	}
	fmt.Println()
}

func ForTest5(){
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d * %d = %d\t",i,j,i*j)
		}
		fmt.Println()
	}
}
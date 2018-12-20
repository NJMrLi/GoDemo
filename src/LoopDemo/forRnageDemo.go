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
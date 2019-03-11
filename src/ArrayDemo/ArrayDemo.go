package ArrayDemo

import "fmt"

//数组类型是值类型
func ArrayTest1(){
	var arryA [3]int = [3]int{1,2,3}
	//创建一个数组B，将B中第二个元素设置为200
	arryB := arryA
	arryB[1] = 200
	//打印A、B数组的值
	fmt.Printf("arryA[0] = %d\t", arryA[0])
	fmt.Printf("arryB[0] = %d\n", arryB[0])
	fmt.Printf("arryA[1] = %d\t", arryA[1])
	fmt.Printf("arryB[1] = %d\n", arryB[1])
	fmt.Printf("arryA[2] = %d\t", arryA[2])
	fmt.Printf("arryB[2] = %d\n", arryB[2])
}
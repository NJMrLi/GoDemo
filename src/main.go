package main

import (
	"LoopDemo"
	"fmt"
)

func main(){

	//HelloWorld.SayHello()
	//SliceDemo.GetSliceCapAndLength()
	//SliceDemo.GetAppendSlice()
	//LoopDemo.TestForRange()
	//LoopDemo.TestPoint()

	fmt.Println("-----ForTest1-------")
	LoopDemo.ForTest1()
	fmt.Println("-----ForTest2-------")
	LoopDemo.ForTest2()
	fmt.Println("-----ForTest3-------")
	LoopDemo.ForTest3()
	fmt.Println("-----ForTest4-------")
	LoopDemo.ForTest4()
	fmt.Println("-----ForTest5-------")
	LoopDemo.ForTest5()

	//闭包
	//FuncDemo.ClosureDemo1()
	//FuncDemo.ClosureDemo2()
	//FuncDemo.ClosureDemo3()
	//FuncDemo.ClosureDemo4()

	//函数和方法的区别
	//FuncDemo.DiffFuncAndMethod()

	//接口
	//InterfaceDemo.ITableDemo()

	//协程
	//GoroutineDemo.GoRoutineExcute()

	//协程同步，使用原子操作
	//协程加
	//GoroutineDemo.AtomicAdd()
	//协程读取
	//GoroutineDemo.AtomicStoreAndRead()
	//协程Gosched
	//GoroutineDemo.GoschedDemo()

	//无缓冲通道
	//ChannelDemo.PlayTennis()
	//ChannelDemo.Running()

	//缓冲通道
	//ChannelDemo.BufferChannelDemo()

	//使用Runner
	//RunnerDemo.RunnerTest()

	//常量举例
	//IotaDemo.IotaTest()

	//String举例1
	//StringDemo.StringTest1()
	//StringDemo.StringTest2()

	//时间
	//dateTime:=DateTimeDemo.DateTimeBasic()
	//calcDateTime := DateTimeDemo.AddDay(dateTime)
	//timeStamp:=DateTimeDemo.GetDateTimeStamp(calcDateTime)
	//DateTimeDemo.GetDateTime(timeStamp)
	//DateTimeDemo.SimpleTicker()
}
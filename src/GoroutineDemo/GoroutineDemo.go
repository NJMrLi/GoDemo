package GoroutineDemo

import (
	"runtime"
	"sync"
	"fmt"
)

func GoRoutineExcute() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(3)
	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()
		for i:=1;i<=10;i++{
			fmt.Print("1")
		}
	}()

	go func() {
		defer wg.Done()
		for i:=1;i<=10;i++ {
			fmt.Print("2")
		}
	}()

	go func() {
		defer wg.Done()
		for i:=1;i<=10;i++ {
			fmt.Print("3")
		}
	}()

	fmt.Println("等待执行结束")
	wg.Wait()

}
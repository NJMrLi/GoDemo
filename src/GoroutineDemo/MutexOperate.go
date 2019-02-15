package GoroutineDemo

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	num int
	wg3 sync.WaitGroup
	mutex sync.Mutex
)

func MutexOperate() {
	wg3.Add(3)

	go increaseCounter(1)
	go increaseCounter(2)
	go increaseCounter(3)

	wg3.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func increaseCounter(id int) {
	defer wg3.Done()

	for count := 0; count < 2; count++ {
		//互斥锁加上
		mutex.Lock()
		{
			value  := num
			runtime.Gosched()
			value++
			num = value
		}
		//释放锁
		mutex.Unlock()
	}
}
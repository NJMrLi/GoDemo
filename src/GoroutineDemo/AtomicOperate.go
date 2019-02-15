package GoroutineDemo

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
	"time"
)

var (
	counter int64
	wg sync.WaitGroup
)

var (
	shutdown int64
	wg2 sync.WaitGroup
)


func AtomicAdd() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func AtomicStoreAndRead() {

	wg.Add(2)

	//两个工作协程
	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("Shutdown Now")

	//原子写入Shutdown的值为1
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}



func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// 读取shutdown的值，如果为1，则退出
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

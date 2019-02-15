package ChannelDemo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var bufferWg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func BufferChannelDemo() {
	//创建了一个10任务的缓冲通道
	tasks := make(chan string, taskLoad)
	bufferWg.Add(numberGoroutines)

	//创建4个Goroutine
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	//向缓冲通道中放入数据
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	close(tasks)

	bufferWg.Wait()
}

func worker(tasks chan string, worker int) {
	defer bufferWg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : 结束工作 \n", worker)
			return
		}

		fmt.Printf("Worker: %d : 开始工作 %s\n", worker, task)

		//随机处理一下工作的时间
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : 完成工作 %s\n", worker, task)
	}
}
package ChannelDemo

import (
"fmt"
"math/rand"
"sync"
"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func PlayTennis() {
	court := make(chan int)
	wg.Add(2)
	//启动了两个协程，一个纳达尔一个德约科维奇
	go player("纳达尔", court)
	go player("德约科维奇", court)

	//将1放到通道中，模拟开球
	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		// 将数据从通道中取出
		ball, ok := <-court
		if !ok {
			fmt.Printf("选手 %s 胜利\n", name)
			return
		}

		//获取一个随机值，如果可以整除13，就让一个人没有击中，进而关闭整个通道
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("选手 %s 没接到\n", name)
			close(court)
			return
		}
		//如果击中球，就将击球的数量+1，放回通道中
		fmt.Printf("选手 %s 击中 %d\n", name, ball)
		ball++
		court <- ball
	}
}

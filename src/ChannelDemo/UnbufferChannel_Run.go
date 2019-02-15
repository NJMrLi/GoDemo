package ChannelDemo

import (
	"fmt"
	"sync"
	"time"
)

var runnerWg sync.WaitGroup

func Running() {
	//创建一个“接力棒”，也就是通道
	baton := make(chan int)
	runnerWg.Add(1)
	//创建第一个跑步走
	go Runner(baton)
	//开始跑
	baton <- 1
	runnerWg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	//选手接过接力棒
	runner := <-baton
	fmt.Printf("第 %d 选手接棒 \n", runner)

	//如果不是第四名选手，那么说明比赛还在继续
	if runner != 4 {
		//创建一名新选手
		newRunner = runner + 1
		fmt.Printf("第 %d 准备接棒 \n", newRunner)
		go Runner(baton)
	}

	//模拟跑步
	time.Sleep(100 * time.Millisecond)
	//如果第四名跑完了，就结束
	if runner == 4 {
		fmt.Printf("第 %d 结束赛跑 \n", runner)
		runnerWg.Done()
		return
	}

	fmt.Printf("第 %d 选手和第 %d 选手交换了接力棒 \n",
		runner,
		newRunner)

	//选手递出接力棒
	baton <- newRunner
}
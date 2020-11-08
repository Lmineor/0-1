package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func helloWorld() {
	fmt.Println("hello world")
	wg.Done() // 计数器-1
}

func main() { // 开启一个主goroutine去执行main函数

	for i := 0; i < 10000; i++ {
		wg.Add(1)       // 计数牌+1
		go helloWorld() // 开启一个goroutine去执行helloWorld函数
	}

	fmt.Println("Hello main")
	wg.Wait() // 阻塞等所有线程都结束
}

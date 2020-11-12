package main

import "fmt"

/*
两个go routine
1. 生成0-100的数字发送给ch1
2. 从ch1中取出数据并计算他们的平方和并放入ch2中
*/

func f1(ch chan<- int) { // 限定只能往ch中发送值
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 <-chan int, ch2 chan<- int) { // 限定只能从ch1接收值，只能往ch2发送值
	// for tmp := range ch1 {
	// 	ch2 <- tmp * tmp
	// }
	// 从通道中取值的方式1
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)

	go f1(ch1)
	go f2(ch1, ch2)

	for ret := range ch2 {
		fmt.Println(ret)
	}
}

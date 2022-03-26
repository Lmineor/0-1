使用场景：  
我需要完成一项任务，但是这项任务需要满足一定条件才可以执行，否则我就等着。  
那我可以怎么获取这个条件呢？一种是循环去获取，一种是条件满足的时候通知我就可以了。显然第二种效率高很多。  
通知的方式的话，golang里面通知可以用channel的方式

```go
	var mail = make(chan string)
    go func() {
        <- mail
        fmt.Println("get chance to do something")
    }()
    time.Sleep(5*time.Second)
    mail <- "moximoxi"
```
但是channel的方式还是比较适合一对一，一对多并不是很适合。下面就来介绍一下另一种方式：`sync.Cond  `
`sync.Cond`就是用于实现条件变量的，是基于sync.Mutex的基础上，增加了一个通知队列，通知的线程会从通知队列中唤醒一个或多个被通知的线程。  
主要有以下几个方法：
```go
	sync.NewCond(&mutex)：生成一个cond，需要传入一个mutex，因为阻塞等待通知的操作以及通知解除阻塞的操作就是基于sync.Mutex来实现的。
	sync.Wait()：用于等待通知
	sync.Signal()：用于发送单个通知
	sync.Broadcat()：用于广播
```
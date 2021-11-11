package worker

import (
	"fmt"
	"time"
)

type Worker struct {
	JobQueue chan string
	Num      int // 工人的个数
	WorkerCh chan struct{}
	StopCh   chan struct{}
}

func NewWorker(num int) *Worker {
	if num == 0 {
		num = 2
	}
	return &Worker{
		JobQueue: make(chan string, 2),
		Num:      num,
		WorkerCh: make(chan struct{}, num),
		StopCh:   make(chan struct{}),
	}
}

func (w *Worker) Do() {
	w.WorkerCh <- struct{}{}
	w.WorkerCh <- struct{}{} // 先分配两个工人
	for {
		select {
		case <-w.StopCh:
			fmt.Println("worker done!")
			return
		case job := <-w.JobQueue:
			<-w.WorkerCh
			// 如果有可用的工人， 就指派一个下载任务
			go w.Download(job)
		}

	}
}

func (w *Worker) HasJobAndWorker() {

}

func (w *Worker) Stop() {
	w.StopCh <- struct{}{}
}

func (w *Worker) Download(file string) {
	str := fmt.Sprintf("%s start download", file)
	fmt.Println(str)
	time.Sleep(10 * time.Second)
	str = fmt.Sprintf("%s is downloaded", file)
	fmt.Println(str)
	w.WorkerCh <- struct{}{}
}

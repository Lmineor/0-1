package scaner

import (
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/test/printer"
	"github.com/test/set"
	"github.com/test/worker"
)

type Scaner struct {
	TargetDir    string
	ScanInterval int
	BucketLen    int
	Logger       *printer.Printer
	Bucket       *set.Set
	stopCh       chan struct{}
	Lock         sync.RWMutex
}

// NewScaner dir: the directory to scan, interval is the interval(s) to process scan
func NewScaner(dir string, interval int) *Scaner {
	if dir == "" {
		dir, _ = os.Getwd()
	}
	if interval == 0 {
		interval = 5
	}
	return &Scaner{
		TargetDir:    dir,
		ScanInterval: interval,
		Logger:       printer.NewPrinter(),
		Bucket:       set.New(),
		stopCh:       make(chan struct{}),
	}
}

func (sc *Scaner) Stop() {
	sc.stopCh <- struct{}{}
}

func (sc *Scaner) Start(worker *worker.Worker) {
	ticker := time.NewTicker(time.Duration(sc.ScanInterval) * time.Second)
	for {
		select {
		case <-sc.stopCh:
			sc.Logger.PrintInfo("scaner done!")
			return
		case <-ticker.C:
			sc.ScanFolder(worker)
		}
	}
}

func (sc *Scaner) ScanFolder(worker *worker.Worker) {
	sc.Lock.Lock()
	defer sc.Lock.Unlock()
	tmpBucket := set.New()
	fileInfoList, err := ioutil.ReadDir(sc.TargetDir)
	if err != nil {
		sc.Logger.PrintInfo("error")
	}
	tmpBucketLen := 0
	for _, info := range fileInfoList {
		if !info.IsDir() {
			tmpBucket.Add(info.Name())
			tmpBucketLen++
		}
	}
	if sc.BucketLen != tmpBucketLen {
		sc.Logger.PrintWarning("file changed...")
		minux := sc.Bucket.Minus(tmpBucket)
		for _, diff := range minux.List() {
			if sc.Bucket.HasItem(diff) {
				sc.Logger.PrintDelete(diff)
			} else {
				sc.Logger.PrintAdd(diff)
				worker.JobQueue <- diff
			}
		}
		sc.BucketLen = tmpBucketLen
		sc.Bucket = tmpBucket
	} else {
		sc.Logger.PrintInfo("nothing changes, still scaning...")
	}
}

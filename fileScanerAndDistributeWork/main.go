package main

import (
	"time"

	"github.com/test/scaner"
	"github.com/test/worker"
)

type App struct {
	Worker *worker.Worker
	Scaner *scaner.Scaner
}

func StartScanFolder(app *App) {
	scaner := app.Scaner
	worker := app.Worker
	go scaner.Start(worker)
	go worker.Do()

}

func main() {
	worker := worker.NewWorker(2)
	scaner := scaner.NewScaner("E:\\test", 0)
	app := &App{Worker: worker, Scaner: scaner}
	StartScanFolder(app)
	time.Sleep(1 * time.Minute)
	scaner.Stop()
	worker.Stop()
	time.Sleep(1 * time.Second)
}

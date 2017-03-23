package main

import (
	"fmt"
	"sync"
	"time"
)

type MultiWorker struct {
	wg sync.WaitGroup
}

func (w *MultiWorker) Run() {
	for i := 0; i < 100; i++ {
		w.wg.Add(1)
		go w.runJob(i)
	}
	w.wg.Wait()
	fmt.Println("All jobs have finished successfully")
}

func (w *MultiWorker) runJob(id int) {
	defer w.wg.Done()
	defer fmt.Printf("Job %d finished\n", id)

	if (id % 2) == 0 {
		time.Sleep(time.Second)
	} else {
		time.Sleep(2 * time.Second)
	}
}

func main() {
	w := &MultiWorker{}
	w.Run()
}

package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errorCount int32
	var r error
	var runners chan bool = make(chan bool, n)
	wg := &sync.WaitGroup{}

	for _, task := range tasks {
		if atomic.LoadInt32(&errorCount) >= int32(m) { // break if error limit
			r = ErrErrorsLimitExceeded
			break
		}

		runners <- true
		wg.Add(1)
		go func(task Task) {
			err := task()
			if err != nil {
				atomic.AddInt32(&errorCount, 1)
			}
			<-runners
			wg.Done()
		}(task)
	}

	wg.Wait()

	return r
}

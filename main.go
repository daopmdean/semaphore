package main

import (
	"fmt"
	"time"

	"github.com/daopmdean/semaphore/semaphore"
)

func main() {
	testSem()
}

func testSem() {
	sem := semaphore.NewSemaphore(3)
	doneC := make(chan bool, 1)
	total := 10
	for i := 1; i <= total; i++ {
		sem.Acquire()
		go func(i int) {
			defer sem.Release()
			doJob(i)
			if i == total {
				doneC <- true
			}
		}(i)
	}
	<-doneC
}

func doJob(i int) {
	fmt.Println("Doing job---", i)
	time.Sleep(time.Second)
}

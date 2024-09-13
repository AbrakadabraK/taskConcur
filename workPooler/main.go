package main

import (
	"fmt"
	"sync"
)

func main() {
	task := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerID int) {
			fmt.Println("Воркер под номером:", workerID, "стартовал")
			worker(task, &wg, workerID)
		}(i)

	}

	go func() {

		for j := 0; j <= 10; j++ {
			task <- j
		}
		close(task)
	}()

	wg.Wait()

}

func worker(task <-chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for t := range task {
		fmt.Println("Воркер №", i, "встал на исполнение. TaskResult:-->", t)
	}
}

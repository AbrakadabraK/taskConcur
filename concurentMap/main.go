package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	data := make(map[int]string)

	dataLoader := make(chan string)
	intC := make(chan int)

	for j := 0; j < 5; j++ {
		wg.Add(1)
		go getAndParseData(&wg, &mu, dataLoader, intC, data)
	}

	go func() {

		for i := 0; i < 10; i++ {
			r := strconv.Itoa(i)
			s := "И вот у нас пошел уже --> " + r
			dataLoader <- s
			intC <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(dataLoader)
	}()
	wg.Wait()
	fmt.Println(data)
}

func getAndParseData(wg *sync.WaitGroup, mu *sync.Mutex, data chan string, intd chan int, mapppubg map[int]string) {

	defer wg.Done()

	for d := range data {
		mu.Lock()
		mapppubg[<-intd] = d
		mu.Unlock()
	}
}

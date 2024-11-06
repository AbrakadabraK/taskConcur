package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Задача: Программа с таймаутом

Напишите программу, которая запускает длительную операцию в горутине.
Если операция не завершится за определенное время, программа должна завершить выполнение с сообщением о таймауте.

Подсказка:

- Используйте time.After и select, чтобы отслеживать время выполнения операции.
*/
var done = make(chan struct{})

func main() {
	var operationWg sync.WaitGroup
	var timeOut sync.WaitGroup

	operationWg.Add(1)
	go operation(&operationWg)

	timeOut.Add(1)
	go func() {
		defer timeOut.Done()
		select {

		case <-time.After(10 * time.Second):
			fmt.Println("timed out")
			close(done)
		}

	}()

	timeOut.Wait()
	operationWg.Wait()
}

func operation(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Я что то делаю очень долго")
			time.Sleep(1 * time.Second)
		}

	}
}

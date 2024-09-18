/*
7. Задача: Producer-Consumer

Напишите программу, которая симулирует взаимодействие производителя и потребителя.

	Производитель генерирует данные и отправляет их в канал, а потребитель получает данные из канала и обрабатывает их.

Подсказка:

- Используйте буферизированный канал для управления потоком данных.
*/
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	wgProduce := &sync.WaitGroup{}
	wgConsumer := &sync.WaitGroup{}
	task := make(chan string)
	event := make(chan string, 5)

	for j := 0; j < 5; j++ {
		wgProduce.Add(1)
		wgConsumer.Add(1)
		go producer(task, event, wgProduce)
		go consumer(event, wgConsumer)
	}

	go func() {
		for i := 0; i < 10; i++ {

			numb := strconv.Itoa(i)
			task <- "Task № " + numb
		}
		close(task)
	}()
	wgConsumer.Wait()
	wgProduce.Wait()
}

func producer(task chan string, event chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case t, ok := <-task:
			if ok {
				event <- t
			}
		default:
			time.Sleep(30 * time.Millisecond)
		}
	}
}

func consumer(event chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case e := <-event:
			fmt.Println("Event finish : ", e)
		default:
			time.Sleep(30 * time.Millisecond)
		}
	}
}

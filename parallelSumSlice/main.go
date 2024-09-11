package main

import (
	"fmt"
	"time"
)

func main() {
	testData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	res := make(chan []int)
	for i := 0; i < 3; i++ {
		go sumData(res)
	}

	go splitData(testData, res)

	time.Sleep(10 * time.Second)
}

func splitData(data []int, res chan []int) {

	for i := 0; i < len(data); i += 3 {
		end := i + 3
		if end > len(data) {
			end = len(data)
		}
		res <- data[i:end] // Отправляем часть данных длиной 3 элемента (или меньше в конце)
	}
	close(res) // Закрываем канал, когда отправка данных завершена
}

func sumData(data <-chan []int) {
	for d := range data { // Чтение данных из канала до его закрытия
		sum := 0
		for _, s := range d {
			sum += s
		}
		fmt.Println("Сумма:", sum)
	}
}

package main

import (
	"fmt"
	"time"
)

// Функция для игрока "Пинг"
func ping(pingChan, pongChan chan int) {
	for {
		ball := <-pingChan
		if ball > 10 {
			close(pongChan) // Закрываем канал, чтобы завершить игру
			return
		}
		// Ожидаем получения мяча
		fmt.Println("Пинг:", ball)
		time.Sleep(500 * time.Millisecond) // Симулируем задержку

		ball++           // Увеличиваем счётчик передач
		pongChan <- ball // Отправляем мяч игроку "Понг"
	}
}

// Функция для игрока "Понг"
func pong(pingChan, pongChan chan int) {
	for {
		ball, ok := <-pongChan // Ожидаем получения мяча
		if !ok {
			// Если канал закрыт, завершаем игру
			fmt.Println("Понг завершил игру")
			return
		}
		fmt.Println("Понг:", ball)
		time.Sleep(500 * time.Millisecond) // Симулируем задержку

		ball++
		pingChan <- ball // Отправляем мяч обратно игроку "Пинг"
	}
}

func main() {
	pingChan := make(chan int) // Канал для передачи "пинга"
	pongChan := make(chan int) // Канал для передачи "понга"

	// Запускаем горутины "Пинг" и "Понг"
	go ping(pingChan, pongChan)
	go pong(pingChan, pongChan)

	// Инициализируем игру с первого "пинга"
	pingChan <- 1

	// Ожидаем завершения игры
	time.Sleep(6 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string)
	go timerToOneSec(result)
	go timerToFiveSec(result)

	for i := 0; i < 2; i++ {
		result := <-result
		fmt.Println(result)
	}
}

func timerToOneSec(res chan<- string) {
	timerOne := time.NewTimer(1 * time.Second)

	<-timerOne.C

	res <- "Hello"

}

func timerToFiveSec(res chan<- string) {
	timerOne := time.NewTimer(5 * time.Second)
	<-timerOne.C

	res <- "Dear User"

}

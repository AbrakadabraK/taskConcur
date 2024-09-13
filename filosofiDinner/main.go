package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const n int = 5

type Philosopher struct {
	id        int
	leftFork  chan struct{}
	rightFork chan struct{}
}

func NewPhilosopher(id int, leftFork, rightFork chan struct{}) *Philosopher {
	return &Philosopher{
		id:        id,
		leftFork:  leftFork,
		rightFork: rightFork,
	}
}

func (p *Philosopher) eat() {
	fmt.Printf("Философ %d ест\n", p.id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func (p *Philosopher) think() {
	fmt.Printf("Философ %d думает\n", p.id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		p.think()

		if p.id%2 == 0 {
			<-p.leftFork
			<-p.rightFork
		} else {
			<-p.rightFork
			<-p.leftFork
		}

		p.eat()

		if p.id%2 == 0 {
			p.leftFork <- struct{}{}
			p.rightFork <- struct{}{}
		} else {
			p.rightFork <- struct{}{}
			p.leftFork <- struct{}{}
		}
	}
}

func main() {
	wg := sync.WaitGroup{}

	forks := make([]chan struct{}, n)

	for i := 0; i < n; i++ {
		forks[i] = make(chan struct{}, 1)
		forks[i] <- struct{}{}
	}

	philosophers := make([]*Philosopher, n)

	for i := 0; i < n; i++ {
		leftFork := forks[i]
		rightFork := forks[(i+1)%n]
		philosophers[i] = NewPhilosopher(i, leftFork, rightFork)
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
	fmt.Println("Все философы поели и подумали")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

var lock1 sync.Mutex
var lock2 sync.Mutex

func main() {

	var fork1 = Fork{input: make(chan int), output: make(chan int)}
	var fork2 = Fork{input: make(chan int), output: make(chan int)}
	var fork3 = Fork{input: make(chan int), output: make(chan int)}
	var fork4 = Fork{input: make(chan int), output: make(chan int)}
	var fork5 = Fork{input: make(chan int), output: make(chan int)}

	var p1 = Philosopher{leftFork: fork1, rightFork: fork5, thinking: true}
	var p2 = Philosopher{leftFork: fork2, rightFork: fork1, thinking: true}
	var p3 = Philosopher{leftFork: fork3, rightFork: fork2, thinking: true}
	var p4 = Philosopher{leftFork: fork4, rightFork: fork3, thinking: true}
	var p5 = Philosopher{leftFork: fork5, rightFork: fork4, thinking: true}

	philosophers := []Philosopher{p1, p2, p3, p4, p5}

	go p1.eat(p5, p2)
	go p2.eat(p1, p3)
	go p3.eat(p2, p4)
	go p4.eat(p3, p5)
	go p5.eat(p4, p1)

	go fork1.run()
	go fork2.run()
	go fork3.run()
	go fork4.run()
	go fork5.run()

	for {
		display(philosophers)
	}
}

func display(ps []Philosopher) {
	for i, p := range ps {
		var status = ""
		if p.thinking {
			status = "Thinking"
		} else {
			status = "Eating"
		}
		fmt.Printf("\rP%d: %v - %d\n", i, status, p.nrEaten)
		time.Sleep(100 * time.Millisecond)
	}
}

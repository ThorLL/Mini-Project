package main

import (
	"sync"
)

var lock1 sync.Mutex
var lock2 sync.Mutex

var p1 Philosopher
var p2 Philosopher
var p3 Philosopher
var p4 Philosopher
var p5 Philosopher

var fork1 = Fork{input: make(chan int), output: make(chan int)}
var fork2 = Fork{input: make(chan int), output: make(chan int)}
var fork3 = Fork{input: make(chan int), output: make(chan int)}
var fork4 = Fork{input: make(chan int), output: make(chan int)}
var fork5 = Fork{input: make(chan int), output: make(chan int)}

func main() {
	p1 = Philosopher{id: 1, leftFork: &fork1, rightFork: &fork5, thinking: true, leftPhi: &p2, rightPhi: &p5}
	p2 = Philosopher{id: 2, leftFork: &fork2, rightFork: &fork1, thinking: true, leftPhi: &p3, rightPhi: &p1}
	p3 = Philosopher{id: 3, leftFork: &fork3, rightFork: &fork2, thinking: true, leftPhi: &p4, rightPhi: &p2}
	p4 = Philosopher{id: 4, leftFork: &fork4, rightFork: &fork3, thinking: true, leftPhi: &p5, rightPhi: &p3}
	p5 = Philosopher{id: 5, leftFork: &fork5, rightFork: &fork4, thinking: true, leftPhi: &p1, rightPhi: &p4}

	go p1.eat()
	go p2.eat()
	go p3.eat()
	go p4.eat()
	go p5.eat()

	go fork1.run()
	go fork2.run()
	go fork3.run()
	go fork4.run()
	go fork5.run()
	i := 0
	for {
		i++
	}
}

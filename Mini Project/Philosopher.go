package main

type Philosopher struct {
	nrEaten   int
	leftFork  Fork
	rightFork Fork
	leftPhi   *Philosopher
	rightPhi  *Philosopher
	input     chan int
	output    chan bool
}

package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id        int
	nrEaten   int
	thinking  bool
	leftFork  *Fork
	rightFork *Fork
	leftPhi   *Philosopher
	rightPhi  *Philosopher
}

func (p *Philosopher) eat() {
	for {
		if p.leftPhi.thinking && p.rightPhi.thinking {
			lock1.Lock()
			p.thinking = false
			p.leftFork.input <- 1
			p.rightFork.input <- 1
			p.nrEaten++
			p.display()
			<-p.leftFork.output
			<-p.rightFork.output
			p.thinking = true
			p.display()
			lock1.Unlock()
		}
		if p.leftPhi.thinking && p.rightPhi.thinking {
			lock2.Lock()
			p.thinking = false
			p.leftFork.input <- 1
			p.rightFork.input <- 1
			p.nrEaten++
			p.display()
			<-p.leftFork.output
			<-p.rightFork.output
			p.thinking = true
			p.display()
			lock2.Unlock()
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func (p Philosopher) display() {
	var status = ""
	if p.thinking {
		status = "Thinking"
	} else {
		status = "Eating"
	}
	fmt.Printf("\rP%d: %v - %d\n", p.id, status, p.nrEaten)
}

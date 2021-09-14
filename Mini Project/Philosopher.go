package main

type Philosopher struct {
	nrEaten   int
	thinking  bool
	leftFork  Fork
	rightFork Fork
}

func (p *Philosopher) eat(leftPhi Philosopher, rightPhi Philosopher) {
	if leftPhi.thinking && rightPhi.thinking {
		lock1.Lock()
		p.thinking = false
		p.leftFork.input <- 1
		p.rightFork.input <- 1
		p.nrEaten++
		<-p.leftFork.output
		<-p.rightFork.output
		p.thinking = true
		lock1.Unlock()
	}
	if leftPhi.thinking && rightPhi.thinking {
		lock2.Lock()
		p.thinking = false
		p.leftFork.input <- 1
		p.rightFork.input <- 1
		p.nrEaten++
		<-p.leftFork.output
		<-p.rightFork.output
		p.thinking = true
		lock2.Unlock()
	}
}

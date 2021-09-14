package main

type Fork struct {
	nrUsed int
	inUse  bool
	input  chan int
	output chan int
}

func (f *Fork) run() {
	<-f.input
	f.inUse = true
	f.nrUsed++
	f.inUse = false
	f.output <- 1
}

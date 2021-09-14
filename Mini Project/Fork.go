package main

type Fork struct {
	nrUsed int
	input  chan int
	output chan bool
}

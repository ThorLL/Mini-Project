package main

type fork struct {
	nrOfUses int
	beingUsed bool
	input chan fork
	output chan fork
}

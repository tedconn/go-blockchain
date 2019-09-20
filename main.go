package main

import "fmt"

func main() {
	chain := []Block{GenerateGenesisBlock()}

	// Pop the last off the block chain.
	// What you really want to do here is get the chain from storage and get
	// its latest hash from there
	h := chain[len(chain)-1].Hash

	// this is a fake block we add
	// but what would be interesting is to take some input from command line

	minedBlock := AddBlock("Ted Conn", h)
	fmt.Println(minedBlock)

}

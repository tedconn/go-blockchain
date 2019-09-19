package main

func main() {
	chain := []Block{GenerateGenesisBlock()}

	// add a block!
	AddBlock(&chain, "first block name")
}

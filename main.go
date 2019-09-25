package main

import (
	"encoding/json"
	"syscall/js"
)

var done = make(chan struct{}, 0)

func main() {
	c := make(chan bool)
	js.Global().Set("addBlock", js.FuncOf(addBlock))
	js.Global().Set("generateGenesisBlock", js.FuncOf(generateGenesisBlock))
	<-c
}

func addBlock(this js.Value, args []js.Value) interface{} {
	message := args[0].String()
	h := args[1].String()
	block := AddBlock(message, h)

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(block)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

func generateGenesisBlock(this js.Value, args []js.Value) interface{} {
	block := GenerateGenesisBlock()

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(block)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

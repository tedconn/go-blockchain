package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	. "github.com/logrusorgru/aurora"
)

type Block struct {
	Timestamp    time.Time
	Name         string
	Hash         string
	PreviousHash string
	Nonce        int
}

func GenerateGenesisBlock() (block Block) {
	ph := sha256.New()
	ph.Write([]byte("0"))
	previousHash := base64.StdEncoding.EncodeToString(ph.Sum(nil))

	name := "genesis"
	hash := CalculateHash(previousHash, time.Now(), name, 1)

	block.Timestamp = time.Now()
	block.Name = name
	block.Hash = hash
	block.PreviousHash = previousHash
	return
}

func CalculateHash(previousHash string, timestamp time.Time, name string, nonce int) (hash string) {
	h := sha256.New()
	s := fmt.Sprintf("%s%v%s%d", previousHash, timestamp, name, nonce)
	h.Write([]byte(s))
	hash = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}

func CheckDifficulty(difficulty string, hash string) bool {
	//fmt.Println(hash)
	return hash[0:len(difficulty)] == difficulty
}

func NextNonce(block *Block) Block {
	n := block.Nonce + 1
	block.Nonce = n
	return UpdateHash(block)
}

func UpdateHash(block *Block) Block {
	h := CalculateHash(block.PreviousHash, block.Timestamp, block.Name, block.Nonce)
	block.Hash = h
	return *block
}

func MineBlock(block Block) Block {
	start := time.Now()
	newBlock := NextNonce(&block)
	if CheckDifficulty("000", newBlock.Hash) {
		elapsed := time.Since(start)
		ms := float64(elapsed) / float64(time.Millisecond)
		fmt.Println(Sprintf(Bold(Cyan("Mined new block in %fms")), Bold(Cyan(ms))))
		return newBlock
	} else {
		return MineBlock(newBlock)
	}
}

func AddBlock(chain *[]Block, name string) {
	l := len(*chain)
	h := (*chain)[l-1].Hash
	block := Block{
		Name:         name,
		Timestamp:    time.Now(),
		PreviousHash: h,
		Nonce:        0,
	}

	newBlock := MineBlock(block)
	s := append(*chain, newBlock)
	printSlice(s)
}

func printSlice(s []Block) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

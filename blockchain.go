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

// Every chain has to have the first block right? This is called the genesis block
// The previous hash for this block is just a hash of "0"
// and the name is "genesis"
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

// calculate the hash
// the block hash is a sha256 hash of the previousHash, timestamp, name and nonce
func CalculateHash(previousHash string, timestamp time.Time, name string, nonce int) (hash string) {
	h := sha256.New()
	s := fmt.Sprintf("%s%v%s%d", previousHash, timestamp, name, nonce)
	h.Write([]byte(s))
	hash = base64.StdEncoding.EncodeToString(h.Sum(nil))
	return
}

// the difficulty is basically what makes blockchain secure
// when we calculate a hash, we add some salt to it (the nonce)
// and we test that our "difficulty string" is the first n characters
func CheckDifficulty(difficulty string, hash string) bool {
	return hash[0:len(difficulty)] == difficulty
}

func NextNonce(block *Block) Block {
	n := block.Nonce + 1
	block.Nonce = n
	return UpdateHash(block)
}

// Calculate a new hash and set it on the provided block
func UpdateHash(block *Block) Block {
	h := CalculateHash(block.PreviousHash, block.Timestamp, block.Name, block.Nonce)
	block.Hash = h
	return *block
}

// Mine a given block by checking testing the difficulty
// and if not, generate a new hash
func MineBlock(block Block) Block {
	start := time.Now()
	if newBlock := NextNonce(&block); CheckDifficulty("000", newBlock.Hash) {
		elapsed := time.Since(start)
		ms := float64(elapsed) / float64(time.Millisecond)
		fmt.Println(Sprintf(Bold(Cyan("Mined new block in %fms")), Bold(Cyan(ms))))
		return newBlock
	} else {
		return MineBlock(newBlock)
	}
}

// Add a new block to the current chain
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

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

// some dude said online that an individual goroutine
// (depending on what it does obviously) can use about 2.5k per goroutine
// important thing to note is that it's memory bound
var worker_pools = 500

var total_jobs = 10000

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

// need to remove the chain here
// just pass in the previous hash
func MineBlock(id int, jobs <-chan Block, results chan<- Block) {
	for j := range jobs {
		if CheckDifficulty("00", j.Hash) {
			results <- j
		}
	}
}

func AddBlock(name string, previousHash string) (mined Block) {
	start := time.Now()
	block := Block{
		Name:         name, // get from input
		Timestamp:    time.Now(),
		PreviousHash: previousHash,
		Nonce:        0,
	}

	// set up some worker pools and
	// set up some jobs to run in them
	jobs := make(chan Block, 100)
	results := make(chan Block, 100)

	// spawn the workers
	for w := 1; w <= worker_pools; w++ {
		go MineBlock(w, jobs, results)
	}

	for j := 1; j <= total_jobs; j++ {
		jobs <- NextNonce(&block)
	}

	mined = <-results
	close(jobs)

	elapsed := time.Since(start)
	ms := float64(elapsed) / float64(time.Millisecond)
	fmt.Println(Sprintf(Bold(Cyan("Mined new block in %fms")), Bold(Cyan(ms))))

	return
}

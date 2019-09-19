package main

import (
	"testing"
	"time"
)

func TestGenerateGenesisBlock(t *testing.T) {
	genesisBlock := GenerateGenesisBlock()
	if genesisBlock.Name != "genesis" {
		t.Errorf("GenerateGenesisBlock() failed, expected Genesis Block to have name \"%s\", got \"%s\"", "genesis", genesisBlock.Name)
	} else {
		t.Logf("GenerateGenesisBlock() success, expected Genesis Block to have name \"%s\", got \"%s\"", "genesis", genesisBlock.Name)
	}

	_, _, s := time.Now().Clock()
	_, _, blockSeconds := genesisBlock.Timestamp.Clock()
	if blockSeconds != s {
		t.Errorf("GenerateGenesisBlock() failed, expected timestamp seconds to be %d, got %d", s, blockSeconds)
	} else {
		t.Logf("GenerateGenesisBlock() success, expected timestamp seconds to be %d, got %d", s, blockSeconds)
	}
}

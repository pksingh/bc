package main

import (
	"strings"
	"time"
)

// Block will a simple leader record
type Block struct {
	data         map[string]interface{} // Will hold the Raw Data
	hash         string                 //Will have the Hast of current Block
	previousHash string                 //Parent Block Hash
	timestamp    time.Time              //Block created timestamp
	pow          int                    //PoW or nonce to match the Hash Difficulty or Pattern
}

// Blockchain is linked of callection of all ledgers/Blocks
type Blockchain struct {
	genesisBlock Block   // First initiation Block - GENESIS Block
	chain        []Block // List of all Blocks
	difficulty   int     // Difficulty level for PoW/nonce
}
}

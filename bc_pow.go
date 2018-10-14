package main

import (
	"crypto/sha1"
	"fmt"
	"strconv"
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

// Lets Create Blockchain
func CreateBlockchain(difficulty int) Blockchain {
	// Lets initialize 0-block or Genesis Block
	genesisBlock := Block{
		hash:      "0", //No data hence the Hash -"0"
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock}, //Lets the chain have the Gen Block
		difficulty,            // Set the difficuly level
	}
}

// Calculate Hash for currenct Block
func (b Block) calculateHash() string {
	data := fmt.Sprintf("%v", b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha1.Sum([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

}

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
	data         map[string]interface{} // Will hold the Data Map
	hash         string                 // Will have the Hast of current Block
	previousHash string                 // Parent Block Hash
	timestamp    time.Time              // Block created timestamp
	pow          int                    // PoW or nonce to match the Hash Difficulty or Pattern
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

// Lets mine : generate Hash to match the difficulty
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
	fmt.Printf("\tPoW: %d\tHash: %s, Data:%v\n", b.pow, b.hash, b.data)
}

// add the new block to the Blockchain
func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

// validate the existing Blockchain
func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
func main() {
	// create a new blockchain instance with a mining difficulty of 4
	difficulty := 4
	fmt.Println("Init Blockchain with Difficulty ", difficulty)
	blockchain := CreateBlockchain(difficulty)

	// record transactions on the blockchain for Alice, Bob, and John
	fmt.Println("\nAdding Block: ", "Ravin -- 5 --> Binod")
	blockchain.addBlock("Ravin", "Binod", 5)
}

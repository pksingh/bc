package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// BlockChain network or simply lets say Blockchain
type PoSNetwork struct {
	BlockHead  *Block   // Head of all Block
	Blockchain []*Block // List of all Blocks
	Validators []*Node  // All validators, Part of
	// difficulty   int  // Difficulty level for PoW/nonce
}

// Node Structure
type Node struct {
	Id      int    // Id of Validator Node
	Stake   int    // Stake of Node
	Address string // Address of Node
}

// Block will a simple leader record
type Block struct {
	Data          string // Will hold the Data Map
	Hash          string // Will have the Hast of current Block
	PrevHash      string // Parent Block Hash
	Timestamp     int64  // Block creation timestamp
	ValidatorAddr string // Address of Validator Node
	ValidatorId   int    // Added to identify quickly(for human) - Not Required in reality
	// PoW        int    // PoW or nonce to match the Hash Difficulty or Pattern
}

// NewBlockHash : Generate Hash for a NewBlock leager
func NewBlockHash(block *Block) string {
	//data, _ := json.Marshal(block)
	//blockInfo := block.Timestamp + block.PrevHash + block.Hash + block.ValidatorAddr
	blockInfo := strconv.FormatInt(block.Timestamp, 10) + block.PrevHash + block.Hash + block.ValidatorAddr + block.Data
	return newHash(blockInfo)
}

// NewHash() to generate Hash for a string
func newHash(s string) string {
	h := md5.New()
	//	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// randAddress : will generate Adress for Validators
func randAddress() string {
	b := make([]byte, 8) // Lets keep 8Byte Address - user configurable
	_, _ = math.Read(b)
	return fmt.Sprintf("%X", b)
}


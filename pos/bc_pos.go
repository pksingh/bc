package main

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


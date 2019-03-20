package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	math "math/rand"
	"strconv"
	"time"
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

// PrintBlockchainInfo() : Will show whole Blockchain
func (n PoSNetwork) PrintBlockchainInfo() {
	for i, block := range n.Blockchain {
		if i == 0 {
			fmt.Printf("Block GEN> ")
		} else {
			fmt.Printf("Block %d > ", i)
		}
		block.PrintBlockInfo()
	}
}

// PrintBlockDetails() : Display Detail Block Stratucre
func (b Block) PrintBlockDetails() {
	fmt.Println("\tTimeStamp:", b.Timestamp)
	fmt.Println("\tPreHash:", b.PrevHash)
	fmt.Println("\tHash:", b.Hash)
	fmt.Println("\tvAddress:", b.ValidatorAddr)
	fmt.Println("\tvId:", b.ValidatorId)
	fmt.Println("\tData:", b.Data)
}

// PrintBlockInfo() : Display Block Stratucre in Single line
func (b Block) PrintBlockInfo() {
	// fmt.Println("\tTS:", b.Timestamp, " Hash:", b.Hash, " vId:", b.ValidatorId)
	fmt.Println("\tTS:", b.Timestamp, " Hash:", b.Hash, " vId:", b.ValidatorId, " vAddr:", b.ValidatorAddr, "Data:", b.Data)
}

// NewNode() : Add new Validator with initial stake to the BC Network
func (n PoSNetwork) NewNode(id, stake int) []*Node {
	newNode := &Node{
		Id:      id,
		Stake:   stake,
		Address: randAddress(),
	}
	n.Validators = append(n.Validators, newNode)
	return n.Validators
}

// randAddress : will generate Address for Validators
func randAddress() string {
	b := make([]byte, 8) // Lets keep 8Byte Address - user configurable
	_, _ = math.Read(b)
	return fmt.Sprintf("%X", b)
}


// GenerateNewBlock() : will generate block by the validator node
func (n PoSNetwork) GenerateNewBlock(Validator *Node, data string) ([]*Block, *Block, error) {
	if err := n.ValidateBlockchain(); err != nil {
		Validator.Stake -= 10
		return n.Blockchain, n.BlockHead, err
	}

	currentTime := time.Now().UnixMicro() //.String()

	newBlock := &Block{
		Data:          data,
		Timestamp:     currentTime,
		PrevHash:      n.BlockHead.Hash,
		Hash:          NewBlockHash(n.BlockHead),
		ValidatorAddr: Validator.Address,
		ValidatorId:   Validator.Id,
	}

	if err := n.ValidateNewBlock(newBlock); err != nil {
		Validator.Stake -= 10
		return n.Blockchain, n.BlockHead, err
	} else {
		n.Blockchain = append(n.Blockchain, newBlock)
	}
	return n.Blockchain, newBlock, nil
}

// ValidateBlockchain() : It will validate the BC
func (n PoSNetwork) ValidateBlockchain() error {
	if len(n.Blockchain) <= 1 {
		return nil
	}

	currBlockIdx := len(n.Blockchain) - 1
	prevBlockIdx := len(n.Blockchain) - 2

	for prevBlockIdx >= 0 {
		currBlock := n.Blockchain[currBlockIdx]
		prevBlock := n.Blockchain[prevBlockIdx]
		if currBlock.PrevHash != prevBlock.Hash {
			return errors.New("blockchain has inconsistent hashes")
		}

		//if currBlock.Timestamp <= prevBlock.Timestamp {
		//if currBlock.Timestamp < prevBlock.Timestamp {
		// if currBlock.Timestamp.After(prevBlock.Timestamp) {
		// 	return errors.New("blockchain has inconsistent timestamps")
		// }

		if NewBlockHash(prevBlock) != currBlock.Hash {
			return errors.New("blockchain has inconsistent hash generation")
		}
		currBlockIdx--
		prevBlockIdx--
	}
	return nil
}

// ValidateNewBlock() : Will validate the Newly added block
func (n PoSNetwork) ValidateNewBlock(newBlock *Block) error {
	if n.BlockHead.Hash != newBlock.PrevHash {
		return errors.New("blockchain BlockHead(HEAD) hash is not equal to new block previous hash")
	}

	//if n.BlockHead.Timestamp >= newBlock.Timestamp {
	//if n.BlockHead.Timestamp > newBlock.Timestamp {
	// if n.BlockHead.Timestamp.Before(newBlock.Timestamp) {
	// 	return errors.New("blockchain BlockHead(HEAD) timestamp is greater than or equal to new block timestamp")
	// }

	if NewBlockHash(n.BlockHead) != newBlock.Hash {
		return errors.New("new block hash of BlockHead(HEAD) does not equal new block hash")
	}
	return nil
}

// SelectWinner() : One from the Validators will be selected
func (n PoSNetwork) SelectWinner() (*Node, error) {
	var winnerPool []*Node
	totalStake := 0
	for _, node := range n.Validators {
		if node.Stake > 0 {
			winnerPool = append(winnerPool, node)
			totalStake += node.Stake
		}
	}
	if winnerPool == nil {
		return nil, errors.New("there are no nodes with stake in the network")
	}
	winnerNumber := math.Intn(totalStake)
	tmp := 0
	for _, node := range n.Validators {
		tmp += node.Stake
		if winnerNumber < tmp {
			//fmt.Println("\tWinner => Id: ", node.Id, " Address:", node.Address, "-Stake:", node.Stake)
			fmt.Print("\tWinner => Id:", node.Id, " Stake:", node.Stake)
			return node, nil
		}
	}
	return nil, errors.New("a winner should have been picked but wasn't")
}

// AddBlock() : Will Add the Block by Winner Node
func (n PoSNetwork) AddBlock(data string) {
	winner, err := n.SelectWinner()
	if err != nil {
		log.Fatal(err)
	}
	winner.Stake += 10
	n.Blockchain, n.BlockHead, err = n.GenerateNewBlock(winner, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	n.BlockHead.PrintBlockInfo()
	fmt.Println()
}

func main() {
	// set random seed
	math.Seed(time.Now().UnixNano())

	// generate an initial PoS network including a blockchain with a genesis block.
	genesisTime := time.Now().UnixMicro()
	pos := &PoSNetwork{
		Blockchain: []*Block{
			{
				Data:          "",
				Timestamp:     genesisTime,
				PrevHash:      "",
				Hash:          newHash(strconv.FormatInt(genesisTime, 10)),
				ValidatorAddr: "",
				ValidatorId:   0,
			},
		},
	}
	pos.BlockHead = pos.Blockchain[0]
	fmt.Println("Init PoSNetwork DONE")

	// instantiate nodes to act as validators in our network
	pos.Validators = pos.NewNode(1, 20)
	pos.Validators = pos.NewNode(2, 30)
	pos.Validators = pos.NewNode(3, 50)
	fmt.Println("InitValidators => ")
	for _, v := range pos.Validators {
		fmt.Println("\tId: ", v.Id, " Address:", v.Address, " Stake:", v.Stake)
		// fmt.Print("\tvId:", v.Id, " Stake:", v.Stake)
	}

	fmt.Println()
	fmt.Println("\nLets Carry out few Transactions ... ")
	pos.AddBlock("Ravin -- 5 --> Binod")
	pos.AddBlock("Binod -- 3 --> Suresh")
	pos.AddBlock("Suresh -- 1 --> Ravin")

	fmt.Println("\nLets Show/Print all Transactions in the Blockchain ...")
	pos.PrintBlockchainInfo()
}

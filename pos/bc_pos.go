package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
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
}

package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain is an array of pointers
type BlockChain struct {
	Blocks []*Block
}

// Block is a single unit in the blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash will hash the data of the current block, along with the hash from the block preceding it
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock Will add a Block type unit to a blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.PrevHash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis needs to be the first block in a chain, as the first block doesn't have an address to point back to
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain will be what starts a new blockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

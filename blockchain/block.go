package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block is a single unit in the blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Commenting out this code as our new proof of work algorithm takes care of the hashing functionality of the blocks

// DeriveHash will hash the data of the current block, along with the hash from the block preceding it
//func (b *Block) DeriveHash() {
//info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
//hash := sha256.Sum256(info)
//b.Hash = hash[:]
//}
///////////////////////////////////////

// CreateBlock creates a block and performs a proof of work algorithm
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis needs to be the first block in a chain, as the first block doesn't have an address to point back to
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

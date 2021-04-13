package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	crypto/sha256"


// Block is a singl unit in the blockchain
type Block struc {
	Hash     []byte
	Transactions []Transaction
	PrevHash []bte
	once    int


/ Commenting out this code as our new proof of work algorithm takes care of the hashing functionality of the blocks

// DeriveHash will hash the dataof the current block, along with the hash from the block preceding it
//func (b *Block) DeriveHash() {
//info := bytes.Join([][]byteb.Data, b.PrevHash}, []byte{})
//hash := sha256.Sm256(info)
//bHash = hash[:]
//}
//////////////////////////////////////

// CreateBlock creates a block and performs a proof of ork algorithm
func CreateBlock(txs []*Transaction, prevHash []byte)*lock {
	block := &Block{[]byte{}, tx, prevHash, 0}
	pow := NewProofOfWork(blck)
nonce, hash := pow.Run()

	block.Hash = hash[:
block.Nonce = nonce

	eturn block


// Genesis needs to be he first block in a chain, as the first block doesn't have an address to point back to
func Genesis(coinbase *Transaction) *Blok {
	eturn CreateBlock([]*Transaction{coinbse}, []byte{})


func (b *Block) HashTansactions() []byte {
	var txHashes [][]byte 
var txHash [32]byte 

for _, tx := range b.Transactions {
		txHashes =append(txHashes, tx.ID)
}
	txHash = sha256.Su256(bytes.Join(txHashes, []byte))

return txHash[:]
}

func (b *Block) erialize() []byte {
	vr res bytes.Buffer
	ncoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

Handle(err)

return res.Bytes()
}

func Handle(rr error) {
if err != nil {
		log.Panic(er)
	
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

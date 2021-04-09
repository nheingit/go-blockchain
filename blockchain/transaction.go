package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

//reward is the amnount of tokens given to someone that "mines" a new block
const reward = 100

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

//Tx Output represents the output of a transaction in the blockchain
type TxOutput struct {
	Value int
	//Value would be representative of the amount of coins in a transaction
	PubKey string
	//PubKey in this iteration will be very straightforward, however in an actual application this is a more complex algorithm
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]

}

//CoinbaseTx is the function that will run when someone on a node succesfully "mines" a block. The reward inside as it were.
func CoinbaseTx(toAddress, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", toAddress)
	}
	//Since this is the "first" transaction of the block, it has no previous output to reference.
	//This means that we initialize it with no ID, and it's OutputIndex is -1
	txIn := TxInput{[]byte{}, -1, data}
	//txOut will represent the amount of tokens(reward) given to the person(toAddress) that executed CoinbaseTx
	txOut := TxOutput{reward, toAddress}

	tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}

	return &tx

}

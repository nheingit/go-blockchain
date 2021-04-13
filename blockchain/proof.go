package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//The Greater the Difference in the algorithm, the harder the work will be to compute
const Difference = 12 // Possbile TODO: Change this constant to a dynamic algorithim

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difference)) //256 is number of bytes in our hash (SHA256), Lsh is "left shift"

	pow := &ProofOfWork{b, target}

	return pow
}

//InitData takes your block and adds a nonce (counter/incrementer) to it.
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			toHex(int64(nonce)),
			toHex(int64(Difference)),
		},
		[]byte{},
	)
	return data
}

//Validate will check our Run() function performed as expected
func (pow *ProofOfWork) Validate() bool {
	var bigIntHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	bigIntHash.SetBytes(hash[:])

	//this will return true if the hash is valid, and false if not
	return bigIntHash.Cmp(pow.Target) == -1

}

//Run will hash our data, turn that hash into a big int, and then compare that big int to our Target which is inside  the Proof Of Work Struct
func (pow *ProofOfWork) Run() (int, []byte) {
	var bigIntHash big.Int
	var hash [32]byte

	nonce := 0
	//this is effectively an infinite loop as our hashes will not reach math.MaxInt64 before finding the target
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		bigIntHash.SetBytes(hash[:])

		if bigIntHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

//ToHex is a utility function that we will use to cast our nonce into a byte
func toHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

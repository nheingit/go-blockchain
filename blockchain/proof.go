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
			pow.Block.Data,
			toHex(int64(nonce)),
			toHex(int64(Difference)),
		},
		[]byte{},
	)
	return data
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

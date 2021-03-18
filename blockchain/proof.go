package blockchain

import "math/big"

//
const Difference = 12 // Possbile TODO: Change this constant to a dynamic algorithim

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difference)) //256 is number of bytes in our hash (SHA256)

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {

}

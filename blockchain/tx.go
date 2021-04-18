package blockchain

type TxOutput struct {
	Value int
	//Value would be representative of the amount of coins in a transaction
	PubKey string
	//The Pubkey is needed to "unlock" any coins within an Output. This indicated that YOU are the one that sent it.
	//You are indentifiable by your PubKey
	//PubKey in this iteration will be very straightforward, however in an actual application this is a more complex algorithm
}

//Important to note that each output is Indivisible.
//You cannot "make change" with any output.
//If the Value is 10, in order to give someone 5, we need to make two five coin outputs.

//TxInput is represntative of a reference to a previous TxOutput
type TxInput struct {
	ID []byte
	//ID will find the Transaction that a specific output is inside of
	Out int
	//Out will be the index of the specific output we found within a transaction.
	//For example if a transaction has 4 outputs, we can use this "out" field to specify which output we are looking for
	Sig string
	//This would be a script that adds data to an outputs' PubKey
	//however for this tutorial the Sig will be indentical to the PubKey.
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

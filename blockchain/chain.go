package blockchain

type BlockChain struct {
	Blocks []*Block
}

func Genesis() *Block {
	// first block on the chain
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

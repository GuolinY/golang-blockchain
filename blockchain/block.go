package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash, Data, PrevHash []byte
}

func (b *Block) DeriveHash() {
	// todo: placeholder
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
	}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	chain.Blocks = append(chain.Blocks, CreateBlock(data, prevBlock.Hash))
}

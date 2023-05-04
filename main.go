package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PreHash []byte

	Hash []byte

	Data []byte
}

func NewBlock(data string, preHash []byte) *Block {

	b := Block{
		PreHash: preHash,
		Hash:    nil,
		Data:    []byte(data),
	}

	b.setHash()

	return &b

}

func (b *Block)setHash(){

	tmp := [][]byte{
		b.PreHash,
		b.Hash,
		b.Data,
	}

	data := bytes.Join(tmp,[]byte{})
	hash := sha256.Sum256(data)

	b.Hash = hash[:]
}

type BlockChain struct {
	Blocks []*Block
}

const genInfo = "创世之块"

func NewBlockChain() *BlockChain {
	genBlock := NewBlock(genInfo, nil)

	bc := BlockChain{
		Blocks: []*Block{genBlock},
	}

	return &bc
}

func main() {
	bc := NewBlockChain()

	for i, block := range bc.Blocks {
		fmt.Printf("当前区块高度 : %d\n", i)
		fmt.Printf("前哈希： %x\n", block.PreHash)
		fmt.Printf("当前区块哈希：%x\n", block.Hash)
		fmt.Printf("Data:%s\n", string(block.Data))
		
	}

}
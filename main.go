package main

import (
	"crypto/sha256"
	"fmt"
)

//1.定义结构
type Block struct {
	//1.前区块哈希
	PrevHash []byte
	//2.当前区块哈希
	Hash []byte
	//3.数据
	Data []byte
}

//2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}

	block.SetHash()

	return &block
}

//3.生成区块
func (block *Block) SetHash() {
	//1.拼装数据
	blockInfo := append(block.PrevHash, block.Data...)
	//2.sha256
	//func Sum256(data []byte) [Size]byte{}
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//4.引入区块链
type BlockChain struct {
	//定义一个区块链数据
	blocks []*Block
}

//5.定义一个区块链
func NewBlockChain() *BlockChain {
	//创建一个创世块，并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

//创世区块
func GenesisBlock() *Block {
	return NewBlock("创世块，老牛逼了！", []byte{})

}

//6.添加区块
//7.重构代码

func main() {
	bc := NewBlockChain()

	for i, block := range bc.blocks {

		fmt.Printf("======== 当前区块高度 %d========\n", i)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值：%x\n", block.Hash)
		fmt.Printf("区块数据：%s\n", block.Data)
	}
}

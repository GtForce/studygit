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
//5.添加区块
//6.重构代码

func main() {
	block := NewBlock("老师转班长一枚比特币！", []byte{})

	fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
	fmt.Printf("当前区块哈希值：%x\n", block.Hash)
	fmt.Printf("区块数据：%s\n", block.Data)
}

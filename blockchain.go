package main

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
func (bc *BlockChain) AddBlock(data string) {
	//如何获取前区块的哈希
	//获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash
	//a.创建新的区块
	block := NewBlock(data, prevHash)
	//b.添加到区块链数据中
	bc.blocks = append(bc.blocks, block)
}

package main

import (
	"fmt"
	"os"
)

//这是一个用来接收参数并且控制区块链操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA      "add data to blockchain"
	printChain                "print all blockchain data"
`

//接收参数的动作，我们放到一个函数中

func (cli *CLI) Run() {
	// go run .\ printChain
	// go run .\ addBlock --data "HelloWorld"

	//1.得到所有的命令
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(Usage)
		return
	}

	//2.分析命令并执行相应动作
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Printf("添加区块\n")

		//确保命令有效
		if len(args) == 4 && args[2] == "--data" {
			//获取命令的数据
			data := args[3]
			cli.AddBlock(data)
		} else {
			fmt.Printf("添加区块参数使用不当，请检查")
			fmt.Printf(Usage)
		}
	case "printChain":
		fmt.Printf("打印区块\n")
		cli.PrinBlockChain()
	default:
		fmt.Printf("无效的命令，请检查\n")
		fmt.Printf(Usage)
	}
}

package main

import (
	"kongyixueyuan.com/publicChain/part15-persistence-creategenesisblock/BLC"
)

func main()  {

	// 创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	// 新区块
	//blockchain.AddBlockToBlockchain("Send 100RMB To zhangqiang",blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	//
	//blockchain.AddBlockToBlockchain("Send 200RMB To changjingkong",blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	//
	//blockchain.AddBlockToBlockchain("Send 300RMB To juncheng",blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	//
	//blockchain.AddBlockToBlockchain("Send 50RMB To haolin",blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	//
	//
	//fmt.Println(blockchain)
	//fmt.Println(blockchain.Blocks)
}

package main

import (
	"fmt"

	"github.com/MobasirSarkar/go-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to mobasir")
	bc.AddBlock("Send 2 more BTC to sarkar")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Println()
	}
}

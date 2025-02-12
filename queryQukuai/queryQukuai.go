package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	// 实现查询区块功能
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number) // 21827718
	blockNumber := big.NewInt(21828630)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Hash().Hex())          // 0x9374fe4323e14221310e074e76870f66ea9b1e702f7248aba494f5b2b8ccc474
	fmt.Println(block.Number().String())     // 21827724
	fmt.Println(block.Time())                // 1739330507
	fmt.Println(block.Difficulty().String()) // 0
	fmt.Println(len(block.Transactions()))   // 25
	fmt.Println(block.Hash())
	//blockHash := common.HexToHash("0x9374fe4323e14221310e074e76870f66ea9b1e702f7248aba494f5b2b8ccc474")
	//fmt.Println(block.Number())
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

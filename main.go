package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/649a9a3ee6004aeeb63610d4298cc00c")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Ethereum network")

	address := common.HexToAddress("0x01Cb459A55343ee8d75317e2eD2cB51e0679dCDd")
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	fmt.Println(pendingBalance)
}

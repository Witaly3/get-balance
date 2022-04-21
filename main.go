package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"

	token "get-balance/token" // for test
)

func main() {
	client, err := ethclient.Dial("https://speedy-nodes-nyc.moralis.io/2c069930ad94968e935f3447/eth/mainnet")
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	instance, err := token.NewUniswap(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x0d4a11d5EEaaC28EC3F61d100daF4d40471f1852")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		dt := time.Now()
		fmt.Println("Current date and time is: ", dt.String())
		fmt.Printf("name: %s\n", name)
		fmt.Printf("symbol: %s\n", symbol)
		fmt.Printf("wei: %s\n", bal)
		time.Sleep(30 * time.Second)
	}
}

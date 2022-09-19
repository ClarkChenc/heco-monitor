package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	EthClient  *ethclient.Client
	HecoClient *ethclient.Client
)

func init() {
	ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/f9a6daa60f90420997b8ca5adc5dd2db")
	if err != nil {
		fmt.Println(err)
		panic("eth client init failed")
	}
	EthClient = ethClient

	hecoClient, err := ethclient.Dial("https://http-mainnet.hecochain.com")
	if err != nil {
		fmt.Println(err)
		panic("heco client init failed")
	}
	HecoClient = hecoClient
}

package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/x-project/heco/abi/erc20"
	"github.com/x-project/heco/util"
	"github.com/x-project/heco/util/dao"
	"github.com/x-project/heco/util/influx"
)

// service 公共初始化
func init() {
	//db 配置文件加载
	util.InitDBconfig()

	//InitDB
	util.InitDB()

	//InitInfluxDB
	influx.InitInflux2()
}

func OnTick() {
	ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/f9a6daa60f90420997b8ca5adc5dd2db")
	if err != nil {
		fmt.Println(err)
		return
	}
	hecoClient, err := ethclient.Dial("https://http-mainnet.hecochain.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, token := range dao.TokenMap {
		if token.EthAddress != "0x0000000000000000000000000000000000000001" {
			ethContract, _ := erc20.NewErc20(common.HexToAddress(token.EthAddress), ethClient)
			amount, _ := ethContract.BalanceOf(nil, common.HexToAddress("0xa929022c9107643515f5c777ce9a910f0d1e490c"))
			token.EthAmount = decimal.NewFromBigInt(amount, 0)
			decimals, _ := ethContract.Decimals(nil)
			token.EthDecimals = int(decimals)

		} else {
			amount, _ := ethClient.BalanceAt(context.Background(), common.HexToAddress("0xa929022c9107643515f5c777ce9a910f0d1e490c"), nil)
			token.EthAmount = decimal.NewFromBigInt(amount, 0)
			token.EthDecimals = 18
		}

		hecoContract, _ := erc20.NewErc20(common.HexToAddress(token.HecoAddress), hecoClient)
		hecoAmount, _ := hecoContract.TotalSupply(nil)
		token.HecoAmount = decimal.NewFromBigInt(hecoAmount, 0)
		hecoDecimals, _ := hecoContract.Decimals(nil)
		token.HecoDecimals = int(hecoDecimals)

		eAmount := (float64)(hecoAmount.Div(hecoAmount,
			big.NewInt(10).Exp(big.NewInt(10), big.NewInt(int64(token.EthDecimals)-3), nil)).Uint64()) / 1000
		influx.WriteRecord("bridge",
			map[string]string{"chain": "eth", "token": token.TokenName},
			map[string]interface{}{"amount": eAmount})

		hAmount := (float64)(hecoAmount.Div(hecoAmount,
			big.NewInt(10).Exp(big.NewInt(10), big.NewInt(int64(token.HecoDecimals)-3), nil)).Uint64()) / 1000
		influx.WriteRecord("bridge",
			map[string]string{"chain": "heco", "token": token.TokenName},
			map[string]interface{}{"amount": hAmount})

		influx.WriteRecord("bridge",
			map[string]string{"chain": "diff", "token": token.TokenName},
			map[string]interface{}{"amount": eAmount - hAmount})
		if err := token.Update(util.GetDB()); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("update successfully")
}

func main() {
	ctx := context.Background()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			OnTick()

		case <-ctx.Done():
			influx.CloseInflux2()
			fmt.Print("ticker stop")
		}
	}
}

package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/x-project/heco/abi/erc20"
	"github.com/x-project/heco/util"
	"github.com/x-project/heco/util/client"
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

func GetBridgeToken() {
	var wg sync.WaitGroup
	for _, t := range dao.TokenMap {
		wg.Add(1)
		go func(token dao.BridgeToken) {
			defer wg.Done()
			var ethAmount *big.Int
			if token.EthAddress != "0x0000000000000000000000000000000000000001" {
				ethContract, err := erc20.NewErc20(common.HexToAddress(token.EthAddress), client.EthClient)
				if err != nil {
					fmt.Println(err)
					return
				}
				ethAmount, err = ethContract.BalanceOf(nil, common.HexToAddress("0xa929022c9107643515f5c777ce9a910f0d1e490c"))
				if err != nil {
					fmt.Println(err)
					return
				}
				if token.EthAddress == "0xdac17f958d2ee523a2206206994597c13d831ec7" {
					// extra usdt deposited in hub address
					hubAmount, _ := ethContract.BalanceOf(nil, common.HexToAddress("0xbc53b706b165d2b7e98f254095d9d342e845e5ac"))
					ethAmount = ethAmount.Add(ethAmount, hubAmount)
				}

				token.EthAmount = decimal.NewFromBigInt(ethAmount, 0)
				decimals, _ := ethContract.Decimals(nil)
				token.EthDecimals = int(decimals)

			} else {
				ethAmount, _ = client.EthClient.BalanceAt(context.Background(), common.HexToAddress("0xa929022c9107643515f5c777ce9a910f0d1e490c"), nil)
				token.EthAmount = decimal.NewFromBigInt(ethAmount, 0)
				token.EthDecimals = 18
			}

			hecoContract, err := erc20.NewErc20(common.HexToAddress(token.HecoAddress), client.HecoClient)
			if err != nil {
				fmt.Println(err)
				return
			}
			hecoAmount, err := hecoContract.TotalSupply(nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			token.HecoAmount = decimal.NewFromBigInt(hecoAmount, 0)

			hecoDecimals, err := hecoContract.Decimals(nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			token.HecoDecimals = int(hecoDecimals)

			if err := token.Update(util.GetDB()); err != nil {
				fmt.Println(err)
			}

			eAmount := (float64)(ethAmount.Div(ethAmount,
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
		}(t)
	}
	wg.Wait()
	fmt.Println("bridge token update successfully")
}

func GetRootToken() {
	var wg sync.WaitGroup
	for _, root_token := range dao.RootTokenMap {
		wg.Add(1)
		go func(root_token dao.RootToken) {
			defer wg.Done()

			// get token's account balance
			tokenBalance := 0.0
			var accountWg sync.WaitGroup
			var mu sync.Mutex

			for _, accountBalance := range dao.RootAccountMap[root_token.TokenName] {
				accountWg.Add(1)
				go func(accountBalance dao.RootAccountToken) {
					defer accountWg.Done()

					balance, err := client.NewHttpQuery().GetBalance(accountBalance.TokenName, accountBalance.Account)
					if err != nil {
						fmt.Println(err)
						return
					}

					accountBalance.AccountBalance = balance
					err = accountBalance.Update(util.GetDB())
					if err != nil {
						fmt.Println(err)
						return
					}

					mu.Lock()
					tokenBalance += balance
					mu.Unlock()
				}(accountBalance)
			}
			accountWg.Wait()

			// get root token balance
			ethContract, err := erc20.NewErc20(common.HexToAddress(root_token.EthAddress), client.EthClient)
			if err != nil {
				fmt.Println(err)
				return
			}

			tokenAmount, err := ethContract.TotalSupply(nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			root_token.EthAmount = decimal.NewFromBigInt(tokenAmount, 0)

			decimals, err := ethContract.Decimals(nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			root_token.EthDecimals = int(decimals)

			root_token.RootDecimals = int(decimals)
			root_token.RootAmount = decimal.NewFromFloat(tokenBalance * math.Pow10(root_token.RootDecimals))

			if err := root_token.Update(util.GetDB()); err != nil {
				fmt.Println(err)
				return
			}
		}(root_token)
	}

	wg.Wait()
	fmt.Println("root token update successfully")
}

func OnTick() {
	go GetBridgeToken()
	go GetRootToken()
}

func main() {
	ctx := context.Background()
	ticker := time.NewTicker(1 * time.Minute)
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

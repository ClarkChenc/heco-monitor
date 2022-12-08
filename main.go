// package main

// import (
// 	"context"
// 	"fmt"
// 	"math"
// 	"math/big"
// 	"sync"
// 	"time"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/shopspring/decimal"
// 	"github.com/x-project/heco/abi/erc20"
// 	"github.com/x-project/heco/util"
// 	"github.com/x-project/heco/util/client"
// 	"github.com/x-project/heco/util/dao"
// 	"github.com/x-project/heco/util/influx"
// 	"github.com/x-project/heco/util/types"
// )

// // service 公共初始化
// func init() {
// 	//db 配置文件加载
// 	util.InitDBconfig()

// 	//InitDB
// 	util.InitDB()

// 	//InitInfluxDB
// 	influx.InitInflux2()
// }

// func GetBridgeToken() {
// 	var wg sync.WaitGroup
// 	for _, t := range dao.TokenMap {
// 		wg.Add(1)
// 		go func(token dao.BridgeToken) {
// 			defer wg.Done()
// 			var ethAmount *big.Int
// 			if token.EthAddress != "0x0000000000000000000000000000000000000001" {
// 				ethContract, err := erc20.NewErc20(common.HexToAddress(token.EthAddress), client.EthClient)
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}
// 				ethAmount, err = ethContract.BalanceOf(nil, common.HexToAddress("0xa929022c9107643515f5c777ce9a910f0d1e490c"))
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}
// 				if token.EthAddress == "0xdac17f958d2ee523a2206206994597c13d831ec7" {
// 					// extra usdt deposited in hub address
// 					hubAmount, _ := ethContract.BalanceOf(nil, common.HexToAddress("0xbc53b706b165d2b7e98f254095d9d342e845e5ac"))
// 					ethAmount = ethAmount.Add(ethAmount, hubAmount)
// 				}

// 				token.EthAmount = decimal.NewFromBigInt(ethAmount, 0)
// 				decimals, _ := ethContract.Decimals(nil)
// 				token.EthDecimals = int(decimals)

// 			} else {
// 				ethAmount, _ = client.EthClient.BalanceAt(context.Background(), common.HexToAddress("0xa929022c9107643515f5c777ce9a910f0d1e490c"), nil)
// 				token.EthAmount = decimal.NewFromBigInt(ethAmount, 0)
// 				token.EthDecimals = 18
// 			}

// 			hecoContract, err := erc20.NewErc20(common.HexToAddress(token.HecoAddress), client.HecoClient)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			hecoAmount, err := hecoContract.TotalSupply(nil)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			token.HecoAmount = decimal.NewFromBigInt(hecoAmount, 0)

// 			hecoDecimals, err := hecoContract.Decimals(nil)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			token.HecoDecimals = int(hecoDecimals)

// 			if err := token.Update(util.GetDB()); err != nil {
// 				fmt.Println(err)
// 			}

// 			eAmount := (float64)(ethAmount.Div(ethAmount,
// 				big.NewInt(10).Exp(big.NewInt(10), big.NewInt(int64(token.EthDecimals)-3), nil)).Uint64()) / 1000
// 			influx.WriteRecord("bridge",
// 				map[string]string{"chain": "eth", "token": token.TokenName},
// 				map[string]interface{}{"amount": eAmount})

// 			hAmount := (float64)(hecoAmount.Div(hecoAmount,
// 				big.NewInt(10).Exp(big.NewInt(10), big.NewInt(int64(token.HecoDecimals)-3), nil)).Uint64()) / 1000
// 			influx.WriteRecord("bridge",
// 				map[string]string{"chain": "heco", "token": token.TokenName},
// 				map[string]interface{}{"amount": hAmount})

// 			influx.WriteRecord("bridge",
// 				map[string]string{"chain": "diff", "token": token.TokenName},
// 				map[string]interface{}{"amount": eAmount - hAmount})
// 		}(t)
// 	}
// 	wg.Wait()
// 	fmt.Println("bridge token update successfully")
// }

// func GetRootToken() {
// 	for _, root_token := range dao.RootTokenMap {
// 		// get token's account balance
// 		tokenBalance := 0.0
// 		var accountWg sync.WaitGroup

// 		accountToken := &dao.RootAccountToken{TokenName: root_token.TokenName}
// 		accountToken.Clear(util.GetDB())

// 		// fmt.Println("token", root_token)
// 		var mu sync.Mutex
// 		for _, accountBalance := range dao.RootAccountMap[root_token.TokenName] {
// 			accountWg.Add(1)
// 			go func(accountBalance dao.RootAccountToken) {
// 				defer accountWg.Done()

// 				balance, err := client.NewHttpQuery().GetBalance(accountBalance.TokenName, accountBalance.Account)
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}

// 				accountBalance.AccountBalance = balance
// 				err = accountBalance.Update(util.GetDB())
// 				if err != nil {
// 					fmt.Println(err)
// 					return
// 				}

// 				mu.Lock()
// 				tokenBalance += balance
// 				mu.Unlock()
// 			}(accountBalance)
// 		}
// 		accountWg.Wait()

// 		// get root token balance
// 		var tokenContract *erc20.Erc20 = nil

// 		var queryClient *ethclient.Client = nil
// 		switch root_token.SideType {
// 		case types.SIDE_ETH:
// 			queryClient = client.EthClient

// 		case types.SIDE_HECO:
// 			queryClient = client.HecoClient
// 		default:
// 			return
// 		}

// 		tokenContract, err := erc20.NewErc20(common.HexToAddress(root_token.SideAddress), queryClient)
// 		if err != nil || tokenContract == nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		tokenAmount, err := tokenContract.TotalSupply(nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		root_token.SideAmount = decimal.NewFromBigInt(tokenAmount, 0)

// 		decimals, err := tokenContract.Decimals(nil)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		root_token.SideDecimals = int(decimals)

// 		root_token.RootDecimals = int(decimals)
// 		root_token.RootAmount = decimal.NewFromFloat(tokenBalance * math.Pow10(root_token.RootDecimals))

// 		if err := root_token.Update(util.GetDB()); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// 	fmt.Println("root token update successfully")
// }

// func OnTick() {
// 	// go GetBridgeToken()
// 	go GetRootToken()
// }

// func main() {
// 	ctx := context.Background()
// 	interval := 5 * time.Minute

// 	ticker := time.NewTicker(1)
// 	defer ticker.Stop()

// 	isReset := false
// 	for {
// 		select {
// 		case <-ticker.C:
// 			if !isReset {
// 				ticker.Reset(interval)
// 				isReset = true
// 			}
// 			OnTick()

// 		case <-ctx.Done():
// 			influx.CloseInflux2()
// 			fmt.Print("ticker stop")
// 		}
// 	}
// }

package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/x-project/heco/abi/erc20"
)

type EvmToken struct {
	Chain      string   `gorm:"primaryKey;chain:id" json:"-"`
	EthAddress string   `gorm:"column:eth_address" json:"ethAddress"`
	EthAmount  *big.Int `gorm:"column:eth_amount" json:"ethAmount"`
}

func main() {
	//var wg sync.WaitGroup
	evmTokens := make(map[string]EvmToken)
	evmTokens["USDT-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
	}
	evmTokens["USDT-bsc"] = EvmToken{
		Chain:      "bsc",
		EthAddress: "0x55d398326f99059fF775485246999027B3197955",
	}
	evmTokens["BTC-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0x0316EB71485b0Ab14103307bf65a021042c6d380",
	}
	evmTokens["BTC-bsc"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c",
	}
	evmTokens["ETH"] = EvmToken{
		Chain:      "eth",
		EthAddress: "",
	}
	evmTokens["HT-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0x6f259637dcD74C767781E37Bc6133cd6A68aa161",
	}
	evmTokens["USDC-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
	}
	evmTokens["XCN-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0xA2cd3D43c775978A96BdBf12d733D5A1ED94fb18",
	}
	evmTokens["HUSD-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0xdF574c24545E5FfEcb9a659c229253D4111d87e1",
	}
	evmTokens["HUSD-heco"] = EvmToken{
		Chain:      "heco",
		EthAddress: "0x0298c2b32eae4da002a15f36fdf7615bea3da047",
	}
	evmTokens["BETH-heco"] = EvmToken{
		Chain:      "heco",
		EthAddress: "0xb6f4c418514dd4680f76d5caa3bb42db4a893acb",
	}
	evmTokens["NEST-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0x04abEdA201850aC0124161F037Efd70c74ddC74C",
	}
	evmTokens["SHIB-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE",
	}
	evmTokens["SHIB-bsc"] = EvmToken{
		Chain:      "bsc",
		EthAddress: "0x2859e4544C4bB03966803b044A93563Bd2D0DD4D",
	}
	evmTokens["FTT-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0x50D1c9771902476076eCFc8B2A83Ad6b9355a4c9",
	}
	evmTokens["LTC-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0x2c000c0093dE75a8fA2FccD3d97b314e20b431C3",
	}
	evmTokens["LTC-bsc"] = EvmToken{
		Chain:      "bsc",
		EthAddress: "0x4338665CBB7B2485A8855A139b75D5e34AB0DB94",
	}
	evmTokens["ETC-etc"] = EvmToken{
		Chain:      "etc",
		EthAddress: "",
	}
	evmTokens["LINK-eth"] = EvmToken{
		Chain:      "eth",
		EthAddress: "0x514910771AF9Ca656af840dff83E8264EcF986CA",
	}

	chainUrl := make(map[string]*ethclient.Client)
	ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/e094bd4872984c20a288050c1d291598")
	if err != nil {
		fmt.Println(err)
		return
	}
	chainUrl["eth"] = ethClient

	bscClient, err := ethclient.Dial("https://bsc-dataseed1.binance.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	chainUrl["bsc"] = bscClient

	hecoClient, err := ethclient.Dial("https://http-mainnet.hecochain.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	chainUrl["heco"] = hecoClient
	etcClient, err := ethclient.Dial("https://www.ethercluster.com/etc")
	if err != nil {
		fmt.Println(err)
		return
	}
	chainUrl["etc"] = etcClient

	for tokenName, token := range evmTokens {
		if token.EthAddress != "" {
			ethContract, err := erc20.NewErc20(common.HexToAddress(token.EthAddress), chainUrl[token.Chain])
			if err != nil {
				fmt.Println(err)
				return
			}
			token.EthAmount, err = ethContract.BalanceOf(nil, common.HexToAddress("0xc57ae759c085c0d23a9bbf8cd3e3d306a0acf7db"))
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			token.EthAmount, _ = chainUrl[token.Chain].BalanceAt(context.Background(), common.HexToAddress("0xc57ae759c085c0d23a9bbf8cd3e3d306a0acf7db"), nil)
		}
		fmt.Printf("%v, %v\n", tokenName, token.EthAmount)
	}
}

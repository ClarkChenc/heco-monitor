package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type Token struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"-"`
	TokenName string `gorm:"column:token_name" json:"token_name"`

	RootChainAddress  string          `gorm:"column:root_chain_address" json:"root_chain_address"`
	RootChainAmount   decimal.Decimal `gorm:"column:root_chain_amount" json:"root_chain_amount"`
	RootChainDecimals int             `gorm:"column:root_chain_decimals" json:"root_chain_decimals"`

	EthAddress  string          `gorm:"column:eth_address" json:"eth_address"`
	EthAmount   decimal.Decimal `gorm:"column:eth_amount" json:"eth_amount"`
	EthDecimals int             `gorm:"column:eth_decimals" json:"eth_decimals"`

	HecoAddress  string          `gorm:"column:heco_address" json:"heco_address"`
	HecoAmount   decimal.Decimal `gorm:"column:heco_amount" json:"heco_amount"`
	HecoDecimals int             `gorm:"column:heco_decimals" json:"heco_decimals"`
}

var (
	TokenMap map[string]Token
)

func init() {
	TokenMap = make(map[string]Token)

	TokenMap["ETH"] = Token{
		TokenName:   "ETH",
		EthAddress:  "0x0000000000000000000000000000000000000001",
		HecoAddress: "0x64FF637fB478863B7468bc97D30a5bF3A428a1fD",
	}
	TokenMap["UDST"] = Token{
		TokenName: "UDST",

		EthAddress:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
		HecoAddress: "0xa71edc38d189767582c38a3145b5873052c3e47a",
	}
	TokenMap["HT"] = Token{
		TokenName: "HT",

		EthAddress:  "0x6f259637dcd74c767781e37bc6133cd6a68aa161",
		HecoAddress: "0x5545153ccfca01fbd7dd11c0b23ba694d9509a6f",
	}
	TokenMap["HBTC"] = Token{
		TokenName: "HBTC",

		EthAddress:  "0x0316eb71485b0ab14103307bf65a021042c6d380",
		HecoAddress: "0x66a79D23E58475D2738179Ca52cd0b41d73f0BEa",
	}
	TokenMap["MX"] = Token{
		TokenName: "MX",

		EthAddress:  "0x11eef04c884e24d9b7b4760e7476d06ddf797f36",
		HecoAddress: "0x8d854e603dc777337134286f5b3408261736a88F",
	}
	TokenMap["SHIB"] = Token{
		TokenName: "SHIB",

		EthAddress:  "0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce",
		HecoAddress: "0xdd86dd2dc0aca2a8f41a680fc1f88ec1b7fc9b09",
	}
	TokenMap["USDC"] = Token{
		TokenName: "USDC",

		EthAddress:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		HecoAddress: "0x9362Bbef4B8313A8Aa9f0c9808B80577Aa26B73B",
	}
	TokenMap["UNI"] = Token{
		TokenName: "UNI",

		EthAddress:  "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
		HecoAddress: "0x22C54cE8321A4015740eE1109D9cBc25815C46E6",
	}

	TokenMap["HPT"] = Token{
		TokenName: "HPT",

		EthAddress:  "0xa66daa57432024023db65477ba87d4e7f5f95213",
		HecoAddress: "0xe499ef4616993730ced0f31fa2703b92b50bb536",
	}
	TokenMap["Seele"] = Token{
		TokenName: "Seele",

		EthAddress:  "0xb1e93236ab6073fdac58ada5564897177d4bcc43",
		HecoAddress: "0xA9634C25DcEA58dBA4Eb06caE307724b41cda241",
	}
	TokenMap["DAI"] = Token{
		TokenName: "DAI",

		EthAddress:  "0x6b175474e89094c44da98b954eedeac495271d0f",
		HecoAddress: "0x3D760a45D0887DFD89A2F5385a236B29Cb46ED2a",
	}
	TokenMap["LINK"] = Token{
		TokenName: "LINK",

		EthAddress:  "0x514910771af9ca656af840dff83e8264ecf986ca",
		HecoAddress: "0x9e004545c59d359f6b7bfb06a26390b087717b42",
	}
	TokenMap["WOO"] = Token{
		TokenName: "WOO",

		EthAddress:  "0x4691937a7508860f876c9c0a2a617e7d9e945d4b",
		HecoAddress: "0x3befb2308bce92da97264077faf37dcd6c8a75e6",
	}

	TokenMap["NEST"] = Token{
		TokenName: "NEST",

		EthAddress:  "0x04abeda201850ac0124161f037efd70c74ddc74c",
		HecoAddress: "0x2e3443e2ded0c769dd7229dac9ba37155f94fdf1",
	}
	TokenMap["AAVE"] = Token{
		TokenName: "AAVE",

		EthAddress:  "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9",
		HecoAddress: "0x202b4936fE1a82A4965220860aE46d7d3939Bb25",
	}
	TokenMap["BIX"] = Token{
		TokenName: "BIX",

		EthAddress:  "0x009c43b42aefac590c719e971020575974122803",
		HecoAddress: "0x1229e81d6f79b0b49ec686749573307f919f2c71",
	}

	TokenMap["TUSD"] = Token{
		TokenName: "TUSD",

		EthAddress:  "0x0000000000085d4780b73119b644ae5ecd22b376",
		HecoAddress: "0x5ee41ab6edd38cdfb9f6b4e6cf7f75c87e170d98",
	}
	TokenMap["LAMB"] = Token{
		TokenName: "LAMB",

		EthAddress:  "0x8971f9fd7196e5cEE2C1032B50F656855af7Dd26",
		HecoAddress: "0xE131F048D85f0391A24435eEFB7763199B587d0e",
	}

	TokenMap["LAMB"] = Token{
		TokenName: "LAMB",

		EthAddress:  "0x8971f9fd7196e5cEE2C1032B50F656855af7Dd26",
		HecoAddress: "0xE131F048D85f0391A24435eEFB7763199B587d0e",
	}

	TokenMap["GOF"] = Token{
		TokenName: "GOF",

		EthAddress:  "0x488e0369f9bc5c40c002ea7c1fe4fd01a198801c",
		HecoAddress: "0x2AAFe3c9118DB36A20dd4A942b6ff3e78981dce1",
	}

	TokenMap["MATIC"] = Token{
		TokenName: "MATIC",

		EthAddress:  "0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0",
		HecoAddress: "0xdB11743fe8B129b49b11236E8a715004BDabe7e5",
	}

	TokenMap["SWFTC"] = Token{
		TokenName: "SWFTC",

		EthAddress:  "0x0bb217E40F8a5Cb79Adf04E1aAb60E5abd0dfC1e",
		HecoAddress: "0x329dda64Cbc4DFD5FA5072b447B3941CE054ebb3",
	}

	TokenMap["SUSHI"] = Token{
		TokenName: "SUSHI",

		EthAddress:  "0x6b3595068778dd592e39a122f4f5a5cf09c90fe2",
		HecoAddress: "0x52e00b2da5bd7940ffe26b609a42f957f31118d5",
	}

	TokenMap["MANA"] = Token{
		TokenName: "MANA",

		EthAddress:  "0x0f5d2fb29fb7d3cfee444a200298f468908cc942",
		HecoAddress: "0x09006b66d89e5213Fc173403AACBA30620A91F4e",
	}

	TokenMap["ACH"] = Token{
		TokenName: "ACH",

		EthAddress:  "0xed04915c23f00a313a544955524eb7dbd823143d",
		HecoAddress: "0x4a31D1Ad7430586752A1888fE947E3E7D52aFfB8",
	}

	TokenMap["CHZ"] = Token{
		TokenName: "CHZ",

		EthAddress:  "0x3506424f91fd33084466f402d5d97f05f8e3b4af",
		HecoAddress: "0xB37AD4461bB12ba3ac110ed20c3FefE173fa66D3",
	}

	TokenMap["HDOT"] = Token{
		TokenName: "HDOT",

		EthAddress:  "0x9ffc3bcde7b68c46a6dc34f0718009925c1867cb",
		HecoAddress: "0xA2c49cEe16a5E5bDEFDe931107dc1fae9f7773E3",
	}
}

func (t *Token) TableName() string {
	return "token"
}

func (t *Token) Update(db *gorm.DB) error {
	res := db.Where("heco_address=?", t.HecoAddress).Last(t)
	if res.Error == nil || errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res = db.Save(t)
		return res.Error
	}
	return res.Error
}

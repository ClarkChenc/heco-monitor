package dao

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type BridgeToken struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"-"`
	TokenName string `gorm:"column:token_name" json:"token_name"`

	EthAddress  string          `gorm:"column:eth_address" json:"eth_address"`
	EthAmount   decimal.Decimal `gorm:"column:eth_amount" json:"eth_amount"`
	EthDecimals int             `gorm:"column:eth_decimals" json:"eth_decimals"`

	HecoAddress  string          `gorm:"column:heco_address" json:"heco_address"`
	HecoAmount   decimal.Decimal `gorm:"column:heco_amount" json:"heco_amount"`
	HecoDecimals int             `gorm:"column:heco_decimals" json:"heco_decimals"`
}

func (t *BridgeToken) TableName() string {
	return "bridge_token"
}

func (t *BridgeToken) Update(db *gorm.DB) error {
	res := db.Where("heco_address=?", t.HecoAddress).Last(t)
	if res.Error == nil || errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res = db.Save(t)
		return res.Error
	}
	return res.Error
}

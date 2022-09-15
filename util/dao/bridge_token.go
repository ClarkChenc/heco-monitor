package dao

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type BridgeToken struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"-"`
	TokenName string `gorm:"column:token_name" json:"tokenName"`

	EthAddress  string          `gorm:"column:eth_address" json:"ethAddress"`
	EthAmount   decimal.Decimal `gorm:"column:eth_amount" json:"ethAmount"`
	EthDecimals int             `gorm:"column:eth_decimals" json:"ethDecimals"`

	HecoAddress  string          `gorm:"column:heco_address" json:"hecoAddress"`
	HecoAmount   decimal.Decimal `gorm:"column:heco_amount" json:"hecoAmount"`
	HecoDecimals int             `gorm:"column:heco_decimals" json:"hecoDecimals"`

	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (t *BridgeToken) TableName() string {
	return "bridge_token"
}

func (t *BridgeToken) Update(db *gorm.DB) error {
	t.UpdateTime = time.Now()
	newVal := *t
	var oldVal BridgeToken
	res := db.Where("heco_address=?", t.HecoAddress).Last(&oldVal).Update(newVal)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res = db.Save(&newVal)
		return res.Error
	}
	return res.Error
}

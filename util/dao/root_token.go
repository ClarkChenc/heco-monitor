package dao

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type RootToken struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"-"`
	TokenName string `gorm:"column:token_name" json:"tokenName"`

	RootAmount   decimal.Decimal `gorm:"column:root_amount" json:"rootAmount"`
	RootDecimals int             `gorm:"column:root_decimals" json:"rootDecimals"`

	EthAddress  string          `gorm:"column:eth_address" json:"ethAddress"`
	EthAmount   decimal.Decimal `gorm:"column:eth_amount" json:"ethAmount"`
	EthDecimals int             `gorm:"column:eth_decimals" json:"ethDecimals"`

	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (t *RootToken) TableName() string {
	return "root_token"
}

func (t *RootToken) Update(db *gorm.DB) error {
	t.UpdateTime = time.Now()
	newVal := *t
	var oldVal RootToken
	res := db.Where("eth_address=?", t.EthAddress).Last(&oldVal).Update(newVal)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res = db.Save(&newVal)
		return res.Error
	}
	return res.Error
}

type RootAccountToken struct {
	ID        uint   `gorm:"primaryKey;column:id" json:"-"`
	TokenName string `gorm:"column:token_name" json:"tokenName"`

	Account        string  `gorm:"column:account" json:"account"`
	AccountBalance float64 `gorm:"column:account_balance" json:"accountBalance"`

	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (t *RootAccountToken) TableName() string {
	return "root_account_token"
}

func (t *RootAccountToken) Update(db *gorm.DB) error {
	t.UpdateTime = time.Now()
	newVal := *t
	var oldVal RootAccountToken
	res := db.Where("token_name=? and account=?", t.TokenName, t.Account).Last(&oldVal).Update(newVal)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		res = db.Save(&newVal)
		return res.Error
	}
	return res.Error
}

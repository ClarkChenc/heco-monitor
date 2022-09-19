package client

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/x-project/heco/util/types"
)

var (
	urlTypePattern *regexp.Regexp

	blockchairBalancePattern *regexp.Regexp
	filScanBalancePattern    *regexp.Regexp
	xtzBalancePattern        *regexp.Regexp
)

type HttpQuery struct {
	url string
}

func init() {
	urlTypePattern = regexp.MustCompile(`https://(.*?)/`)
	blockchairBalancePattern = regexp.MustCompile(`field="(?:balance|balance_total)"[\s\S]+?>([\d|\.|\,]+?)<`)
	filScanBalancePattern = regexp.MustCompile(`Balance[\s\S]+?> ([\d|\.|\,]+?) FIL`)
	xtzBalancePattern = regexp.MustCompile(`(.*)`)
}

func NewHttpQuery() *HttpQuery {
	return &HttpQuery{}
}

func getHttpContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	return string(body), nil
}

func getBalanceURL(tokenName string, account string) (url string, err error) {
	switch tokenName {
	case types.HBCH:
		url = "https://blockchair.com/" + "bitcoin-cash/" + "address/" + account
	case types.HBSV:
		url = "https://blockchair.com/" + "bitcoin-sv/" + "address/" + account
	case types.HBTC:
		url = "https://blockchair.com/" + "bitcoin/" + "address/" + account
	case types.HDOT:
		url = "https://blockchair.com/" + "polkadot/" + "address/" + account
	case types.HLTC:
		url = "https://blockchair.com/" + "litecoin/" + "address/" + account
	case types.HFIL:
		url = "https://filfox.info/en/address/" + account
	case types.HXTZ:
		url = "https://api.tzkt.io/v1/accounts/" + account + "/balance"
	default:
		err = fmt.Errorf("invalid token", tokenName)
	}
	return
}

func (q *HttpQuery) GetBalance(tokenName string, account string) (float64, error) {
	url, err := getBalanceURL(tokenName, account)
	if err != nil {
		return 0, err
	}
	content, err := getHttpContent(url)
	if err != nil {
		fmt.Println("failed to get http content", err)
		return 0, err
	}

	urlMatch := urlTypePattern.FindStringSubmatch(url)
	if len(urlMatch) == 0 {
		err := fmt.Errorf("invalid url", url)
		return 0, err
	}

	var matches []string
	k := 1.0
	switch urlMatch[1] {
	case "blockchair.com":
		matches = blockchairBalancePattern.FindStringSubmatch(content)

	case "filfox.info":
		matches = filScanBalancePattern.FindStringSubmatch(content)

	case "api.tzkt.io":
		matches = xtzBalancePattern.FindStringSubmatch(content)
		k = math.Pow10(-6)
	}

	if len(matches) < 2 {
		err := fmt.Errorf("failed to get balance", url)
		return 0, err
	}

	strBalance := matches[1]
	strBalance = strings.Replace(strBalance, ",", "", -1)

	balance, err := strconv.ParseFloat(strBalance, 64)
	if err != nil {
		return 0, err
	}
	balance = balance * k
	return balance, nil
}

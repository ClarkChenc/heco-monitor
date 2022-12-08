package client

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/x-project/heco/util/types"
)

var (
	urlTypePattern *regexp.Regexp
	digitPattern   *regexp.Regexp
)

type HttpQuery struct {
	url string
}

func init() {
	urlTypePattern = regexp.MustCompile(`(?:https|http)://(.*?)/`)
	digitPattern = regexp.MustCompile(`([\d|\.|,]+)`)
}

func NewHttpQuery() *HttpQuery {
	return &HttpQuery{}
}

func getHttpHtmlContent(url string, selector string) (string, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("ignore-certificate-errors", true), //忽略错误
		chromedp.Flag("disable-web-security", true),      //禁用网络安全标志
		// chromedp.NoFirstRun,                              //设置网站不是首次运行
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	defer cancel()
	_ = chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 20*time.Second)
	defer cancel()

	var htmlContent string
	sel := `document.querySelector("body")`
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		return "", err
	}

	return htmlContent, nil
}

func getTargetData(htmlContent string, selector string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var str string
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str = selection.Text()
	})
	return str, nil
}

func getBalanceURL(tokenName string, account string) (url string, err error) {
	switch tokenName {
	case types.HXTZ:
		url = "https://tzkt.io/tz1gMsMPNBnq4AqnjrAgG536DP12uVLFgRbt/" + account + "operations/"
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
		url = "https://filscan.io/address/general?address=" + account
	case types.ADA:
		url = "https://blockchair.com/" + "cardano/" + "address/" + account
	case types.XRP:
		url = "https://blockchair.com/" + "ripple/" + "address/" + account
	case types.DOGE:
		url = "https://blockchair.com/" + "dogecoin/" + "address/" + account
	case types.EOS:
		url = "https://bloks.io/account/" + account
	// case types.NFT:
	// 	url = ""

	case types.XCH:
		url = "https://alltheblocks.net/chia/address/" + account

	case types.KAVA:
		url = "https://atomscan.com/kava/accounts/" + account

	case types.ZEC:
		url = "https://blockchair.com/" + "zcash/" + "address/" + account

	case types.ETC:
		url = "https://etcblockexplorer.com/address/" + account

	case types.TT:
		url = "https://v2.viewblock.io/thundercore/address/" + account

	case types.THETA:
		url = "http://www.thetascan.io/address/?address=" + account

	case types.TRX:
		url = "https://tronscan.org/#/address/" + account

	case types.KSM:
		url = "https://blockchair.com/" + "kusama/" + "address/" + account

	case types.LUNA:
		url = "https://atomscan.com/terra/accounts/" + account

	case types.VET:
		url = "https://vechainstats.com/account/" + account

	case types.SOL:
		url = "https://solana.fm/address/" + account

	case types.ELA:
		url = "https://blockchain.elastos.org/address/" + account

	case types.CTXC:
		url = "https://cerebro.cortexlabs.ai/address/" + account

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
	urlMatch := urlTypePattern.FindStringSubmatch(url)
	if len(urlMatch) == 0 {
		err := fmt.Errorf("invalid url", url)
		return 0, err
	}

	var selector string

	switch urlMatch[1] {
	case "blockchair.com":
		selector = `#app > div.body > div > div > div.page-address__balance > div > div > div.account-hash__balance.d-grid > div.account-hash__balance__row--main.d-flex.ai-center.jc-between > span.account-hash__balance__values > span:nth-child(1) > span:nth-child(1) > span.wb-ba`

	case "tzstats.com":
		selector = `#account-info > div:nth-child(1) > span.DataView__Text-sc-1qnax1m-1.kewtyK`

	case "filscout.com":
		selector = `#__layout > div > div > div.contains > div > div > div:nth-child(2) > div.overview.card-radius > ul > li:nth-child(4) > span:nth-child(2)`

	case "filscan.io":
		selector = `#app > div.content-container > div > div > div > div > div.common-top > div.general-overview > div > div:nth-child(4) > div > div > div`

	case "bloks.io":
		selector = `#info > div:nth-child(1) > table > tbody > tr:nth-child(1) > td:nth-child(2)`

	case "alltheblocks.net":
		selector = `#__BVID__197 > tbody > tr > td:nth-child(5) > div`

	case "atomscan.com":
		selector = `#app > div.main-content > div > div > div:nth-child(2) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1) > div > div:nth-child(1) > div.column.is-9`

	case "www.mintscan.io":
		selector = `#__next > main > section > div > div.Account_container__pc9IN > section.Section_container__3OCWW.AccountTokens_container__1-HuH > div > div.TokensTable_container__1LQ7A > div.TokensTable_list__2XPAR > div > div:nth-child(2)`

	case "etcblockexplorer.com":
		selector = `#wrap > div > div:nth-child(2) > div > div > div > div > div:nth-child(2) > p.title.is-5`

	case "scan.thundercore.com":
		selector = `#root > div > div.jss4 > div > div:nth-child(2) > div > div > div:nth-child(2) > div.jss299`

	case "v2.viewblock.io":
		selector = `#__next > div.sc-bczRLJ.lcttxH > div.sc-bczRLJ.sc-38ee1c9c-0.lcttxH.cHvOzI > div > div.sc-bczRLJ.sc-b9475ed6-2.fZnLNI.fTLDkx > div > div.sc-bczRLJ.iIGKPS > span:nth-child(1)`

	case "www.thetascan.io":
		selector = `body > div.main > center > div:nth-child(4) > div:nth-child(10)`

	case "tronscan.org":
		selector = `#root > div.tron-wrapper > main > div > div > div:nth-child(1) > div.info-wrap > div.address-info > div > div.address-overview > div > span.address-asset-num`

	case "vechainstats.com":
		selector = `body > main > div > div.gind-token-area > div > div > div > div:nth-child(1) > div > div:nth-child(2) > div.gind-tkninr-right > p`

	case "solana.fm":
		selector = `#__next > div.main-content.pb-4.bg-blue-800 > div.mx-2.sm\:mx-4.md\:m-0 > div > div.card > div > div > div.flex.justify-between > div.mb-6.text-lg > span > span`

	case "blockchain.elastos.org":
		selector = `#wrap > section > section > div:nth-child(3) > div.row > div.col-md-10 > table > tbody > tr:nth-child(3) > td.ellipsis.text-right.ng-binding`

	case "cerebro.cortexlabs.ai":
		selector = `#addressInfo > tr.ng-scope > td`

	case "tzkt.io":
		selector = `#account > aside > div.v-navigation-drawer__content > div > div:nth-child(3) > div:nth-child(1) > div > div.v-list-item__title`

	}

	var matches []string
	for {
		content, err := getHttpHtmlContent(url, selector)
		if err != nil {
			fmt.Println("failed to get content", "url", url, "err", err)
			continue
		}

		balanceStr, err := getTargetData(content, selector)
		if err != nil {
			fmt.Println("failed to get target data", "url", url, "content", content, "err", err)
			continue
		}

		matches = digitPattern.FindStringSubmatch(balanceStr)
		if len(matches) < 2 {
			err = fmt.Errorf("invalid digit form", "url", url, "content", balanceStr)
			continue
		}
		break
	}

	balanceStr := matches[1]
	balanceStr = strings.Replace(balanceStr, ",", "", -1)

	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		return 0, err
	}
	// fmt.Println("test:", "url", url, "token", tokenName, "account", account, "balance", balance)
	return balance, nil
}

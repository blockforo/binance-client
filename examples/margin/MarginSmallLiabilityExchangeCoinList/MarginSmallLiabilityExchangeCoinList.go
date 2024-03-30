package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginSmallLiabilityExchangeCoinList()
}

func MarginSmallLiabilityExchangeCoinList() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginSmallLiabilityExchangeCoinListService - /sapi/v1/margin/exchange-small-liability
	marginSmallLiabilityExchangeCoinList, err := client.NewMarginSmallLiabilityExchangeCoinListService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginSmallLiabilityExchangeCoinList))
}

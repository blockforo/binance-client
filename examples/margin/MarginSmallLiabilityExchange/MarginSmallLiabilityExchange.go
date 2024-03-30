package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginSmallLiabilityExchange()
}

func MarginSmallLiabilityExchange() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginSmallLiabilityExchangeService - /sapi/v1/margin/exchange-small-liability
	marginSmallLiabilityExchange, err := client.NewMarginSmallLiabilityExchangeService().
		AssetNames("BTC,ETH,BNB").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginSmallLiabilityExchange))
}

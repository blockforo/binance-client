package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginSmallLiabilityExchangeHistory()
}

func MarginSmallLiabilityExchangeHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginSmallLiabilityExchangeHistoryService - /sapi/v1/margin/exchange-small-liability-history
	marginSmallLiabilityExchangeHistory, err := client.NewMarginSmallLiabilityExchangeHistoryService().
		Current(1).Size(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginSmallLiabilityExchangeHistory))
}

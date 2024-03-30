package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginInterestRateHistory()
}

func MarginInterestRateHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginInterestRateHistoryService - /sapi/v1/margin/interestRateHistory
	marginInterestRateHistory, err := client.NewMarginInterestRateHistoryService().Asset("USDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginInterestRateHistory))
}

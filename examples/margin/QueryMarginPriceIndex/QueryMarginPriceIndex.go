package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	QueryMarginPriceIndex()
}

func QueryMarginPriceIndex() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// QueryMarginPriceIndexService - /sapi/v1/margin/priceIndex
	queryMarginPriceIndex, err := client.NewQueryMarginPriceIndexService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(queryMarginPriceIndex))
}

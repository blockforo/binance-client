package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginIsolatedMarginTier()
}

func MarginIsolatedMarginTier() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedMarginTierService - /sapi/v1/margin/isolatedMarginTier
	marginIsolatedMarginTier, err := client.NewMarginIsolatedMarginTierService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginIsolatedMarginTier))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginIsolatedAccountDisable()
}

func MarginIsolatedAccountDisable() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedAccountDisableService - /sapi/v1/margin/isolated/account
	marginIsolatedAccountDisable, err := client.NewMarginIsolatedAccountDisableService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginIsolatedAccountDisable))
}

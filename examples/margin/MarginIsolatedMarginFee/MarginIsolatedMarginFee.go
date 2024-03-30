package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginIsolatedMarginFee()
}

func MarginIsolatedMarginFee() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedMarginFeeService - /sapi/v1/margin/isolatedMarginData
	marginIsolatedMarginFee, err := client.NewMarginIsolatedMarginFeeService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginIsolatedMarginFee))
}

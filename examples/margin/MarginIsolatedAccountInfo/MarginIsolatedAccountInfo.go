package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginIsolatedAccountInfo()
}

func MarginIsolatedAccountInfo() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedAccountInfoService - /sapi/v1/margin/isolated/account
	marginIsolatedAccountInfo, err := client.NewMarginIsolatedAccountInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginIsolatedAccountInfo))
}

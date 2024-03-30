package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginIsolatedAccountLimit()
}

func MarginIsolatedAccountLimit() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedAccountLimitService - /sapi/v1/margin/isolated/accountLimit
	marginIsolatedAccountLimit, err := client.NewMarginIsolatedAccountLimitService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginIsolatedAccountLimit))
}

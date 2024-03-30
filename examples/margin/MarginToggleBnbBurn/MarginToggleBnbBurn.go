package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginToggleBnbBurn()
}

func MarginToggleBnbBurn() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginToggleBnbBurnService - /sapi/v1/bnbBurn
	marginToggleBnbBurn, err := client.NewMarginToggleBnbBurnService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginToggleBnbBurn))
}

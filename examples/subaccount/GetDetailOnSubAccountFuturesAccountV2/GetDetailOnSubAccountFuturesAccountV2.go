package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	GetDetailOnSubAccountFuturesAccountV2()
}

func GetDetailOnSubAccountFuturesAccountV2() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Get Detail on Sub-account's Futures Account V2 (For Master Account) - /sapi/v1/sub-account/futures/internalTransfer
	getDetailOnSubAccountFuturesAccountV2, err := client.NewGetDetailOnSubAccountFuturesAccountV2Service().Email("email@email.com").
		FuturesType(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(getDetailOnSubAccountFuturesAccountV2))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	QuerySubAccountAssets()
}

func QuerySubAccountAssets() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Query Sub-account Assets (For Master Account) - /sapi/v3/sub-account/assets
	querySubAccountAssets, err := client.NewQuerySubAccountAssetsService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(querySubAccountAssets))
}

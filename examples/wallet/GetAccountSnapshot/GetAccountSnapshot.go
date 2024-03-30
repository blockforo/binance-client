package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	AccountSnapshot()
}

func AccountSnapshot() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// GetAccountSnapshotService get all orders from account - /sapi/v1/accountSnapshot
	accountSnapshot, err := client.NewGetAccountSnapshotService().MarketType("SPOT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(accountSnapshot))
}

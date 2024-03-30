package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	QueryManagedSubAccountSnapshot()
}

func QueryManagedSubAccountSnapshot() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	withdrawAssetsFromTheManagedSubAccount, err := client.NewQueryManagedSubAccountSnapshotService().Email("email@email.com").
		SubType("BTC").StartTime(123123123).EndTime(123132123).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(withdrawAssetsFromTheManagedSubAccount))
}

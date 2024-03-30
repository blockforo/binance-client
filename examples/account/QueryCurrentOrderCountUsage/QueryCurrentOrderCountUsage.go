package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	QueryCurrentOrderCountUsage()
}

func QueryCurrentOrderCountUsage() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Query Current Order Count Usage (TRADE)
	getQueryCurrentOrderCountUsageService, err := client.NewGetQueryCurrentOrderCountUsageService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(getQueryCurrentOrderCountUsageService))
}

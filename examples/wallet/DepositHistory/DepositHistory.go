package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	DepositHistory()
}

func DepositHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// DepositHistoryService - /sapi/v1/capital/deposit/hisrec
	depositHistory, err := client.NewDepositHistoryService().Coin("BTC").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(depositHistory))
}

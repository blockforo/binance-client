package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WithdrawHistory()
}

func WithdrawHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// WithdrawHistoryService - /sapi/v1/capital/withdraw/history
	withdrawHistory, err := client.NewWithdrawHistoryService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(withdrawHistory))
}

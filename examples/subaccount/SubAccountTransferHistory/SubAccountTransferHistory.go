package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	SubAccountTransferHistory()
}

func SubAccountTransferHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Sub-account Transfer History (For Sub-account) - /sapi/v1/sub-account/transfer/subUserHistory
	subAccountTransferHistory, err := client.NewSubAccountTransferHistoryService().Asset("BTC").
		TransferType(1).StartTime(1234567891011).EndTime(1234567891011).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(subAccountTransferHistory))
}

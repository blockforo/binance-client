package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	UserUniversalTransferHistory()
}

func UserUniversalTransferHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// UserUniversalTransferHistoryService - /sapi/v1/asset/transfer
	userUniversalTransferHistory, err := client.NewUserUniversalTransferHistoryService().
		TransferType("MAIN_UMFUTURE").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(userUniversalTransferHistory))
}

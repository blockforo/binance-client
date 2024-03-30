package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginIsolatedAccountTransferHistory()
}

func MarginIsolatedAccountTransferHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedAccountTransferHistoryService - /sapi/v1/margin/isolated/transfer
	marginIsolatedAccountTransferHistory, err := client.NewMarginIsolatedAccountTransferHistoryService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginIsolatedAccountTransferHistory))
}

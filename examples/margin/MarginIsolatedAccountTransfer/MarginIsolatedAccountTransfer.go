package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginIsolatedAccountTransfer()
}

func MarginIsolatedAccountTransfer() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginIsolatedAccountTransferService - /sapi/v1/margin/isolated/transfer
	marginIsolatedAccountTransfer, err := client.NewMarginIsolatedAccountTransferService().Asset("USDT").
		Symbol("BTCUSDT").TransFrom("SPOT").TransTo("ISOLATED_MARGIN").Amount(100).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginIsolatedAccountTransfer))
}

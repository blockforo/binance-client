package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginAccountQueryMaxTransferOutAmount()
}

func MarginAccountQueryMaxTransferOutAmount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountQueryMaxTransferOutAmountService - /sapi/v1/margin/maxTransferable
	marginAccountQueryMaxTransferOutAmount, err := client.NewMarginAccountQueryMaxTransferOutAmountService().
		Asset("USDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginAccountQueryMaxTransferOutAmount))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginAccountOrder()
}

func MarginAccountOrder() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountOrderService - /sapi/v1/margin/order
	marginAccountOrder, err := client.NewMarginAccountOrderService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginAccountOrder))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginAccountCancelOrder()
}

func MarginAccountCancelOrder() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountCancelOrderService - /sapi/v1/margin/order
	marginAccountCancelOrder, err := client.NewMarginAccountCancelOrderService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginAccountCancelOrder))
}

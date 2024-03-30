package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginAccountCancelAllOrders()
}

func MarginAccountCancelAllOrders() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountCancelAllOrdersService - /sapi/v1/margin/openOrders
	marginAccountCancelAllOrders, err := client.NewMarginAccountCancelAllOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginAccountCancelAllOrders))
}

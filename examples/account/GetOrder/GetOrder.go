package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	QueryOrder()
}

func QueryOrder() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Binance Query Order (USER_DATA) - GET /api/v3/order
	queryOrder, err := client.NewGetOrderService().Symbol("BTCUSDT").
		OrderId(20064739).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(queryOrder))
}

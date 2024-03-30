package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	GetCurrentOpenOrders()
}

func GetCurrentOpenOrders() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Binance Get current open orders - GET /api/v3/openOrders
	getCurrentOpenOrders, err := client.NewGetOpenOrdersService().Symbol("BTCUSDT").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(getCurrentOpenOrders))
}

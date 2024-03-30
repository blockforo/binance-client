package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	TestNewOrder()
}

func TestNewOrder() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Binance Test New Order endpoint - POST /api/v3/order/test
	testNewOrder, err := client.NewTestNewOrder().Symbol("BTCUSDT").
		Side("BUY").OrderType("MARKET").Quantity(0.001).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(testNewOrder))
}

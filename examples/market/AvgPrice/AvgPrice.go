package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	AvgPrice()
}

func AvgPrice() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// AvgPrice
	avgPrice, err := client.NewAvgPriceService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(avgPrice))
}

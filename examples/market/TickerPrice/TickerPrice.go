package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	TickerPrice()
}

func TickerPrice() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// TickerPrice
	tickerPrice, err := client.NewTickerPriceService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(tickerPrice))
}

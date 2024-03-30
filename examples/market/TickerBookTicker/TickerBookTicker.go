package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	TickerBookTicker()
}

func TickerBookTicker() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// TickerBookTicker
	tickerBookTicker, err := client.NewTickerBookTickerService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(tickerBookTicker))
}

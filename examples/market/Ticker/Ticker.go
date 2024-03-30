package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	Ticker()
}

func Ticker() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// Ticker
	ticker, err := client.NewTickerService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(ticker))
}

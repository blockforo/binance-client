package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	Ticker24hr()
}

func Ticker24hr() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// Ticker24hr
	ticker24hr, err := client.NewTicker24hrService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(ticker24hr))
}

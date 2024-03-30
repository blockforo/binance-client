package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	UiKlines()
}

func UiKlines() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// UiKlines
	uiKlines, err := client.NewUIKlinesService().
		Symbol("BTCUSDT").Interval("1m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(uiKlines))
}

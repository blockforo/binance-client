package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	Klines()
}

func Klines() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// Klines
	klines, err := client.NewKlinesService().
		Symbol("BTCUSDT").Interval("1m").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(klines))
}

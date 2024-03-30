package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	HistoricalTradeLookup()
}

func HistoricalTradeLookup() {
	apiKey := "your api key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, "", baseURL)

	historicalTradeLookup, err := client.NewHistoricalTradeLookupService().
		Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(historicalTradeLookup))
}

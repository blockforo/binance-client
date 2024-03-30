package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	AggTradesList()
}

func AggTradesList() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// AggTradesList
	aggTradesList, err := client.NewAggTradesListService().
		Symbol("BTCUSDT").Limit(20).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(aggTradesList))
}

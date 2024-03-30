package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	ExchangeInfo()
}

func ExchangeInfo() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// ExchangeInfo
	exchangeInfo, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(exchangeInfo))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginCurrentOrderCount()
}

func MarginCurrentOrderCount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginCurrentOrderCountService - /sapi/v1/margin/rateLimit/order
	marginCurrentOrderCount, err := client.NewMarginCurrentOrderCountService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginCurrentOrderCount))
}

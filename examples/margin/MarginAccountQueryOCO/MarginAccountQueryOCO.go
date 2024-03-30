package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginAccountQueryOCO()
}

func MarginAccountQueryOCO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountQueryOCOService - /sapi/v1/margin/orderList
	marginAccountQueryOCO, err := client.NewMarginAccountQueryOCOService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginAccountQueryOCO))
}

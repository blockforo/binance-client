package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginAccountQueryAllOCO()
}

func MarginAccountQueryAllOCO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginAccountQueryAllOCOService - /sapi/v1/margin/allOrderList
	marginAccountQueryAllOCO, err := client.NewMarginAccountQueryAllOCOService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginAccountQueryAllOCO))
}

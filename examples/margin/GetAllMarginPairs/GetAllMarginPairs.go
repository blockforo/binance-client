package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	GetAllMarginPairs()
}

func GetAllMarginPairs() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// GetAllMarginPairsService - /sapi/v1/margin/allPairs
	getAllMarginPairs, err := client.NewGetAllMarginPairsService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(getAllMarginPairs))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	QueryAllOCO()
}

func QueryAllOCO() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Binance Query all OCO (USER_DATA) (GET /api/v3/allOrderList)
	queryAllOCO, err := client.NewQueryAllOCOService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(queryAllOCO))
}

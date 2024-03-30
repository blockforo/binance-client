package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	GetSubAccountStatus()
}

func GetSubAccountStatus() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Get Sub-account's Status on Margin/Futures (For Master Account) - /sapi/v1/sub-account/status
	getSubAccountStatus, err := client.NewGetSubAccountStatusService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(getSubAccountStatus))
}

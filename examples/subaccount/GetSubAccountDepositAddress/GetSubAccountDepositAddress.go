package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	GetSubAccountDepositAddress()
}

func GetSubAccountDepositAddress() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Get Sub-account Deposit Address (For Master Account) - /sapi/v1/capital/deposit/subAddress
	getSubAccountDepositAddress, err := client.NewGetSubAccountDepositAddressService().Email("from@email.com").
		Coin("BTC").Network("BTC").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(getSubAccountDepositAddress))
}

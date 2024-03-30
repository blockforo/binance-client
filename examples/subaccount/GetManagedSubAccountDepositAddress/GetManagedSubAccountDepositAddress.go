package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	GetManagedSubAccountDepositAddress()
}

func GetManagedSubAccountDepositAddress() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Get Managed Sub-account Deposit Address (For Investor Master Account) - /sapi/v1/managed-subaccount/deposit/address
	GetManagedSubAccountDepositAddress, err := client.NewGetManagedSubAccountDepositAddressService().Email("from@email.com").
		Coin("BTC").Network("BTC").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(GetManagedSubAccountDepositAddress))
}

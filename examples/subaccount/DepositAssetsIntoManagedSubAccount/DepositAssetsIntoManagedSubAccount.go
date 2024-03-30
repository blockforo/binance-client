package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	DepositAssetsIntoManagedSubAccount()
}

func DepositAssetsIntoManagedSubAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Deposit Assets Into The Managed Sub-account（For Investor Master Account） - /sapi/v1/sub-account/managed-subaccount/deposit
	depositAssetsIntoManagedSubAccount, err := client.NewDepositAssetsIntoManagedSubAccountService().ToEmail("to@email.com").
		Asset("BTC").Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(depositAssetsIntoManagedSubAccount))
}

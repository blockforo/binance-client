package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WithdrawAssetsFromTheManagedSubAccount()
}

func WithdrawAssetsFromTheManagedSubAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	withdrawAssetsFromTheManagedSubAccount, err := client.NewWithdrawAssetsFromTheManagedSubAccountService().FromEmail("email@email.com").
		Asset("BTC").Amount(1.5).TransferDate(123132123).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(withdrawAssetsFromTheManagedSubAccount))
}

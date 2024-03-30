package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	EnableMarginForSubAccount()
}

func EnableMarginForSubAccount() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Enable Margin for Sub-account (For Master Account) - /sapi/v1/sub-account/margin/enable
	enableMarginForSubAccount, err := client.NewEnableMarginForSubAccountService().Email("from@email.com").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(enableMarginForSubAccount))
}

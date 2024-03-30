package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	DisableFastWithdrawSwitch()
}

func DisableFastWithdrawSwitch() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// DisableFastWithdrawSwitchService  - /sapi/v1/account/disableFastWithdrawSwitch
	disableFastWithdrawSwitch, err := client.NewDisableFastWithdrawSwitchService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(disableFastWithdrawSwitch))
}

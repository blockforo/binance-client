package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	PersonalLeftQuota()
}

func PersonalLeftQuota() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	personalLeftQuota, err := client.NewPersonalLeftQuotaService().Product("STAKING").
		ProductId("AXS*90").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(personalLeftQuota))
}

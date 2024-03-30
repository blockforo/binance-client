package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	UpdateIPRestrictionForSubAccountAPIKey()
}

func UpdateIPRestrictionForSubAccountAPIKey() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Update IP Restriction for Sub-Account API key (For Master Account) - /sapi/v2/sub-account/subaccountApi/ipRestriction
	updateIPRestrictionForSubAccountAPIKey, err := client.NewUpdateIPRestrictionForSubAccountAPIKeyService().Email("email@email.com").
		SubAccountApiKey("123123").Status("").IpAddress("127.0.0.1").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(updateIPRestrictionForSubAccountAPIKey))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	RedeemStakingProduct()
}

func RedeemStakingProduct() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	redeemStakingProduct, err := client.NewRedeemStakingProductService().Product("STAKING").
		ProductId("AXS*90").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(redeemStakingProduct))
}

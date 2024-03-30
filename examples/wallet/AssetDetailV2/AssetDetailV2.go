package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	AssetDetailV2()
}

func AssetDetailV2() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// AssetDetailV2Service - /sapi/v1/asset/assetDetail
	assetDetailV2, err := client.NewAssetDetailV2Service().Asset("BTC").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(assetDetailV2))
}

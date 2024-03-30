package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	AssetDividendRecord()
}

func AssetDividendRecord() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// AssetDividendRecordService - /sapi/v1/asset/assetDividend
	assetDividendRecord, err := client.NewAssetDividendRecordService().Asset("BTC").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(assetDividendRecord))
}

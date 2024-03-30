package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	ForceLiquidationRecord()
}

func ForceLiquidationRecord() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// ForceLiquidationRecordService - /sapi/v1/margin/forceLiquidationRec
	forceLiquidationRecord, err := client.NewForceLiquidationRecordService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(forceLiquidationRecord))
}

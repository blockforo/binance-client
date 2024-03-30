package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	TransferToMaster()
}

func TransferToMaster() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Transfer to Master (For Sub-account) - /sapi/v1/sub-account/transfer/subToMaster
	transferToMaster, err := client.NewTransferToMasterService().Asset("BTC").
		Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(transferToMaster))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	TransferToSubAccountOfSameMaster()
}

func TransferToSubAccountOfSameMaster() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Transfer to Sub-account of Same Master (For Sub-account) - /sapi/v1/sub-account/transfer/subToSub
	transferToSubAccountOfSameMaster, err := client.NewTransferToSubAccountOfSameMasterService().ToEmail("from@email.com").Asset("BTC").
		Amount(0.01).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(transferToSubAccountOfSameMaster))
}

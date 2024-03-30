package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	DustTransfer()
}

func DustTransfer() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// DustTransferService - /sapi/v1/asset/dust
	dustTransfer, err := client.NewDustTransferService().Asset([]string{"ETH", "LTC", "TRX"}).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(dustTransfer))
}

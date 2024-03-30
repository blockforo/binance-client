package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	SubAccountFuturesAssetTransferHistory()
}

func SubAccountFuturesAssetTransferHistory() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Query Sub-account Futures Asset Transfer History (For Master Account) - /sapi/v1/sub-account/futures/internalTransfer
	subaccountFuturesAssetTransferHistory, err := client.NewQuerySubAccountFuturesAssetTransferHistoryService().Email("from@email.com").
		FuturesType(1).StartTime(1234567891011).EndTime(1).Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(subaccountFuturesAssetTransferHistory))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	QuerySubAccountTransactionTatistics()
}

func QuerySubAccountTransactionTatistics() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	transactionTatistics, err := client.NewQuerySubAccountTransactionTatistics().Email("email@email.com").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(transactionTatistics))
}

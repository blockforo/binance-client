package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginCrossMarginFee()
}

func MarginCrossMarginFee() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginCrossMarginFeeService - /sapi/v1/margin/crossMarginData
	marginCrossMarginFee, err := client.NewMarginCrossMarginFeeService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginCrossMarginFee))
}

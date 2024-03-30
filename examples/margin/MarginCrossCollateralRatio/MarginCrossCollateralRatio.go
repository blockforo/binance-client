package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	MarginCrossCollateralRatio()
}

func MarginCrossCollateralRatio() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// MarginCrossCollateralRatioService - /sapi/v1/margin/crossMarginCollateralRatio
	marginCrossCollateralRatio, err := client.NewMarginCrossCollateralRatioService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(marginCrossCollateralRatio))
}

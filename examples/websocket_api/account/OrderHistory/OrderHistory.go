package main

import (
	"context"
	"fmt"
	"log"

	binance "github.com/blockforo/binance-client"
)

func main() {
	OrderHistoryExample()
}

func OrderHistoryExample() {
	client := binance.NewWebsocketAPIClient("api_key", "secret_key")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewAccountOrderHistoryService().Symbol("BTCUSDT").StartTime(1684566357000).Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance.PrettyPrint(response))

	client.WaitForCloseSignal()
}

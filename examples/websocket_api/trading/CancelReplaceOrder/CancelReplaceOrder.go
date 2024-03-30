package main

import (
	"context"
	"fmt"
	"log"

	binance "github.com/blockforo/binance-client"
)

func main() {
	CancelReplaceOrderExample()
}

func CancelReplaceOrderExample() {
	client := binance.NewWebsocketAPIClient("api_key", "secret_key", "wss://testnet.binance.vision/ws-api/v3")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewCancelReplaceOrderService().Symbol("BTCUSDT").CancelReplaceMode("STOP_ON_FAILURE").
		Side("BUY").OrderType("MARKET").CancelOrderID(123123132).Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance.PrettyPrint(response))

	client.WaitForCloseSignal()
}

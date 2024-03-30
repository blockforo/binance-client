package main

import (
	"context"
	"fmt"
	"log"

	binance "github.com/blockforo/binance-client"
)

func main() {
	StartUserDataStream()
}

func StartUserDataStream() {
	client := binance.NewWebsocketAPIClient("API_KEY", "")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	response, err := client.NewStartUserDataStreamService().Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance.PrettyPrint(response.Result.ListenKey))

	client.WaitForCloseSignal()
}

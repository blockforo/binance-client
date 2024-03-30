package main

import (
	"context"
	"fmt"
	"log"

	binance "github.com/blockforo/binance-client"
)

func main() {
	PingUserDataStream()
}

func PingUserDataStream() {
	client := binance.NewWebsocketAPIClient("API_KEY", "")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewPingUserDataStreamService().ListenKey("LISTEN_KEY").Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance.PrettyPrint(response.Response))

	client.WaitForCloseSignal()
}

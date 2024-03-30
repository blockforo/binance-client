package main

import (
	"context"
	"fmt"
	"log"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WsStreamAndWsAPI()
}

func WsStreamAndWsAPI() {
	// Websocket Stream
	websocketStreamClient := binance.NewWebsocketStreamClient(false)
	wsTradeHandler := func(event *binance.WsTradeEvent) {
		fmt.Println(binance.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsTradeServe("LTCBTC", wsTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Websocket API
	client := binance.NewWebsocketAPIClient("api_key", "secret_key", "wss://testnet.binance.vision/ws-api/v3")
	err2 := client.Connect()
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return
	}

	response, err := client.NewPlaceNewOrderService().Symbol("BTCUSDT").Side("BUY").OrderType("MARKET").Quantity(0.01).
		Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance.PrettyPrint(response))

	<-doneCh
	client.WaitForCloseSignal()
}

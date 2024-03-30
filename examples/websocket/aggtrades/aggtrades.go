package main

import (
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	AggTradesExample()
}

func AggTradesExample() {
	// Initialise Websocket Client with Testnet base URL and false for "isCombined" parameter
	websocketStreamClient := binance.NewWebsocketStreamClient(false, "wss://testnet.binance.vision")

	wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
		fmt.Println(binance.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsAggTradeServe("BTCUSDT", wsAggTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}

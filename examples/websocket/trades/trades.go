package main

import (
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WsTradeExample()
}

func WsTradeExample() {
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
	<-doneCh
}

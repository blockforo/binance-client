package main

import (
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WsKlineExample()
}

func WsKlineExample() {
	websocketStreamClient := binance.NewWebsocketStreamClient(false)
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println(binance.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsKlineServe("LTCBTC", "1m", wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}

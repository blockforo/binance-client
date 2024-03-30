package main

import (
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WsBookTickerExample()
}

func WsBookTickerExample() {
	websocketStreamClient := binance.NewWebsocketStreamClient(false)
	wsBookTickerHandler := func(event *binance.WsBookTickerEvent) {
		fmt.Println(binance.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsBookTickerServe("LTCBTC", wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}

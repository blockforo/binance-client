package main

import (
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WsAllMarketTickersExample()
}

func WsAllMarketTickersExample() {
	websocketStreamClient := binance.NewWebsocketStreamClient(false)
	wsAllMarketTickersHandler := func(event binance.WsAllMarketTickersStatEvent) {
		fmt.Println(binance.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsAllMarketTickersStatServe(wsAllMarketTickersHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}

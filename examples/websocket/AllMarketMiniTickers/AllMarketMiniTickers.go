package main

import (
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WsAllMarketMiniTickers()
}

func WsAllMarketMiniTickers() {
	websocketStreamClient := binance.NewWebsocketStreamClient(false)
	wsAllMarketMiniTickersHandler := func(event binance.WsAllMarketMiniTickersStatEvent) {
		fmt.Println(binance.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, _, err := websocketStreamClient.WsAllMarketMiniTickersStatServe(wsAllMarketMiniTickersHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneCh
}

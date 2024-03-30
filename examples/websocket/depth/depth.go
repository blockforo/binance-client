package main

import (
	"fmt"
	"time"

	binance "github.com/blockforo/binance-client"
)

func main() {
	WsDepthHandlerExample()
}

func WsDepthHandlerExample() {
	websocketStreamClient := binance.NewWebsocketStreamClient(false)
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		fmt.Println(binance.PrettyPrint(event))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneCh, stopCh, err := websocketStreamClient.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopCh to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopCh <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneCh
}

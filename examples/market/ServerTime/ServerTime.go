package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	ServerTime()
}

func ServerTime() {

	client := binance.NewClient("", "")

	// set to debug mode
	client.Debug = true

	// NewServerTimeService
	serverTime, err := client.NewServerTimeService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(serverTime))
}

package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	Ping()
}

func Ping() {
	baseURL := "https://api.binance.com"

	client := binance.NewClient("", "", baseURL)

	// NewPingService
	ping := client.NewPingService().Do(context.Background())
	if ping == nil {
		fmt.Println("Success")
		return
	}
	fmt.Println(binance.PrettyPrint(ping))
}

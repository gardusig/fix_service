package main

import (
	"fmt"

	"github.com/gardusig/fix_service/protocol/fix"
)

const fixSettingsFilepath = "/Users/gardusig/github/fix_service/order_sender/config/fix.cfg"

func main() {
	fmt.Println("Starting algo-engine...")
	fixClient, err := fix.NewClientFIX(fixSettingsFilepath)
	if err != nil {
		panic(err)
	}
	fixClient.Start()
	defer fixClient.Stop()
	fmt.Println("Started algo-engine")
	select {}
}

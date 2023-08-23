package main

import (
	"fmt"

	"fix_service/protocol/fix"
)

const fixSettingsFilepath = "/Users/gardusig/github/fix_service/order_sender/config/fix.cfg"

func main() {
	fmt.Println("Starting algo-engine...")
	fixClient := fix.NewClientFIX(fixSettingsFilepath)
	fixClient.Start()
	defer fixClient.Stop()
	fmt.Println("Started algo-engine")
	select {}
}

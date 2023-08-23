package main

import (
	"fmt"
)

const fixSettingsFilepath = "/Users/gardusig/beyond/trading-bot/matching_engine/config/fix.cfg"

func main() {
	fmt.Println("Starting algo-engine...")
	fixClient := fix.NewClientFIX(fixSettingsFilepath)
	fixClient.Start()
	defer fixClient.Stop()
	fmt.Println("Started algo-engine")
	select {}
}

package main

import (
	"fmt"
	"path"

	"github.com/gardusig/fix_service/protocol/fix"
)

func main() {
	fmt.Println("Starting client...")
	client, err := fix.NewClientFIX(path.Join("config", "fix.cfg"))
	if err != nil {
		panic(err)
	}
	client.Start()
	defer client.Stop()
	fmt.Println("Started client")
	select {}
}

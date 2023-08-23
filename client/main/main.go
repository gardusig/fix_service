package main

import (
	"fmt"
	"path"
	"time"

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
	time.Sleep(5 * time.Second)
}

package main

import (
	"fmt"
	"path"

	"github.com/gardusig/fix_service/protocol/fix"
)

func main() {
	fmt.Println("Starting server...")
	server, err := fix.NewServerFIX(path.Join("config", "fix.cfg"))
	if err != nil {
		panic(err)
	}
	server.Start()
	defer server.Stop()
	fmt.Println("Started server")
	select {}
}

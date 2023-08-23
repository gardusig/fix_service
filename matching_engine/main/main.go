package main

import (
	"fmt"
	"matching_engine/server"
	"os"
	"path"

	"github.com/quickfixgo/quickfix"
)

var (
	settings     *quickfix.Settings
	serverApp    server.ServerApp
	storeFactory quickfix.MessageStoreFactory
	logFactory   quickfix.LogFactory
)

func init() {
	serverApp = server.ServerApp{}
	storeFactory = quickfix.NewMemoryStoreFactory()
	logFactory = quickfix.NewScreenLogFactory()
	settings = quickfix.NewSettings()
	fmt.Println("init done")
}

func startServer() {
	fmt.Println("Starting the matching engine...")
	acceptor, err := quickfix.NewAcceptor(
		serverApp,
		storeFactory,
		settings,
		logFactory,
	)
	if err != nil {
		fmt.Println("Failed to start server, reason:", err)
		panic(err)
	}
	acceptor.Start()
	defer acceptor.Stop()
	fmt.Println("Started the matching engine")
	select {}
}

func main() {
	startServer()
	cfgFileName := path.Join("matching_engine", "config", "fix.cfg")
	cfg, err := os.Open(cfgFileName)
	if err != nil {
		panic(err)
		return fmt.Errorf("error opening %v, %v", cfgFileName, err)
	}
}

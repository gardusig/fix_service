package server

import (
	"fmt"

	"github.com/quickfixgo/quickfix"
)

type ServerApp struct{}

func (a ServerApp) OnCreate(sessionID quickfix.SessionID) {
	fmt.Println("Created session:", sessionID)
}

func (a ServerApp) OnLogon(sessionID quickfix.SessionID) {
	fmt.Println("Sending login message. sessionId: ", sessionID)
}

func (a ServerApp) OnLogout(sessionID quickfix.SessionID) {
	fmt.Println("Sending logout message. sessionId: ", sessionID)
}

func (a ServerApp) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	fmt.Println("Sending heartbeat. sessionId: ", sessionID, ", msg: ", msg)
}

func (a ServerApp) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	fmt.Println("Received heartbeat. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a ServerApp) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error {
	fmt.Println("Sending message. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a ServerApp) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	fmt.Println("Received message. sessionId: ", sessionID, ", msg:", msg)
	return
}

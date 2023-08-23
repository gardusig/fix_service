package fix

import (
	"fmt"

	"github.com/quickfixgo/quickfix"
)

type GenericApp struct{}

func (a GenericApp) OnCreate(sessionID quickfix.SessionID) {
	fmt.Println("Created session:", sessionID)
}

func (a GenericApp) OnLogon(sessionID quickfix.SessionID) {
	fmt.Println("Sending login message. sessionId: ", sessionID)
}

func (a GenericApp) OnLogout(sessionID quickfix.SessionID) {
	fmt.Println("Sending logout message. sessionId: ", sessionID)
}

func (a GenericApp) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	fmt.Println("Sending heartbeat. sessionId: ", sessionID, ", msg: ", msg)
}

func (a GenericApp) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	fmt.Println("Received heartbeat. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a GenericApp) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error {
	fmt.Println("Sending message. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a GenericApp) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	fmt.Println("Received message. sessionId: ", sessionID, ", msg:", msg)
	return nil
}

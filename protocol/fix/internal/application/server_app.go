package application

import (
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type ServerApp struct{}

func (a ServerApp) OnCreate(sessionID quickfix.SessionID) {
	logrus.Debug("Created session:", sessionID)
}

func (a ServerApp) OnLogon(sessionID quickfix.SessionID) {
	logrus.Debug("Sending login message. sessionId: ", sessionID)
}

func (a ServerApp) OnLogout(sessionID quickfix.SessionID) {
	logrus.Debug("Sending logout message. sessionId: ", sessionID)
}

func (a ServerApp) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	logrus.Debug("Sending heartbeat. sessionId: ", sessionID, ", msg: ", msg)
}

func (a ServerApp) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	logrus.Debug("Received heartbeat. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a ServerApp) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error {
	logrus.Debug("Sending message. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a ServerApp) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	logrus.Debug("Received message. sessionId: ", sessionID, ", msg:", msg)
	return nil
}

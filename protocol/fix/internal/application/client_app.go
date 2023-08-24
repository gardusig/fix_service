package application

import (
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type AppClient struct{}

func (a AppClient) OnCreate(sessionID quickfix.SessionID) {
	logrus.Debug("Created session: ", sessionID)
}

func (a AppClient) OnLogon(sessionID quickfix.SessionID) {
	logrus.Debug("Sending login message. sessionId: ", sessionID)
}

func (a AppClient) OnLogout(sessionID quickfix.SessionID) {
	logrus.Debug("Sending logout message. sessionId: ", sessionID)
}

func (a AppClient) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	logrus.Debug("Sending heartbeat. sessionId: ", sessionID, ", msg: ", msg)
}

func (a AppClient) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	logrus.Debug("Received heartbeat. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a AppClient) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) error {
	logrus.Debug("Sending message. sessionId: ", sessionID, ", msg: ", msg)
	return nil
}

func (a AppClient) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	logrus.Debug("Received message. sessionId: ", sessionID, ", msg:", msg)
	return nil
}

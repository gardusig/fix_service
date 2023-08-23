package fix

import (
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type ServerFIX struct {
	messageStoreFactory quickfix.MessageStoreFactory
	logFactory          quickfix.LogFactory

	settings    *quickfix.Settings
	application quickfix.Application
	acceptor    *quickfix.Acceptor
}

func NewServerFIX(filepath string) (*ClientFIX, error) {
	settings, err := getSettingsFromFile(filepath)
	if err != nil {
		return nil, err
	}
	client := ClientFIX{
		settings:            settings,
		application:         GenericApp{},
		messageStoreFactory: quickfix.NewMemoryStoreFactory(),
		logFactory:          quickfix.NewScreenLogFactory(),
	}
	return &client, nil
}

func (s ServerFIX) Start() error {
	logrus.Debug("Starting FIX client")
	var err error
	s.acceptor, err = quickfix.NewAcceptor(
		s.application,
		s.messageStoreFactory,
		s.settings,
		s.logFactory,
	)
	if err != nil {
		logrus.Debug("Failed to create fix initiator, reason: ", err.Error())
		return err
	}
	return s.acceptor.Start()
}

func (c ServerFIX) Stop() {
	c.acceptor.Stop()
}

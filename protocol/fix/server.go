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
	initiator   *quickfix.Initiator
}

func NewServerFIX(filepath string) (*ClientFIX, error) {
	settings, err := getSettingsFromFile(filepath)
	if err != nil {
		return nil, err
	}
	app := GenericApp{}
	client := ClientFIX{
		settings:            settings,
		application:         app,
		messageStoreFactory: quickfix.NewMemoryStoreFactory(),
		logFactory:          quickfix.NewScreenLogFactory(),
	}
	return &client, nil
}

func (c ServerFIX) Start() error {
	logrus.Debug("Starting FIX client")
	var err error
	c.initiator, err = quickfix.NewInitiator(
		c.application,
		c.messageStoreFactory,
		c.settings,
		c.logFactory,
	)
	if err != nil {
		logrus.Debug("Failed to create fix initiator, reason: ", err.Error())
		return err
	}
	return c.initiator.Start()
}

func (c ServerFIX) Stop() {
	c.initiator.Stop()
}

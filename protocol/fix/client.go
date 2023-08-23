package fix

import (
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type ClientFIX struct {
	messageStoreFactory quickfix.MessageStoreFactory
	logFactory          quickfix.LogFactory

	settings    *quickfix.Settings
	application quickfix.Application
	initiator   *quickfix.Initiator
}

func NewClientFIX(filepath string) (*ClientFIX, error) {
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

func (c ClientFIX) Start() error {
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

func (c ClientFIX) Stop() {
	c.initiator.Stop()
}

package fix

import (
	"github.com/gardusig/fix_service/protocol/fix/internal"
	"github.com/gardusig/fix_service/protocol/fix/internal/application"
	"github.com/quickfixgo/quickfix"
	"github.com/sirupsen/logrus"
)

type Client struct {
	application         quickfix.Application
	logFactory          quickfix.LogFactory
	messageStoreFactory quickfix.MessageStoreFactory
	settings            *quickfix.Settings

	initiator *quickfix.Initiator
}

func NewClientFIX(filepath string) (*Client, error) {
	settings, err := internal.GetSettingsFromFile(filepath)
	if err != nil {
		return nil, err
	}
	client := Client{
		application:         application.AppClient{},
		logFactory:          quickfix.NewScreenLogFactory(),
		messageStoreFactory: quickfix.NewMemoryStoreFactory(),
		settings:            settings,
	}
	return &client, nil
}

func (c Client) Start() error {
	logrus.Debug("Starting FIX client...")
	initiator, err := quickfix.NewInitiator(
		c.application,
		c.messageStoreFactory,
		c.settings,
		c.logFactory,
	)
	if err != nil {
		logrus.Debug("Failed to create fix initiator, reason: ", err.Error())
		return err
	}
	c.initiator = initiator
	return c.initiator.Start()
}

func (c Client) Stop() {
	c.initiator.Stop()
}

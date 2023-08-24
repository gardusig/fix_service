package internal

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/quickfixgo/quickfix"
)

func GetSettingsFromFile(filepath string) (*quickfix.Settings, error) {
	cfg, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening %v, %v", filepath, err)
	}
	defer cfg.Close()
	stringData, readErr := io.ReadAll(cfg)
	if readErr != nil {
		return nil, fmt.Errorf("error reading cfg: %s,", readErr)
	}
	reader := bytes.NewReader(stringData)
	return quickfix.ParseSettings(reader)
}

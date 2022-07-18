package telegram

import (
	"testing"
)

func TestSendMessage(t *testing.T) {
	config, err := readConfig("telegram.cfg")
	if err != nil {
		t.Errorf("SendMessage():readConfig() error = %v", err)
		return
	}
	if err := SendMessage(config.ID, config.Token, "Teste"); err != nil {
		t.Errorf("SendMessage() error = %v", err)
	}
}

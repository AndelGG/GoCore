package telegram

import (
	"awesomeProject/internal/infrastructure/telegram"
	"testing"
)

func TestSendMessage(t *testing.T) {
	tg := telegram.TelegramApi{Api: ""}

	if err := tg.SendMessage(100, "rofl"); err != nil {
		t.Error(err)
	}
}

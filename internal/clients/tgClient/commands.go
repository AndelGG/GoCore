package tgClient

import (
	"context"
	"log"
	"strings"
)

const (
	startCmd             = "/start"
	messageToDeepSeekCmd = "/f"
)

func handleCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	switch text {
	case startCmd:
		return p.sendRandom(ctx, chatID, username)
	case messageToDeepSeekCmd:
		return p.sendHelp(ctx, chatID)
	default:
		return p.tg.SendMessage(ctx, chatID, msgUnknownCommand)
	}
}

func sendToDeepseek(ctx context.Context, chatID int, pageURL string, username string) (err error) {
	resp, err := q.request.SendMessage(rawMsg)
	if err != nil {
		q.respond.CatchError(ErrChatBot)
	}

	err = q.respond.SendMessage(resp)
	if err != nil {
		panic(err)
	}
}

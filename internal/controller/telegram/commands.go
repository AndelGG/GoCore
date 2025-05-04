package telegram

import (
	"awesomeProject/internal/lib/e"
	"log"
	"strings"
)

const (
	StartCmd = "/start"
)

// doCmd is api router
func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isRespCmd(text) {
		p.sendResponseToChat(chatID, text)
	}

	switch text {
	case StartCmd:
		return p.tg.SendMessage(chatID, msgHello)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}

	return nil
}

func isRespCmd(text string) bool {
	if []rune(text)[0] != '/' {
		return true
	}
	return false
}

func (p *Processor) sendResponseToChat(chatID int, text string) (err error) {
	defer func() { err = e.WrapIfErr("cant do command: saved page", err) }()

	response := p.useCase.Message(text)
	p.tg.SendMessage(chatID, response)

	if err := p.tg.SendMessage(chatID, msgSent); err != nil {
		return err
	}

	return nil
}

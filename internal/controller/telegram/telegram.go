package telegram

import (
	"awesomeProject/internal/clients/tgClient"
	"awesomeProject/internal/controller"
	"awesomeProject/internal/lib/e"
	"awesomeProject/internal/usecases"
	"errors"
	"fmt"
)

type Processor struct {
	tg      *tgClient.Client
	offset  int
	useCase usecases.ResponderUseCase
}

type Meta struct {
	ChatID   int
	UserName string
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client *tgClient.Client, useCase usecases.ResponderUseCase) *Processor {
	return &Processor{
		tg:      client,
		useCase: useCase,
	}
}

// Fetch convert Update to Event
func (p *Processor) Fetch(limit int) ([]controller.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	fmt.Println(updates)

	if err != nil {
		return nil, e.Wrap("cant get events", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]controller.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *Processor) Process(event controller.Event) error {
	switch event.Type {
	case controller.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("can't process message", ErrUnknownEventType)
	}

}

func (p *Processor) processMessage(event controller.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("cant process message", err)
	}

	if err := p.doCmd(event.Text, meta.ChatID, meta.UserName); err != nil {
		return e.Wrap("cant process message", err)
	}

	return nil
}

func meta(event controller.Event) (Meta, error) {
	res, ok := event.Meta.(Meta) // Type assertion
	if !ok {
		return Meta{}, e.Wrap("cant get meta", ErrUnknownMetaType)
	}

	return res, nil
}

func event(upd tgClient.Update) controller.Event {
	updType := fetchType(upd)

	res := controller.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == controller.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			UserName: upd.Message.From.Username,
		}
	}

	return res
}

func fetchText(upd tgClient.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}

func fetchType(upd tgClient.Update) controller.Type {
	if upd.Message == nil {
		return controller.Unknown
	}
	return controller.Message
}

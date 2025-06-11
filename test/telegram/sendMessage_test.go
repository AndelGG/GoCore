package telegram

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/infrastructure/deepSeek"
	"awesomeProject/internal/infrastructure/telegram"
	"awesomeProject/internal/lib/req"
	"awesomeProject/internal/useCases"
	"testing"
)

func TestSendMessage(t *testing.T) {
	ШпэкЕблан := telegram.Api{Api: ""}
	БизнесШпэк := useCases.TelegramReplyUseCase{TgResponder: ШпэкЕблан}

	msg := &domain.ServiceMessage{ChatId: 675043740, RequestText: "соси хуй"}

	if err := БизнесШпэк.SendMessage(msg); err != nil {
		t.Error(err)
	}
}

func TestDeepSeekResponse(t *testing.T) {
	t.Run("deepseek-infrastructure", func(t *testing.T) {
		chat := deepSeek.New("")

		res, err := chat.RequestToChatBot("Привет", "deepseek-chat", 40)
		if err != nil {
			t.Error(err)
		}

		if len(res.Choices) == 0 {
			t.Error(err)
		}
	})

	t.Run("request to api", func(t *testing.T) {
		scheme := `{
	 "messages": [
	   {
	     "content": "You are a helpful assistant",
	     "role": "system"
	   },
	   {
	     "content": "Hello",
	     "role": "user"
		}
	 ],
	 "model": "deepseek-chat",
	 "frequency_penalty": 0,
	 "max_tokens": 40,
	 "presence_penalty": 0,
	 "response_format": {
	   "type": "text"
	 },
	 "stop": null,
	 "stream": false,
	 "stream_options": null,
	 "temperature": 1,
	 "top_p": 1,
	 "tools": null,
	 "tool_choice": "none",
	 "logprobs": false,
	 "top_logprobs": null
	}`

		req.MakeAuthorizationRequest(scheme, "")
	})
}

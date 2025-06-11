package deepSeek

import (
	"awesomeProject/internal/lib/req"
	"encoding/json"
	"fmt"
)

type ChatBot struct {
	Api string
}

func New(api string) *ChatBot {
	return &ChatBot{api}
}

func (c *ChatBot) RequestToChatBot(text, model string, maxToken int) (ResponseScheme, error) {

	const op = "infrastructure.deepSeek.RequestToChatBot"

	scheme := returnResponseScheme(maxToken, text, model)

	body, err := req.MakeAuthorizationRequest(scheme, c.Api)
	if err != nil {
		return ResponseScheme{}, fmt.Errorf("%s: %w", op, err)
	}

	// TODO: refactor
	var responseObject ResponseScheme
	if err := json.Unmarshal(body, &responseObject); err != nil {
		fmt.Println(fmt.Sprintf("body: %b", body))
		return ResponseScheme{}, fmt.Errorf("%s: %w", op, err)
	}

	return responseObject, nil
}

func returnResponseScheme(maxToken int, text, model string) string {
	return fmt.Sprintf(`{
	 "messages": [
	   {
	     "content": "You are a helpful assistant",
	     "role": "system"
	   },
	   {
	     "content": "%s",
	     "role": "user"
		}
	 ],
	 "model": "%s",
	 "frequency_penalty": 0,
	 "max_tokens": %d,
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
	}`, text, model, maxToken)
}

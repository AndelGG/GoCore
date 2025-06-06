package deepSeek

import (
	"awesomeProject/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ChatBot struct {
	Api string
}

func New(api string) *ChatBot {
	return &ChatBot{api}
}

func (c *ChatBot) RequestToChatBot(message *domain.ServiceMessage) (domain.ResponseScheme, error) {

	const op = "chatbot.dp.reqToBot"

	url := "https://api.deepseek.com/chat/completions"
	method := "POST"

	req, err := createRequest(*message, method, url)
	if err != nil {
		return domain.ResponseScheme{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Api))

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return domain.ResponseScheme{}, fmt.Errorf("%s: %w", op, err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var responseObject domain.ResponseScheme
	if err := json.Unmarshal(body, &responseObject); err != nil {
		fmt.Println(fmt.Sprintf("body: %b", body))
		return domain.ResponseScheme{}, fmt.Errorf("%s: %w", op, err)
	}

	return responseObject, nil

}

func createRequest(msg domain.ServiceMessage, method, url string) (*http.Request, error) {
	const op = "chatbot.createRequest"

	payload := strings.NewReader(fmt.Sprintf(`{
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
  "model": "deepseek-chat",
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
}`, msg.Response, msg.MaxToken))

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return req, nil
}

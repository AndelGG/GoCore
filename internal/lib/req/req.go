package req

import (
	"awesomeProject/internal/domain"
	"fmt"
	"net/http"
	"strings"
)

const (
	deepSeekAPI   = "https://api.deepseek.com/chat/completions"
	openRouterAPI = "https://openrouter.ai/api/v1"
)

func CreateRequest(r *domain.ServiceMessage, ApiKey string) (*http.Request, error) {
	const op = "lib.createRequest"
	method := "POST"
	var url string

	if r.Model == "deepseek-chat" || r.Model == "deepseek-reasoner" {
		url = deepSeekAPI
	} else {
		url = openRouterAPI
	}

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
}`, r.RequestText, r.MaxToken))

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ApiKey))

	return req, nil
}

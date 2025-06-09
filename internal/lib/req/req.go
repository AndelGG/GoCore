package req

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	deepSeekAPI   = "https://api.deepseek.com/chat/completions"
	openRouterAPI = "https://openrouter.ai/api/v1"
)

func MakeAuthorizationRequest(url, scheme, apiKey string) ([]byte, error) {
	payload := strings.NewReader(scheme)

	req, err := CreateRequest(url, http.MethodPost, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}

func MakeGetRequest(url string, query url.Values) ([]byte, error) {
	req, err := CreateRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}

func CreateRequest(url, method string, payload io.Reader) (*http.Request, error) {
	const op = "lib.createRequest"

	//if r.Model == "deepseek-chat" || r.Model == "deepseek-reasoner" {
	//	url = deepSeekAPI
	//} else {
	//	url = openRouterAPI
	//}

	//		fmt.Sprintf(`{
	//  "messages": [
	//    {
	//      "content": "You are a helpful assistant",
	//      "role": "system"
	//    },
	//    {
	//      "content": "%s",
	//      "role": "user"
	//	}
	//  ],
	//  "model": "deepseek-chat",
	//  "frequency_penalty": 0,
	//  "max_tokens": %d,
	//  "presence_penalty": 0,
	//  "response_format": {
	//    "type": "text"
	//  },
	//  "stop": null,
	//  "stream": false,
	//  "stream_options": null,
	//  "temperature": 1,
	//  "top_p": 1,
	//  "tools": null,
	//  "tool_choice": "none",
	//  "logprobs": false,
	//  "top_logprobs": null
	//}`, r.RequestText, r.MaxToken)
	//)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

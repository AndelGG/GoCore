package req

import (
	"context"
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

func MakeAuthorizationRequest(ctx context.Context, scheme, apiKey string) ([]byte, error) {
	payload := strings.NewReader(scheme)

	url := deepSeekAPI

	req, err := CreateRequest(ctx, url, http.MethodPost, payload)
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

func MakeGetRequest(ctx context.Context, url string, query url.Values) ([]byte, error) {
	req, err := CreateRequest(ctx, url, http.MethodGet, nil)
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

func CreateRequest(ctx context.Context, url, method string, payload io.Reader) (*http.Request, error) {
	const op = "lib.createRequest"

	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

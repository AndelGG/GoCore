package tgClient

import (
	"awesomeProject/internal/lib/e"
	"awesomeProject/internal/lib/tg"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const getUpdatesMethod = "getUpdates"
const sendMessageMethod = "sendMessage"

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) (updates []tg.Update, err error) {
	defer func() { err = e.WrapIfErr("cant get updates", err) }()
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(q, getUpdatesMethod)
	if err != nil {
		return nil, err
	}

	var res tg.UpdatesResponse

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res.Result, nil
}

func (c *Client) SendMessage(chatID int, txt string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", txt)

	_, err := c.doRequest(q, sendMessageMethod)
	if err != nil {
		return e.Wrap("cant send message", err)
	}
	return nil
}

func (c *Client) doRequest(query url.Values, method string) (data []byte, err error) {
	defer func() { err = e.WrapIfErr("cant do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

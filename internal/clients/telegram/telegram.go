// Package telegram contains logic for interacting with a telegram server
package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/m-kuzmin/golang-telegram-bot/internal/e"
)

// Client struct can be used to interact with a telegram server
type Client struct {
	host     string
	basePath string
	client   http.Client
}

// New returns a telegram client that is used for talking to the telegram servers
func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

// Updates fetches updates from the telegram server
func (c *Client) Updates(offset, limit int) (_ []Update, err error) {
	defer func() { err = e.Wrap("Error while fetching updates: %w", err) }()

	// Prepare request parameters from func args
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	// Send a request
	data, err := c.doRequest("getUpdates", q)
	if err != nil {
		return nil, err
	}

	// Parse the responce and return the results
	var res updatesResponce
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	if res.Ok {
		return res.Result, nil
	} else {
		return nil, fmt.Errorf(res.Description)
	}
}

// SendMessage sends a message to `chat` with `text`
func (c *Client) SendMessage(chatID int, text string) error {
	// Prepare the request
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	// Try to send the messages to chat
	_, err := c.doRequest("sendMessage", q)
	if err != nil {
		return e.Wrap("Error while sending message", err)
	}
	return nil
}

// Abstracts an API call to a telegram server
func (c *Client) doRequest(endpoint string, query url.Values) (data []byte, err error) {
	// automatically wraps an error into a context message
	defer func() { err = e.Wrap("Error while doing request. Endpoint:", err) }()

	// Set up the request
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, endpoint),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()

	// Make an API call
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	// Dont forget to stop reading the request.
	defer func() { _ = resp.Body.Close() }()

	// Read the responce and return it
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

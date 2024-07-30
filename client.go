package pterodactyl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default pterodactyl URL
const HostURL string = "https://panel.localhost"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient -
func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default pterodactyl URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If token not provided, return empty client
	if token == nil {
		return &c, nil
	}

	c.Token = *token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "Application/vnd.pterodactyl.v1+json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, nil
}

func (c *Client) prepareBody(body interface{}) io.Reader {
	b, _ := json.Marshal(body)
	return bytes.NewReader(b)
}

package emqx

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	json "github.com/json-iterator/go"
)

type Client struct {
	apiurl,
	appid,
	appscret string
	hcli *http.Client
}

type ClientOptions func(*Client)

func WithHttpClient(hcli *http.Client) ClientOptions {
	return func(c *Client) {
		c.hcli = hcli
	}
}

func NewClient(apiurl, appid, appscret string, opts ...ClientOptions) *Client {
	c := &Client{
		appid:    appid,
		appscret: appscret,
		hcli:     http.DefaultClient,
	}
	c.SetApiUrl(apiurl)
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) SetApiUrl(apiurl string) {
	u, err := url.Parse(apiurl)
	if err != nil {
		c.apiurl = apiurl
	}
	u.Path = `/api/v5`
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	c.apiurl = u.String()
}

func (c *Client) newRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	up, err := url.JoinPath(c.apiurl, path)
	if err != nil {
		return nil, err
	}

	var br io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		br = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, up, br)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.appid, c.appscret)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

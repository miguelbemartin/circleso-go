package circle

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiURL = "https://app.circle.so/api/v1/"

// ErrAccessTokenNotSet is returned when Client methods are called that require
// an access token to be set.
var ErrAccessTokenNotSet = errors.New("access token not set on client")

// Client for working with the Circle API.
type Client struct {
	apiToken     string
	baseURL      string
	c            *http.Client
	errorHandler func(e error) error
}

// Option sets an optional setting on the Client.
type Option func(*Client)

// NewClient returns a new client for working with the Circle API.
func NewClient(apiToken string) *Client {
	return &Client{
		apiToken: apiToken,
		baseURL:  apiURL,
		c:        http.DefaultClient,
	}
}

func (c *Client) newRequest(
	ctx context.Context, method, endpoint string, body interface{}, queryOps interface{},
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, nil)
	req.Header.Add("Authorization", "Token "+c.apiToken)
	if err != nil {
		return nil, err
	}

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	}

	if queryOps != nil {
		vs, err := query.Values(queryOps)
		if err != nil {
			return nil, err
		}
		appendQueryValues(req, vs)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // nolint: errcheck

	if resp.StatusCode >= 299 {
		e := NewError(resp)
		if c.errorHandler != nil {
			return c.errorHandler(e)
		}
		return e
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}
	return nil
}

func appendQueryValues(req *http.Request, values url.Values) {
	q := req.URL.Query()
	for k, vs := range values {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()
}

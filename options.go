package vk

import (
	"github.com/go-vk-api/vk/httputil"
)

// Option is a configuration option to initialize a client.
type Option func(*Client) error

// WithToken overrides the client token with the specified one.
func WithToken(token string) Option {
	return func(c *Client) error {
		c.Token = token

		return nil
	}
}

// WithHttpClient overrides the client http client with the specified one.
func WithHttpClient(doer httputil.RequestDoer) Option {
	return func(c *Client) error {
		c.HttpClient = doer

		return nil
	}
}

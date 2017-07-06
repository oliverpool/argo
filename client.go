package argo

import (
	"encoding/base64"
)

type Client struct {
	Caller        Caller
	DefaultOption Option
}

func (c Client) AddUri(uris []string, options ...Option) (Response, error) {
	opt := c.mergeOptions(options...)
	return c.Caller.Call("aria2.addUri", uris, opt)
}

func (c Client) AddTorrent(content []byte, options ...Option) (Response, error) {
	return c.AddTorrentWithWebSeed(content, []string{}, options...)
}
func (c Client) AddTorrentWithWebSeed(content []byte, uris []string, options ...Option) (Response, error) {
	opt := c.mergeOptions(options...)
	b64 := base64.StdEncoding.EncodeToString(content)
	return c.Caller.Call("aria2.addTorrent", b64, uris, opt)
}

func (c Client) Shutdown() (Response, error) {
	return c.Caller.Call("aria2.shutdown")
}

// Close gracefully closes the Caller
func (c Client) Close() (err error) {
	return c.Caller.Close()
}

func (c Client) mergeOptions(options ...Option) Option {
	newOption := mergeOptions(options...)
	return mergeOptions(c.DefaultOption, newOption)
}

func mergeOptions(options ...Option) Option {
	opt := Option{}
	for _, o := range options {
		for k, v := range o {
			opt[k] = v
		}
	}
	return opt
}

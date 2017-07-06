package argo

import (
	"encoding/base64"
)

type Client struct {
	Caller        Caller
	DefaultOption Option
}

// AddUri adds a new download
// - uris is an array of HTTP/FTP/SFTP/BitTorrent URIs (strings) pointing to the same resource. If you mix URIs pointing to different resources, then the download may fail or be corrupted without aria2 complaining. When adding BitTorrent Magnet URIs, uris must have only one element and it should be BitTorrent Magnet URI. options is a struct and its members are pairs of option name and value.
// This method returns the GID of the newly registered download.
func (c Client) AddUri(uris []string, options ...Option) (Response, error) {
	opt := c.mergeOptions(options...)
	return c.Caller.Call("aria2.addUri", uris, opt)
}

// AddTorrent adds a BitTorrent download by uploading the content of a ".torrent" file. If you want to add a BitTorrent Magnet URI, use the AddUri() method instead.
// This method returns the GID of the newly registered download
// If --rpc-save-upload-metadata is true, the uploaded data is saved as a file named as the hex string of SHA-1 hash of data plus ".torrent" in the directory specified by --dir option. E.g. a file name might be 0a3893293e27ac0490424c06de4d09242215f0a6.torrent. If a file with the same name already exists, it is overwritten!
// If the file cannot be saved successfully or --rpc-save-upload-metadata is false, the downloads added by this method are not saved by --save-session.
func (c Client) AddTorrent(content []byte, options ...Option) (Response, error) {
	return c.AddTorrentWithWebSeed(content, []string{}, options...)
}

// AddTorrentWithWebSeed adds a BitTorrent download by uploading the content of a ".torrent" file. If you want to add a BitTorrent Magnet URI, use the AddUri() method instead.
// uris is an array of URIs (string). uris is used for Web-seeding. For single file torrents, the URI can be a complete URI pointing to the resource; if URI ends with /, name in torrent file is added. For multi-file torrents, name and path in torrent are added to form a URI for each file.
// This method returns the GID of the newly registered download
// If --rpc-save-upload-metadata is true, the uploaded data is saved as a file named as the hex string of SHA-1 hash of data plus ".torrent" in the directory specified by --dir option. E.g. a file name might be 0a3893293e27ac0490424c06de4d09242215f0a6.torrent. If a file with the same name already exists, it is overwritten!
// If the file cannot be saved successfully or --rpc-save-upload-metadata is false, the downloads added by this method are not saved by --save-session.
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
	customOption := mergeOptions(options...)
	return mergeOptions(c.DefaultOption, customOption)
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

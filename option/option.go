// Package option gathers some of the options that the aria2 client (and the daemon) may take in argument.
// Full list on https://aria2.github.io/manual/en/html/aria2c.html#id2
//
// For a custom option, simply do argo.Option{"key": "value"}
package option

import (
	"github.com/oliverpool/argo"
)

// Dir sets the directory to store the downloaded file.
func Dir(dir string) argo.Option {
	return argo.Option{"dir": dir}
}

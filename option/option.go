// Package option gathers some of the options that the aria2 client (and the daemon) may take in argument.
// Full list on https://aria2.github.io/manual/en/html/aria2c.html#id2
//
// For a custom option, simply do argo.Option{"key": "value"}
package option

import (
	"path/filepath"

	"github.com/oliverpool/argo"
)

// Dir sets the directory to store the downloaded file.
func Dir(dir string) argo.Option {
	return argo.Option{"dir": dir}
}

// Out sets the file name of the downloaded file.
// It is always relative to the directory given in --dir option. When the --force-sequential option is used, this option is ignored.
//
// Note: You cannot specify a file name for Metalink or BitTorrent downloads. The file name specified here is only used when the URIs fed to aria2 are given on the command line directly, but not when using --input-file, --force-sequential option.
func Out(file string) argo.Option {
	return argo.Option{"out": file}
}

// Dst sets the destination file (combination of Dir and Out).
func Dst(fullpath string) argo.Option {
	dir, file := filepath.Split(fullpath)
	return argo.Option{"dir": dir, "out": file}
}

// Package option gathers some of the options the aria2 methods may take in argument.
// Full list on https://aria2.github.io/manual/en/html/aria2c.html#id2
package option

import (
	"github.com/oliverpool/argo"
)

func Dir(dir string) argo.Option {
	return argo.Option{"dir": dir}
}

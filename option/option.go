package option

import (
	"github.com/oliverpool/argo"
)

func Dir(dir string) argo.Option {
	return argo.Option{"dir": dir}
}

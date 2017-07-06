package daemon

import (
	"fmt"

	"github.com/oliverpool/argo"
)

type option func(a *Aria2)

// CmdOption applies some options on the command
// Usage: a.Option(daemon.Port("6800"), daemon.ListenAll)
func (a *Aria2) CmdOption(opts ...option) {
	for _, opt := range opts {
		opt(a)
	}
}

// Option applies some options on the command
// Usage: a.Option(daemon.Port("6800"), daemon.ListenAll)
func (a *Aria2) Option(opts ...argo.Option) {
	for _, opt := range opts {
		for k, raw := range opt {
			switch v := raw.(type) {
			case bool:
				if v {
					AppendArg("--" + k)(a)
				} else {
					AppendArg("--" + k + "=false")(a)
				}
			default:
				s := fmt.Sprintf("--%v=%v", k, v)
				AppendArg(s)(a)
			}
		}
	}
}

var ListenAll = argo.Option{"rpc-listen-all": true}
var EnableRPC = argo.Option{"enable-rpc": true}

func Secret(secret string) argo.Option {
	if secret != "" {
		return argo.Option{"rpc-secret": secret}
	} else {
		return argo.Option{}
	}
}

func Port(port string) argo.Option {
	return argo.Option{"rpc-listen-port": port}
}

// Log level to output to console. LEVEL is either debug, info, notice, warn or error. Default: notice
func Log(level string) argo.Option {
	return argo.Option{"console-log-level": level}
}

func AppendArg(args ...string) option {
	return func(a *Aria2) {
		a.args = append(a.args, args...)
	}
}

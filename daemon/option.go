package daemon

import (
	"fmt"

	"github.com/oliverpool/argo"
)

// EnableRPC enables JSON-RPC/XML-RPC server.
// It is strongly recommended to set secret authorization token using daemon.Secret option
// See also daemon.Port option
var EnableRPC = argo.Option{"enable-rpc": true}

// ListenAll listens incoming JSON-RPC/XML-RPC requests on all network interfaces.
// Default: listen only on local loopback interface (when option omitted).
var ListenAll = argo.Option{"rpc-listen-all": true}

// LogLevel to output to console. LEVEL is either debug, info, notice, warn or error.
// Default: notice (when option omitted)
func LogLevel(level string) argo.Option {
	return argo.Option{"console-log-level": level}
}

// Secret set RPC secret authorization token.
func Secret(secret string) argo.Option {
	if secret != "" {
		return argo.Option{"rpc-secret": secret}
	}
	return argo.Option{}
}

// Port specifies a port number for JSON-RPC/XML-RPC server to listen to.
// Possible Values: 1024 -65535.
// Default: 6800 (when option omitted)
func Port(port string) argo.Option {
	return argo.Option{"rpc-listen-port": port}
}

// Option applies some options on the command.
//   aria2.Option(daemon.Port("6800"), daemon.ListenAll)
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

type cmdOption func(a *Aria2)

// CmdOption applies some cmdOptions on the command.
//   aria2.CmdOption(daemon.AppendArg("some_raw_arg"))
func (a *Aria2) CmdOption(opts ...cmdOption) {
	for _, opt := range opts {
		opt(a)
	}
}

// AppendArg creates an cmdOption to appends string to the command line.
//   aria2.CmdOption(daemon.AppendArg("some_raw_arg"))
func AppendArg(args ...string) cmdOption {
	return func(a *Aria2) {
		a.args = append(a.args, args...)
	}
}

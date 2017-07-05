package daemon

type option func(a *Aria2)

// Option applies some options on the command
// Usage: a.Option(daemon.Port("6800"), daemon.ListenAll)
func (a *Aria2) Option(opts ...option) {
	for _, opt := range opts {
		opt(a)
	}
}

func AppendArg(args ...string) option {
	return func(a *Aria2) {
		a.args = append(a.args, args...)
	}
}

var ListenAll = AppendArg("--rpc-listen-all")
var EnableRPC = AppendArg("--enable-rpc")

func Secret(secret string) option {
	return AppendArg("--rpc-secret=" + secret)
}

func Port(port string) option {
	return AppendArg("--rpc-listen-port=" + port)
}

// Log level to output to console. LEVEL is either debug, info, notice, warn or error. Default: notice
func Log(level string) option {
	return AppendArg("--console-log-level=" + level)
}

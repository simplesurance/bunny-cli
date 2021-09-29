package global

import (
	"sync/atomic"

	bunny "github.com/simplesurance/bunny-go"
)

// Args contains commandline arguments that are shared across multiple
// commands.
type Args struct {
	APIKey string
	Client *bunny.Client
}

var args atomic.Value

func init() {
	args.Store(Args{})
}

// Get returns a copy of the globally stored arguments.
// The function is thread-safe.
func Get() Args {
	return (args.Load().(Args))
}

// Set sets the global variable to a.
// The function is thread-safe.
func Set(a Args) {
	args.Store(a)
}

package exithandler

import (
	"context"
	"os/signal"
	"syscall"
)

// HandlerFunc represents an exit function to be executed when the program receives a SIGINT ou SIGTERM signal.
type HandlerFunc func(context.Context)

// Handle inits a new context listening to the global signals and waits on Done channel to execute the given input function.
func Handle(parent context.Context, exithandler HandlerFunc) {
	ctx, cancel := signal.NotifyContext(parent, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	<-ctx.Done()
	exithandler(ctx)
}

// HandleFunc returns a function initializing a new context listening to the global signals and awaiting on Done channel to execute the given input function.
func HandleFunc(parent context.Context, exithandler HandlerFunc) func() {
	return func() { Handle(parent, exithandler) }
}

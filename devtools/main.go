package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/imversed/imversed/devtools/cmd"
)

func main() {
	ctx := InitCtx(context.Background())

	err := cmd.New(ctx).ExecuteContext(ctx)

	if ctx.Err() == context.Canceled || err == context.Canceled {
		fmt.Println("aborted")
		return
	}

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}

// InitCtx creates a new context from ctx that is canceled when an exit signal received.
func InitCtx(ctx context.Context) context.Context {
	var (
		ctxend, cancel = context.WithCancel(ctx)
		quit           = make(chan os.Signal, 1)
	)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		cancel()
	}()
	return ctxend
}

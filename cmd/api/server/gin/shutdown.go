package gin

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func NewGracefullShutdown() context.Context {
	ctx, cancelFunc := context.WithCancel(context.Background())

	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(sigint, syscall.SIGTERM)

	go func() {
		sig := <-sigint

		log.Printf("received shutdown signal: %s", sig.String())

		cancelFunc()
	}()

	return ctx
}

package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)

	defer func() {
		r := recover()
		done()

		if r != nil {
			panic(r)
		}
	}()

	<-ctx.Done()
}

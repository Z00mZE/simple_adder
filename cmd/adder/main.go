package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Z00mZE/simple_adder/internal/app/adder"
)

func main() {
	ctx, ctxClose := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer ctxClose()

	app, appClose, appError := adder.InitApplication(ctx)
	defer appClose()

	if appError != nil {
		os.Exit(1)
	}

	app.Run()
}

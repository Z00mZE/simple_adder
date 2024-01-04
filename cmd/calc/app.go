package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Z00mZE/wireone/cmd/calc/wire"
)

func main() {
	ctx, ctxClose := context.WithCancel(context.Background())
	defer ctxClose()

	app, appError := wire.Application(ctx)
	if appError != nil {
		fmt.Println("error", appError)
		os.Exit(1)
	}

	app.Run()
}

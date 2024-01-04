package calc

import (
	"context"
	"fmt"
	"os"

	"github.com/Z00mZE/wireone/domain/storage"
)

type Application struct {
	messageProvider storage.Message
	ctx             context.Context
}

func NewApplication(ctx context.Context, msg storage.Message) *Application {
	return &Application{ctx: ctx, messageProvider: msg}
}
func (a *Application) Run() {
	fmt.Println("pong", a.messageProvider.Ping())
}

func Run() {
	ctx, ctxClose := context.WithCancel(context.Background())
	defer ctxClose()

	app, appError := initialize(ctx)
	if appError != nil {
		fmt.Println("error", appError)
		os.Exit(1)
	}

	app.Run()
}

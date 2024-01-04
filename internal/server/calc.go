package server

import (
	"context"
	"fmt"

	"github.com/Z00mZE/wireone/domain/storage"
)

type Server struct {
	messageProvider storage.Message
	ctx             context.Context
}

func NewCalc(ctx context.Context, msg storage.Message) *Server {
	return &Server{ctx: ctx, messageProvider: msg}
}
func (a *Server) Run() {
	fmt.Println("pong", a.messageProvider.Ping())
}

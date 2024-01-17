package app

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"

	"github.com/Z00mZE/simple_adder/internal/adder/domain/transport"
	"github.com/Z00mZE/simple_adder/internal/adder/transport/rpc"
	"github.com/Z00mZE/simple_adder/internal/pkg/logger/sl"
)

type Adder struct {
	ctx              context.Context
	logger           *slog.Logger
	server           transport.Server
	serviceRegistrar transport.ServiceRegistrar
}

func NewAdder(ctx context.Context, srv transport.Server, serviceRegistrar transport.ServiceRegistrar, _ *rpc.Dispatcher, logger *slog.Logger) *Adder {
	return &Adder{
		ctx:              ctx,
		logger:           logger,
		server:           srv,
		serviceRegistrar: serviceRegistrar,
	}
}

func (a *Adder) Run() {
	serverErrorCh := make(chan error)

	go func() {
		if err := a.server.Start(); err != nil {
			serverErrorCh <- errors.Wrap(err, "error occurred on gRPC serve")
		}
	}()

loop:
	for {
		select {
		case <-a.ctx.Done():
			a.logger.Info("parent context was closed")
			break loop
		case serverError := <-serverErrorCh:
			a.logger.Error("shutting down the server", sl.Error(serverError))
			break loop
		}
	}
}

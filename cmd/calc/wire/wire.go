//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/wireone/internal/pkg/logger"
	server2 "github.com/Z00mZE/wireone/internal/server"
	"github.com/Z00mZE/wireone/internal/storage/message"
	"github.com/Z00mZE/wireone/internal/trasnport/server"

	"github.com/Z00mZE/wireone/internal/config"
)

func Application(ctx context.Context) (*server2.Server, error) {
	panic(
		wire.Build(
			wire.NewSet(
				server2.NewCalc,
				logger.WireSet,
				server.WireListenerSet,
				server.WireGrpcSet,
				config.WireSet,
				message.WireSet,
			),
		),
	)
	return new(server2.Server), nil
}

//go:build wireinject
// +build wireinject

package adder

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/simple_adder/internal/app"
	"github.com/Z00mZE/simple_adder/internal/conf"
	"github.com/Z00mZE/simple_adder/internal/service/sum"
	"github.com/Z00mZE/simple_adder/internal/transport/rpc"
	"github.com/Z00mZE/simple_adder/pkg/logger"
)

func InitApplication(ctx context.Context) (*app.Adder, func(), error) {
	panic(
		wire.Build(
			wire.NewSet(
				app.NewAdder,
				conf.NewConfig,
				logger.NewLogger,

				sum.WireSet,
				rpc.WireDispatcherSet,
				rpc.WireServerSet,
			),
		),
	)
	return new(app.Adder), nil, nil
}

//go:build wireinject
// +build wireinject

package wire

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/simple_adder/internal/adder/app"
	"github.com/Z00mZE/simple_adder/internal/adder/conf"
	rpc2 "github.com/Z00mZE/simple_adder/internal/adder/transport/rpc"
	"github.com/Z00mZE/simple_adder/internal/adder/usecase/sum"
	"github.com/Z00mZE/simple_adder/internal/pkg/logger"
)

func InitApplication(ctx context.Context) (*app.Adder, func(), error) {
	panic(
		wire.Build(
			wire.NewSet(
				app.NewAdder,
				conf.NewConfig,

				logger.WireSet,
				sum.WireSet,
				rpc2.WireDispatcherSet,
				rpc2.WireServerSet,
			),
		),
	)
	return new(app.Adder), nil, nil
}

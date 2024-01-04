//go:build wireinject
// +build wireinject

package calc

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/wireone/internal/storage/message"

	"github.com/Z00mZE/wireone/internal/config"
)

func initialize(ctx context.Context) (*Application, error) {
	panic(
		wire.Build(
			wire.NewSet(
				NewApplication,
				config.WireSet,
				message.WireSet,
			),
		),
	)
	return &Application{}, nil
}

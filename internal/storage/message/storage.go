package message

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/wireone/domain/registry"
	"github.com/Z00mZE/wireone/domain/storage"
)

type Storage struct {
	ctx     context.Context
	cfg     registry.Settings
	message string
}

var WireSet = wire.NewSet(
	NewStorage,
	wire.Bind(new(storage.Message), new(*Storage)),
)

func NewStorage(ctx context.Context, cfg registry.Settings) *Storage {
	return &Storage{
		ctx:     ctx,
		message: "hello, worlds",
		cfg:     cfg,
	}
}
func (s *Storage) Ping() string {
	return s.message + ":" + s.cfg.Name()
}

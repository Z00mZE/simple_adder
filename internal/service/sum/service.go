package sum

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/simple_adder/internal/domain/service"
)

var WireSet = wire.NewSet(
	NewService,
	wire.Bind(new(service.Sum), new(*Service)),
)

type Service struct {
}

func NewService() *Service {
	return new(Service)
}

func (s *Service) Sum(_ context.Context, a int64, b int64) (int64, error) {
	return a + b, nil
}

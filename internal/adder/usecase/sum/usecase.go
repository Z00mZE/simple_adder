package sum

import (
	"context"
	"log/slog"

	"github.com/google/wire"

	"github.com/Z00mZE/simple_adder/internal/adder/domain/usecase"
)

var WireSet = wire.NewSet(
	NewUseCase,
	wire.Bind(new(usecase.Sum), new(*UseCase)),
)

type UseCase struct {
	logger *slog.Logger
}

func NewUseCase(logger *slog.Logger) *UseCase {
	const op = `/internal/usecase/sum/usecase`
	return &UseCase{
		logger: logger.With(slog.String("id", op)),
	}
}

func (s *UseCase) Sum(_ context.Context, a int64, b int64) (int64, error) {
	s.logger.Info(
		"call sum method",
		slog.Group(
			"call params",
			slog.Int64("a", a),
			slog.Int64("b", b),
		),
	)
	return a + b, nil
}

package usecase

import (
	"context"
)

type Sum interface {
	Sum(ctx context.Context, a int64, b int64) (int64, error)
}

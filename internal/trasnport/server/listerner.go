package server

import (
	"context"
	"net"
	"strconv"

	"github.com/google/wire"

	"github.com/Z00mZE/wireone/domain/registry"
	"github.com/Z00mZE/wireone/domain/transport"
)

var WireListenerSet = wire.NewSet(NewGrpcListener)

func NewGrpcListener(ctx context.Context, cfg registry.Settings) (transport.GrpcListener, error) {
	const tcp = `tcp`

	ls, lsError := new(net.ListenConfig).Listen(ctx, tcp, net.JoinHostPort("", strconv.Itoa(int(cfg.GrpcServerPort()))))

	return ls.(transport.GrpcListener), lsError
}

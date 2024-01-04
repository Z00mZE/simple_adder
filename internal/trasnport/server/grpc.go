package server

import (
	"context"
	"log/slog"
	"net"
	"strconv"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Z00mZE/wireone/domain/registry"
	"github.com/Z00mZE/wireone/domain/transport"
)

type GrpcServer struct {
	grpc     *grpc.Server
	listener net.Listener
	logger   *slog.Logger
}

var WireGrpcSet = wire.NewSet(
	NewGrpcServer,
	wire.Bind(new(transport.GrpcServer), new(*GrpcServer)),
)

func NewGrpcServer(ctx context.Context, cfg registry.Settings, logger *slog.Logger) (*GrpcServer, func(), error) {
	const tcp = `tcp
`
	logger = logger.With(slog.String("id", "grpc server"))

	var listenConfig net.ListenConfig

	listener, listenerError := listenConfig.Listen(ctx, tcp, net.JoinHostPort("", strconv.Itoa(int(cfg.GrpcServerPort()))))

	if listenerError != nil {
		return nil, nil, errors.Wrap(listenerError, "an error occurred while creating the listener")
	}

	self := &GrpcServer{
		logger:   logger,
		listener: listener,
		grpc: grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				recovery.UnaryServerInterceptor(
					recovery.WithRecoveryHandler(
						func(p interface{}) (err error) {
							logger.Error("Recovered from panic", slog.Any("panic", p))
							return status.Errorf(codes.Internal, "internal error")
						},
					),
				),
				logging.UnaryServerInterceptor(
					logging.LoggerFunc(
						func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
							logger.Log(ctx, slog.Level(lvl), msg, fields...)
						},
					),
					logging.WithLogOnEvents(
						//logging.StartCall, logging.FinishCall,
						logging.PayloadReceived, logging.PayloadSent,
					),
				),
			),
		),
	}

	return self, self.Stop, nil
}
func (s *GrpcServer) Start() error {
	const op = "grpc_server.Start"

	s.logger.
		With(slog.String("op", op)).
		Info("grpc server started", slog.String("addr", s.listener.Addr().String()))

	return s.grpc.Serve(s.listener)
}
func (s *GrpcServer) Stop() {
	const op = "grpc_server.Stop"

	s.logger.
		With(slog.String("op", op)).
		Info(
			"stopping gRPC server",
			slog.String("addr", s.listener.Addr().String()),
		)

	s.grpc.GracefulStop()
}

package rpc

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
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/Z00mZE/simple_adder/internal/adder/conf"
	"github.com/Z00mZE/simple_adder/internal/adder/domain/transport"
)

type Server struct {
	logger   *slog.Logger
	grpc     *grpc.Server
	listener net.Listener
}

var WireServerSet = wire.NewSet(
	NewServer,
	wire.Bind(new(transport.Server), new(*Server)),
	wire.Bind(new(transport.ServiceRegistrar), new(*Server)),
)

func NewServer(ctx context.Context, cfg *conf.Settings, logger *slog.Logger) (*Server, func(), error) {
	const tcpNetwork = `tcp`

	logger = logger.With(slog.String("id", "grpc server"))

	var listenConfig net.ListenConfig

	listener, listenerError := listenConfig.Listen(ctx, tcpNetwork, net.JoinHostPort("", strconv.Itoa(int(cfg.GrpcPort))))
	if listenerError != nil {
		return nil, nil, errors.Wrap(listenerError, "an error occurred while creating the listener")
	}

	self := &Server{
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
						//logging.StartCall,
						logging.PayloadReceived,
						logging.PayloadSent,
						//logging.FinishCall,
					),
				),
			),
		),
	}

	return self, self.Stop, nil
}
func (s *Server) Start() error {
	const op = "grpc_server.Start"

	reflection.Register(s.grpc)

	s.logger.
		With(slog.String("op", op)).
		Info("grpc server started", slog.String("addr", s.listener.Addr().String()))

	return s.grpc.Serve(s.listener)
}
func (s *Server) Stop() {
	const op = "grpc_server.Stop"

	s.logger.
		With(slog.String("op", op)).
		Info(
			"stopping gRPC server",
			slog.String("addr", s.listener.Addr().String()),
		)

	s.grpc.GracefulStop()
}

func (s *Server) ServiceRegistrar() *grpc.Server {
	return s.grpc
}

package transport

import (
	"google.golang.org/grpc"
)

type Server interface {
	Start() error
	Stop()
}
type ServiceRegistrar interface {
	ServiceRegistrar() *grpc.Server
}

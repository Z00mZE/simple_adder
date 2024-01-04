package transport

type server interface {
	Start() error
	Stop()
}

type GrpcServer server

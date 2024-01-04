package registry

type Settings interface {
	Name() string
	GrpcServerPort() uint16
}

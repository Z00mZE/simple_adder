package config

import (
	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"

	"github.com/Z00mZE/wireone/domain/registry"
)

var WireSet = wire.NewSet(
	NewConfig,
	wire.Bind(new(registry.Settings), new(*Settings)),
)

type Settings struct {
	AppName  string `envconfig:"NAME" default:"vasya"`
	GrpcPort uint16 `envconfig:"GRPC_PORT" default:"50051"`
}

func NewConfig() (*Settings, error) {
	var conf = new(Settings)
	if err := envconfig.Process("", conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func (c *Settings) Name() string {
	return c.AppName
}

func (c *Settings) GrpcServerPort() uint16 {
	return c.GrpcPort
}

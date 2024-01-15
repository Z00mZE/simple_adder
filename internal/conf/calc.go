package conf

import (
	"github.com/kelseyhightower/envconfig"
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

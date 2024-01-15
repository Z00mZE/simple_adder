package rpc

import (
	"context"

	"github.com/google/wire"

	"github.com/Z00mZE/simple_adder/internal/domain/service"
	"github.com/Z00mZE/simple_adder/internal/domain/transport"
	"github.com/Z00mZE/simple_adder/pb"
)

var WireDispatcherSet = wire.NewSet(
	NewDispatcher,
	wire.Bind(new(pb.CalcServer), new(*Dispatcher)),
)

type Dispatcher struct {
	sumService service.Sum
}

func NewDispatcher(srv transport.ServiceRegistrar, sumService service.Sum) *Dispatcher {
	self := &Dispatcher{
		sumService: sumService,
	}

	pb.RegisterCalcServer(srv.ServiceRegistrar(), self)

	return self
}

func (d *Dispatcher) Sum(ctx context.Context, request *pb.SumRequest) (*pb.SumResponse, error) {
	result, resultError := d.sumService.Sum(ctx, request.A, request.B)
	if resultError != nil {
		return nil, resultError
	}

	out := &pb.SumResponse{C: result}

	return out, nil
}

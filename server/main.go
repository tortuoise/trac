package server

import (
	"context"
	"net"

	"github.com/golang/glog"
	"github.com/tortuoise/trac/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error {
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	defer func() {
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", network, address, err)
		}
	}()

	s := grpc.NewServer()
	pb.RegisterTracServer(s, newTracServer())
	//pb.RegisterRouteGuideServer(s, newServer())

        reflection.Register(s)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()
	return s.Serve(l)
}

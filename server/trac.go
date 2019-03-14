package server

import (
	"context"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/tortuoise/trac/pb"
	"github.com/tortuoise/trac/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var _ pb.TracServer = &tracServer{}

type tracServer struct{}

func newTracServer() pb.TracServer {
	return new(tracServer)
}

func (s *tracServer) Post(ctx context.Context, msg *pb.WrappedCoordinate) (*empty.Empty, error) {
	glog.Infof("Post %v \n", msg)
	id, err := data.PutCoordinate(&data.WrappedCoordinate{msg.User, msg.Id, int64(msg.Coord.Point.Latitude), int64(msg.Coord.Point.Longitude), int64(msg.Coord.Altitude)})
	if err != nil {
		glog.Infof("Failed to put to store:%s\n", err)
	}
        glog.Infof("Id: %v \n", id)
	return &empty.Empty{}, nil
}

func (s *tracServer) GetLast(ctx context.Context, msg *pb.CoordinateRequest) (*pb.Coordinate, error) {
	glog.Info(msg)
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return &pb.Coordinate{}, nil
}

func (s *tracServer) Get(ctx context.Context, msg *pb.TrackRequest) (*pb.Track, error) {
	glog.Infof("Get %v \n", msg)
	return &pb.Track{}, nil
}

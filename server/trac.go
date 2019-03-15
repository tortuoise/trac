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
	id, err := data.PutCoordinate(&data.WrappedCoordinate{UserId:msg.User, Latitude:(msg.Coord.Point.Latitude), Longitude:(msg.Coord.Point.Longitude), Altitude:(msg.Coord.Altitude), Timestamp: msg.TimestampValue.String()})
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
        c, err := data.GetCoordinate(msg.User)
        if err != nil {
                glog.Errorf("GetLast %v \n", err)
                return nil, err
        }
	return &pb.Coordinate{c.Altitude, &pb.Point{c.Latitude,c.Longitude}}, nil
}

func (s *tracServer) Get(ctx context.Context, msg *pb.TrackRequest) (*pb.Track, error) {
	glog.Infof("Get %v \n", msg)
        cs, err := data.GetTrack(msg.User)
        if err != nil {
                glog.Errorf("Get %v \n", err)
                return nil, err
        }
        // convert []*data.Coordinate to []*pb.Coordinate
        coords := make([]*pb.Coordinate,0)
        for _, c := range cs {
                coords = append(coords, &pb.Coordinate{c.Altitude, &pb.Point{c.Latitude, c.Longitude}})
        }
	return &pb.Track{Coords:coords}, nil
}

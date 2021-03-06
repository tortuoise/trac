// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trac.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	trac.proto

It has these top-level messages:
	TimePeriod
	Coordinate
	CoordinateRequest
	Track
	TrackRequest
	WrappedCoordinate
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"
import routeguide "github.com/tortuoise/trac/pb/routeguide"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TimePeriod struct {
	Start *google_protobuf2.Timestamp `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *google_protobuf2.Timestamp `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *TimePeriod) Reset()                    { *m = TimePeriod{} }
func (m *TimePeriod) String() string            { return proto.CompactTextString(m) }
func (*TimePeriod) ProtoMessage()               {}
func (*TimePeriod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TimePeriod) GetStart() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *TimePeriod) GetEnd() *google_protobuf2.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type Coordinate struct {
	Altitude float32           `protobuf:"fixed32,1,opt,name=altitude" json:"altitude,omitempty"`
	Point    *routeguide.Point `protobuf:"bytes,2,opt,name=point" json:"point,omitempty"`
}

func (m *Coordinate) Reset()                    { *m = Coordinate{} }
func (m *Coordinate) String() string            { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()               {}
func (*Coordinate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Coordinate) GetAltitude() float32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Coordinate) GetPoint() *routeguide.Point {
	if m != nil {
		return m.Point
	}
	return nil
}

type CoordinateRequest struct {
	User int64 `protobuf:"varint,1,opt,name=user" json:"user,omitempty"`
	Id   int64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *CoordinateRequest) Reset()                    { *m = CoordinateRequest{} }
func (m *CoordinateRequest) String() string            { return proto.CompactTextString(m) }
func (*CoordinateRequest) ProtoMessage()               {}
func (*CoordinateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CoordinateRequest) GetUser() int64 {
	if m != nil {
		return m.User
	}
	return 0
}

func (m *CoordinateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Track struct {
	Id     int64         `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Coords []*Coordinate `protobuf:"bytes,2,rep,name=coords" json:"coords,omitempty"`
}

func (m *Track) Reset()                    { *m = Track{} }
func (m *Track) String() string            { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()               {}
func (*Track) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Track) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Track) GetCoords() []*Coordinate {
	if m != nil {
		return m.Coords
	}
	return nil
}

type TrackRequest struct {
	User int64 `protobuf:"varint,1,opt,name=user" json:"user,omitempty"`
	// oneof trackfilter {
	Period *TimePeriod `protobuf:"bytes,2,opt,name=period" json:"period,omitempty"`
	Track  string      `protobuf:"bytes,3,opt,name=track" json:"track,omitempty"`
}

func (m *TrackRequest) Reset()                    { *m = TrackRequest{} }
func (m *TrackRequest) String() string            { return proto.CompactTextString(m) }
func (*TrackRequest) ProtoMessage()               {}
func (*TrackRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TrackRequest) GetUser() int64 {
	if m != nil {
		return m.User
	}
	return 0
}

func (m *TrackRequest) GetPeriod() *TimePeriod {
	if m != nil {
		return m.Period
	}
	return nil
}

func (m *TrackRequest) GetTrack() string {
	if m != nil {
		return m.Track
	}
	return ""
}

type WrappedCoordinate struct {
	User           int64                       `protobuf:"varint,1,opt,name=user" json:"user,omitempty"`
	Id             int64                       `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Coord          *Coordinate                 `protobuf:"bytes,3,opt,name=coord" json:"coord,omitempty"`
	TimestampValue *google_protobuf2.Timestamp `protobuf:"bytes,4,opt,name=timestamp_value,json=timestampValue" json:"timestamp_value,omitempty"`
	Track          string                      `protobuf:"bytes,5,opt,name=track" json:"track,omitempty"`
}

func (m *WrappedCoordinate) Reset()                    { *m = WrappedCoordinate{} }
func (m *WrappedCoordinate) String() string            { return proto.CompactTextString(m) }
func (*WrappedCoordinate) ProtoMessage()               {}
func (*WrappedCoordinate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WrappedCoordinate) GetUser() int64 {
	if m != nil {
		return m.User
	}
	return 0
}

func (m *WrappedCoordinate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WrappedCoordinate) GetCoord() *Coordinate {
	if m != nil {
		return m.Coord
	}
	return nil
}

func (m *WrappedCoordinate) GetTimestampValue() *google_protobuf2.Timestamp {
	if m != nil {
		return m.TimestampValue
	}
	return nil
}

func (m *WrappedCoordinate) GetTrack() string {
	if m != nil {
		return m.Track
	}
	return ""
}

func init() {
	proto.RegisterType((*TimePeriod)(nil), "trac.TimePeriod")
	proto.RegisterType((*Coordinate)(nil), "trac.Coordinate")
	proto.RegisterType((*CoordinateRequest)(nil), "trac.CoordinateRequest")
	proto.RegisterType((*Track)(nil), "trac.Track")
	proto.RegisterType((*TrackRequest)(nil), "trac.TrackRequest")
	proto.RegisterType((*WrappedCoordinate)(nil), "trac.WrappedCoordinate")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Trac service

type TracClient interface {
	Post(ctx context.Context, in *WrappedCoordinate, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	GetLast(ctx context.Context, in *CoordinateRequest, opts ...grpc.CallOption) (*Coordinate, error)
	Get(ctx context.Context, in *TrackRequest, opts ...grpc.CallOption) (*Track, error)
}

type tracClient struct {
	cc *grpc.ClientConn
}

func NewTracClient(cc *grpc.ClientConn) TracClient {
	return &tracClient{cc}
}

func (c *tracClient) Post(ctx context.Context, in *WrappedCoordinate, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/trac.Trac/Post", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracClient) GetLast(ctx context.Context, in *CoordinateRequest, opts ...grpc.CallOption) (*Coordinate, error) {
	out := new(Coordinate)
	err := grpc.Invoke(ctx, "/trac.Trac/GetLast", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tracClient) Get(ctx context.Context, in *TrackRequest, opts ...grpc.CallOption) (*Track, error) {
	out := new(Track)
	err := grpc.Invoke(ctx, "/trac.Trac/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Trac service

type TracServer interface {
	Post(context.Context, *WrappedCoordinate) (*google_protobuf1.Empty, error)
	GetLast(context.Context, *CoordinateRequest) (*Coordinate, error)
	Get(context.Context, *TrackRequest) (*Track, error)
}

func RegisterTracServer(s *grpc.Server, srv TracServer) {
	s.RegisterService(&_Trac_serviceDesc, srv)
}

func _Trac_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrappedCoordinate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trac.Trac/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracServer).Post(ctx, req.(*WrappedCoordinate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trac_GetLast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoordinateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracServer).GetLast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trac.Trac/GetLast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracServer).GetLast(ctx, req.(*CoordinateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trac_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TracServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trac.Trac/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TracServer).Get(ctx, req.(*TrackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Trac_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trac.Trac",
	HandlerType: (*TracServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Post",
			Handler:    _Trac_Post_Handler,
		},
		{
			MethodName: "GetLast",
			Handler:    _Trac_GetLast_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Trac_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trac.proto",
}

func init() { proto.RegisterFile("trac.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0x3b, 0x0e, 0xed, 0x04, 0x5a, 0x32, 0x54, 0x6d, 0x30, 0x95, 0x88, 0xf6, 0x00,
	0x11, 0x42, 0x36, 0x84, 0x03, 0x12, 0x37, 0x5a, 0x50, 0x2e, 0x45, 0x0a, 0x56, 0x04, 0x12, 0x17,
	0xb4, 0x89, 0x97, 0x68, 0x55, 0xc7, 0xbb, 0xac, 0xd7, 0x15, 0x08, 0x71, 0xe1, 0x15, 0x78, 0x03,
	0x1e, 0x83, 0xd7, 0xe0, 0x15, 0xb8, 0xf0, 0x16, 0x68, 0xd7, 0x9b, 0xd8, 0x22, 0x88, 0x72, 0xf3,
	0xfa, 0x9f, 0xf9, 0x66, 0xfe, 0x99, 0x01, 0xd0, 0x8a, 0x2e, 0x62, 0xa9, 0x84, 0x16, 0xd8, 0x31,
	0xdf, 0xd1, 0xf1, 0x52, 0x88, 0x65, 0xce, 0x12, 0x2a, 0x79, 0x42, 0x8b, 0x42, 0x68, 0xaa, 0xb9,
	0x28, 0xca, 0x3a, 0x26, 0xba, 0xe5, 0x54, 0xfb, 0x9a, 0x57, 0xef, 0x12, 0xb6, 0x92, 0xfa, 0xa3,
	0x13, 0x6f, 0xff, 0x29, 0x6a, 0xbe, 0x62, 0xa5, 0xa6, 0x2b, 0xb9, 0xce, 0x56, 0xa2, 0xd2, 0x6c,
	0x59, 0xf1, 0x8c, 0x25, 0xcd, 0x67, 0x2d, 0x92, 0x1c, 0x60, 0xc6, 0x57, 0x6c, 0xca, 0x14, 0x17,
	0x19, 0x3e, 0x80, 0xb0, 0xd4, 0x54, 0xe9, 0x81, 0x37, 0xf4, 0x46, 0xbd, 0x71, 0x14, 0xd7, 0xec,
	0x78, 0xcd, 0x8e, 0x67, 0x6b, 0x76, 0x5a, 0x07, 0xe2, 0x7d, 0x08, 0x58, 0x91, 0x0d, 0xfc, 0x4b,
	0xe3, 0x4d, 0x18, 0x79, 0x09, 0x70, 0x2a, 0x84, 0xca, 0x78, 0x41, 0x35, 0xc3, 0x08, 0x76, 0x68,
	0xae, 0xb9, 0xae, 0x32, 0x66, 0x0b, 0xfa, 0xe9, 0xe6, 0x8d, 0x77, 0x21, 0x94, 0x82, 0x17, 0xda,
	0x91, 0xfb, 0x71, 0xab, 0xf3, 0xa9, 0x11, 0xd2, 0x5a, 0x27, 0x8f, 0xa1, 0xdf, 0x20, 0x53, 0xf6,
	0xbe, 0x62, 0xa5, 0x46, 0x84, 0x4e, 0x55, 0x32, 0x65, 0xa9, 0x41, 0x6a, 0xbf, 0x71, 0x0f, 0x7c,
	0x5e, 0x37, 0x1a, 0xa4, 0x3e, 0xcf, 0xc8, 0x53, 0x08, 0x67, 0x8a, 0x2e, 0xce, 0x9d, 0xe0, 0xad,
	0x05, 0x1c, 0x41, 0x77, 0x61, 0x88, 0xe5, 0xc0, 0x1f, 0x06, 0xa3, 0xde, 0xf8, 0x7a, 0x6c, 0xd7,
	0xd5, 0xaa, 0xe2, 0x74, 0x32, 0x87, 0xab, 0x16, 0xf1, 0xaf, 0xb2, 0x23, 0xe8, 0x4a, 0x3b, 0x5c,
	0xe7, 0xc4, 0xd1, 0x9a, 0xa1, 0xa7, 0x4e, 0xc7, 0x03, 0x08, 0x8d, 0x74, 0x3e, 0x08, 0x86, 0xde,
	0x68, 0x37, 0xad, 0x1f, 0xe4, 0xbb, 0x07, 0xfd, 0xd7, 0x8a, 0x4a, 0xc9, 0xb2, 0xd6, 0xe8, 0xfe,
	0xc3, 0x20, 0xde, 0x81, 0xd0, 0xf6, 0x69, 0x79, 0x7f, 0xb3, 0x51, 0xcb, 0x78, 0x0a, 0xfb, 0x9b,
	0x93, 0x79, 0x7b, 0x41, 0xf3, 0x8a, 0x0d, 0x3a, 0x97, 0xae, 0x73, 0x6f, 0x93, 0xf2, 0xca, 0x64,
	0x34, 0xcd, 0x87, 0xad, 0xe6, 0xc7, 0xbf, 0x3c, 0xe8, 0x98, 0x09, 0xe1, 0x19, 0x74, 0xa6, 0xa2,
	0xd4, 0x78, 0x54, 0x37, 0xb1, 0x65, 0x28, 0x3a, 0xdc, 0xaa, 0xf5, 0xdc, 0xdc, 0x38, 0xb9, 0xf1,
	0xe5, 0xc7, 0xcf, 0xaf, 0xfe, 0x35, 0xb2, 0x93, 0x5c, 0x3c, 0x4c, 0x4c, 0xee, 0x13, 0xef, 0x1e,
	0xbe, 0x80, 0x2b, 0x13, 0xa6, 0xcf, 0x68, 0x03, 0xdc, 0x3a, 0x81, 0x68, 0xcb, 0x2e, 0x39, 0xb2,
	0xa8, 0x3e, 0xee, 0xaf, 0x51, 0xc9, 0x27, 0x33, 0xb7, 0xcf, 0xf8, 0x0c, 0x82, 0x09, 0xd3, 0x88,
	0x6e, 0x33, 0xad, 0x8d, 0x46, 0xbd, 0xd6, 0x3f, 0x72, 0x6c, 0x01, 0x87, 0x78, 0xb0, 0x01, 0xe4,
	0xbc, 0xd4, 0x8e, 0x72, 0x32, 0x86, 0x9b, 0x5c, 0xc4, 0x4b, 0x25, 0x17, 0xad, 0x5b, 0x65, 0x1f,
	0xe8, 0x4a, 0xe6, 0xec, 0x64, 0xd7, 0x10, 0xa6, 0xc6, 0xda, 0xd4, 0x7b, 0xe3, 0xcb, 0xf9, 0x37,
	0x3f, 0x48, 0x67, 0x93, 0x79, 0xd7, 0xba, 0x7d, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x54, 0x98,
	0xb3, 0xf9, 0x11, 0x04, 0x00, 0x00,
}

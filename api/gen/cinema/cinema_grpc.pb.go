// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.1
// source: cinema.proto

package cinema

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CinemaServiceClient is the client API for CinemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CinemaServiceClient interface {
	ConfigureCinema(ctx context.Context, in *CinemaConfig, opts ...grpc.CallOption) (*StatusResponse, error)
	QueryReservedSeats(ctx context.Context, in *QueryReservedSeatsRequest, opts ...grpc.CallOption) (*QueryReservedSeatsResponse, error)
	ReserveSeats(ctx context.Context, in *SeatReservationRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	CancelSeats(ctx context.Context, in *SeatCancellationRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type cinemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCinemaServiceClient(cc grpc.ClientConnInterface) CinemaServiceClient {
	return &cinemaServiceClient{cc}
}

func (c *cinemaServiceClient) ConfigureCinema(ctx context.Context, in *CinemaConfig, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/cinema.CinemaService/ConfigureCinema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaServiceClient) QueryReservedSeats(ctx context.Context, in *QueryReservedSeatsRequest, opts ...grpc.CallOption) (*QueryReservedSeatsResponse, error) {
	out := new(QueryReservedSeatsResponse)
	err := c.cc.Invoke(ctx, "/cinema.CinemaService/QueryReservedSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaServiceClient) ReserveSeats(ctx context.Context, in *SeatReservationRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/cinema.CinemaService/ReserveSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cinemaServiceClient) CancelSeats(ctx context.Context, in *SeatCancellationRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/cinema.CinemaService/CancelSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CinemaServiceServer is the server API for CinemaService service.
// All implementations must embed UnimplementedCinemaServiceServer
// for forward compatibility
type CinemaServiceServer interface {
	ConfigureCinema(context.Context, *CinemaConfig) (*StatusResponse, error)
	QueryReservedSeats(context.Context, *QueryReservedSeatsRequest) (*QueryReservedSeatsResponse, error)
	ReserveSeats(context.Context, *SeatReservationRequest) (*StatusResponse, error)
	CancelSeats(context.Context, *SeatCancellationRequest) (*StatusResponse, error)
	mustEmbedUnimplementedCinemaServiceServer()
}

// UnimplementedCinemaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCinemaServiceServer struct {
}

func (UnimplementedCinemaServiceServer) ConfigureCinema(context.Context, *CinemaConfig) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureCinema not implemented")
}
func (UnimplementedCinemaServiceServer) QueryReservedSeats(context.Context, *QueryReservedSeatsRequest) (*QueryReservedSeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryReservedSeats not implemented")
}
func (UnimplementedCinemaServiceServer) ReserveSeats(context.Context, *SeatReservationRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveSeats not implemented")
}
func (UnimplementedCinemaServiceServer) CancelSeats(context.Context, *SeatCancellationRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelSeats not implemented")
}
func (UnimplementedCinemaServiceServer) mustEmbedUnimplementedCinemaServiceServer() {}

// UnsafeCinemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CinemaServiceServer will
// result in compilation errors.
type UnsafeCinemaServiceServer interface {
	mustEmbedUnimplementedCinemaServiceServer()
}

func RegisterCinemaServiceServer(s grpc.ServiceRegistrar, srv CinemaServiceServer) {
	s.RegisterService(&CinemaService_ServiceDesc, srv)
}

func _CinemaService_ConfigureCinema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CinemaConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServiceServer).ConfigureCinema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.CinemaService/ConfigureCinema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServiceServer).ConfigureCinema(ctx, req.(*CinemaConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _CinemaService_QueryReservedSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReservedSeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServiceServer).QueryReservedSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.CinemaService/QueryReservedSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServiceServer).QueryReservedSeats(ctx, req.(*QueryReservedSeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CinemaService_ReserveSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeatReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServiceServer).ReserveSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.CinemaService/ReserveSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServiceServer).ReserveSeats(ctx, req.(*SeatReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CinemaService_CancelSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeatCancellationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CinemaServiceServer).CancelSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cinema.CinemaService/CancelSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CinemaServiceServer).CancelSeats(ctx, req.(*SeatCancellationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CinemaService_ServiceDesc is the grpc.ServiceDesc for CinemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CinemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cinema.CinemaService",
	HandlerType: (*CinemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureCinema",
			Handler:    _CinemaService_ConfigureCinema_Handler,
		},
		{
			MethodName: "QueryReservedSeats",
			Handler:    _CinemaService_QueryReservedSeats_Handler,
		},
		{
			MethodName: "ReserveSeats",
			Handler:    _CinemaService_ReserveSeats_Handler,
		},
		{
			MethodName: "CancelSeats",
			Handler:    _CinemaService_CancelSeats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cinema.proto",
}

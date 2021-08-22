// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package query

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	QueryRange(ctx context.Context, in *QueryRangeRequest, opts ...grpc.CallOption) (*QueryRangeResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	Series(ctx context.Context, in *SeriesRequest, opts ...grpc.CallOption) (*SeriesResponse, error)
	Labels(ctx context.Context, in *LabelsRequest, opts ...grpc.CallOption) (*LabelsResponse, error)
	Values(ctx context.Context, in *ValuesRequest, opts ...grpc.CallOption) (*ValuesResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) QueryRange(ctx context.Context, in *QueryRangeRequest, opts ...grpc.CallOption) (*QueryRangeResponse, error) {
	out := new(QueryRangeResponse)
	err := c.cc.Invoke(ctx, "/parca.query.Query/QueryRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/parca.query.Query/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Series(ctx context.Context, in *SeriesRequest, opts ...grpc.CallOption) (*SeriesResponse, error) {
	out := new(SeriesResponse)
	err := c.cc.Invoke(ctx, "/parca.query.Query/Series", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Labels(ctx context.Context, in *LabelsRequest, opts ...grpc.CallOption) (*LabelsResponse, error) {
	out := new(LabelsResponse)
	err := c.cc.Invoke(ctx, "/parca.query.Query/Labels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Values(ctx context.Context, in *ValuesRequest, opts ...grpc.CallOption) (*ValuesResponse, error) {
	out := new(ValuesResponse)
	err := c.cc.Invoke(ctx, "/parca.query.Query/Values", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations should embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	QueryRange(context.Context, *QueryRangeRequest) (*QueryRangeResponse, error)
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Series(context.Context, *SeriesRequest) (*SeriesResponse, error)
	Labels(context.Context, *LabelsRequest) (*LabelsResponse, error)
	Values(context.Context, *ValuesRequest) (*ValuesResponse, error)
}

// UnimplementedQueryServer should be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) QueryRange(context.Context, *QueryRangeRequest) (*QueryRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRange not implemented")
}
func (UnimplementedQueryServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedQueryServer) Series(context.Context, *SeriesRequest) (*SeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Series not implemented")
}
func (UnimplementedQueryServer) Labels(context.Context, *LabelsRequest) (*LabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Labels not implemented")
}
func (UnimplementedQueryServer) Values(context.Context, *ValuesRequest) (*ValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Values not implemented")
}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_QueryRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parca.query.Query/QueryRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryRange(ctx, req.(*QueryRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parca.query.Query/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Series_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Series(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parca.query.Query/Series",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Series(ctx, req.(*SeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Labels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Labels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parca.query.Query/Labels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Labels(ctx, req.(*LabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Values_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Values(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parca.query.Query/Values",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Values(ctx, req.(*ValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parca.query.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryRange",
			Handler:    _Query_QueryRange_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Query_Query_Handler,
		},
		{
			MethodName: "Series",
			Handler:    _Query_Series_Handler,
		},
		{
			MethodName: "Labels",
			Handler:    _Query_Labels_Handler,
		},
		{
			MethodName: "Values",
			Handler:    _Query_Values_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query/query.proto",
}

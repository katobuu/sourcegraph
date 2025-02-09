// 🔔 IMPORTANT: Be VERY careful not to introduce breaking changes to this
// spec - raw protocol buffer wire format messages are persisted to database
// as a cache, and Sourcegraph instances rely on this format to emit telemetry
// to the managed Sourcegraph Telemetry Gateway service.
//
// Tests in ./internal/telemetrygateway/v1/backcompat_test.go can be used to
// assert compatibility with snapshots created by older versions of this spec.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: telemetrygateway.proto

package v1

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

const (
	TelemeteryGatewayService_RecordEvents_FullMethodName = "/telemetrygateway.v1.TelemeteryGatewayService/RecordEvents"
)

// TelemeteryGatewayServiceClient is the client API for TelemeteryGatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelemeteryGatewayServiceClient interface {
	// RecordEvents streams telemetry events in batches to the Telemetry Gateway
	// service. Events should only be considered delivered if recording is
	// acknowledged in RecordEventsResponse.
	//
	// 🚨 SECURITY: Callers should check the attributes of the Event type to ensure
	// that only the appropriate fields are exported, as some fields should only
	// be exported on an allowlist basis.
	RecordEvents(ctx context.Context, opts ...grpc.CallOption) (TelemeteryGatewayService_RecordEventsClient, error)
}

type telemeteryGatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTelemeteryGatewayServiceClient(cc grpc.ClientConnInterface) TelemeteryGatewayServiceClient {
	return &telemeteryGatewayServiceClient{cc}
}

func (c *telemeteryGatewayServiceClient) RecordEvents(ctx context.Context, opts ...grpc.CallOption) (TelemeteryGatewayService_RecordEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TelemeteryGatewayService_ServiceDesc.Streams[0], TelemeteryGatewayService_RecordEvents_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &telemeteryGatewayServiceRecordEventsClient{stream}
	return x, nil
}

type TelemeteryGatewayService_RecordEventsClient interface {
	Send(*RecordEventsRequest) error
	Recv() (*RecordEventsResponse, error)
	grpc.ClientStream
}

type telemeteryGatewayServiceRecordEventsClient struct {
	grpc.ClientStream
}

func (x *telemeteryGatewayServiceRecordEventsClient) Send(m *RecordEventsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *telemeteryGatewayServiceRecordEventsClient) Recv() (*RecordEventsResponse, error) {
	m := new(RecordEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TelemeteryGatewayServiceServer is the server API for TelemeteryGatewayService service.
// All implementations must embed UnimplementedTelemeteryGatewayServiceServer
// for forward compatibility
type TelemeteryGatewayServiceServer interface {
	// RecordEvents streams telemetry events in batches to the Telemetry Gateway
	// service. Events should only be considered delivered if recording is
	// acknowledged in RecordEventsResponse.
	//
	// 🚨 SECURITY: Callers should check the attributes of the Event type to ensure
	// that only the appropriate fields are exported, as some fields should only
	// be exported on an allowlist basis.
	RecordEvents(TelemeteryGatewayService_RecordEventsServer) error
	mustEmbedUnimplementedTelemeteryGatewayServiceServer()
}

// UnimplementedTelemeteryGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTelemeteryGatewayServiceServer struct {
}

func (UnimplementedTelemeteryGatewayServiceServer) RecordEvents(TelemeteryGatewayService_RecordEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordEvents not implemented")
}
func (UnimplementedTelemeteryGatewayServiceServer) mustEmbedUnimplementedTelemeteryGatewayServiceServer() {
}

// UnsafeTelemeteryGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelemeteryGatewayServiceServer will
// result in compilation errors.
type UnsafeTelemeteryGatewayServiceServer interface {
	mustEmbedUnimplementedTelemeteryGatewayServiceServer()
}

func RegisterTelemeteryGatewayServiceServer(s grpc.ServiceRegistrar, srv TelemeteryGatewayServiceServer) {
	s.RegisterService(&TelemeteryGatewayService_ServiceDesc, srv)
}

func _TelemeteryGatewayService_RecordEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TelemeteryGatewayServiceServer).RecordEvents(&telemeteryGatewayServiceRecordEventsServer{stream})
}

type TelemeteryGatewayService_RecordEventsServer interface {
	Send(*RecordEventsResponse) error
	Recv() (*RecordEventsRequest, error)
	grpc.ServerStream
}

type telemeteryGatewayServiceRecordEventsServer struct {
	grpc.ServerStream
}

func (x *telemeteryGatewayServiceRecordEventsServer) Send(m *RecordEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *telemeteryGatewayServiceRecordEventsServer) Recv() (*RecordEventsRequest, error) {
	m := new(RecordEventsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TelemeteryGatewayService_ServiceDesc is the grpc.ServiceDesc for TelemeteryGatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelemeteryGatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "telemetrygateway.v1.TelemeteryGatewayService",
	HandlerType: (*TelemeteryGatewayServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordEvents",
			Handler:       _TelemeteryGatewayService_RecordEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "telemetrygateway.proto",
}

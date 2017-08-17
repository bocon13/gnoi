// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openconfig/gnoi/layer2/layer2.proto

/*
Package gnoi_layer2 is a generated protocol buffer package.

It is generated from these files:
	github.com/openconfig/gnoi/layer2/layer2.proto

It has these top-level messages:
	ClearNeighborDiscoveryRequest
	ClearNeighborDiscoveryResponse
	ClearSpanningTreeRequest
	ClearSpanningTreeResponse
	PerformBERTRequest
	PerformBERTResponse
	ClearLLDPInterfaceRequest
	ClearLLDPInterfaceResponse
	SendWakeOnLANRequest
	SendWakeOnLANResponse
*/
package gnoi_layer2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import gnoi "github.com/openconfig/reference/rpc/gnoi"

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

type PerformBERTResponse_BERTState int32

const (
	PerformBERTResponse_UNKNOWN  PerformBERTResponse_BERTState = 0
	PerformBERTResponse_DISABLED PerformBERTResponse_BERTState = 1
	PerformBERTResponse_RUNNING  PerformBERTResponse_BERTState = 2
	PerformBERTResponse_COMPLETE PerformBERTResponse_BERTState = 3
	PerformBERTResponse_ERROR    PerformBERTResponse_BERTState = 4
)

var PerformBERTResponse_BERTState_name = map[int32]string{
	0: "UNKNOWN",
	1: "DISABLED",
	2: "RUNNING",
	3: "COMPLETE",
	4: "ERROR",
}
var PerformBERTResponse_BERTState_value = map[string]int32{
	"UNKNOWN":  0,
	"DISABLED": 1,
	"RUNNING":  2,
	"COMPLETE": 3,
	"ERROR":    4,
}

func (x PerformBERTResponse_BERTState) String() string {
	return proto.EnumName(PerformBERTResponse_BERTState_name, int32(x))
}
func (PerformBERTResponse_BERTState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type ClearNeighborDiscoveryRequest struct {
	Protocol gnoi.L3Protocol `protobuf:"varint,1,opt,name=protocol,enum=gnoi.L3Protocol" json:"protocol,omitempty"`
	Address  string          `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *ClearNeighborDiscoveryRequest) Reset()                    { *m = ClearNeighborDiscoveryRequest{} }
func (m *ClearNeighborDiscoveryRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearNeighborDiscoveryRequest) ProtoMessage()               {}
func (*ClearNeighborDiscoveryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClearNeighborDiscoveryRequest) GetProtocol() gnoi.L3Protocol {
	if m != nil {
		return m.Protocol
	}
	return gnoi.L3Protocol_UNSPECIFIED
}

func (m *ClearNeighborDiscoveryRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ClearNeighborDiscoveryResponse struct {
}

func (m *ClearNeighborDiscoveryResponse) Reset()                    { *m = ClearNeighborDiscoveryResponse{} }
func (m *ClearNeighborDiscoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearNeighborDiscoveryResponse) ProtoMessage()               {}
func (*ClearNeighborDiscoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ClearSpanningTreeRequest struct {
	Interface *gnoi.Path `protobuf:"bytes,1,opt,name=interface" json:"interface,omitempty"`
}

func (m *ClearSpanningTreeRequest) Reset()                    { *m = ClearSpanningTreeRequest{} }
func (m *ClearSpanningTreeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearSpanningTreeRequest) ProtoMessage()               {}
func (*ClearSpanningTreeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClearSpanningTreeRequest) GetInterface() *gnoi.Path {
	if m != nil {
		return m.Interface
	}
	return nil
}

type ClearSpanningTreeResponse struct {
}

func (m *ClearSpanningTreeResponse) Reset()                    { *m = ClearSpanningTreeResponse{} }
func (m *ClearSpanningTreeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearSpanningTreeResponse) ProtoMessage()               {}
func (*ClearSpanningTreeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type PerformBERTRequest struct {
	// ID for retrieving a previous BERT run data - optional.
	Id        string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Interface *gnoi.Path `protobuf:"bytes,2,opt,name=interface" json:"interface,omitempty"`
}

func (m *PerformBERTRequest) Reset()                    { *m = PerformBERTRequest{} }
func (m *PerformBERTRequest) String() string            { return proto.CompactTextString(m) }
func (*PerformBERTRequest) ProtoMessage()               {}
func (*PerformBERTRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PerformBERTRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PerformBERTRequest) GetInterface() *gnoi.Path {
	if m != nil {
		return m.Interface
	}
	return nil
}

type PerformBERTResponse struct {
	Id            string                        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	State         PerformBERTResponse_BERTState `protobuf:"varint,2,opt,name=state,enum=gnoi.layer2.PerformBERTResponse_BERTState" json:"state,omitempty"`
	ElapsedPeriod int64                         `protobuf:"varint,3,opt,name=elapsed_period,json=elapsedPeriod" json:"elapsed_period,omitempty"`
	Pattern       []byte                        `protobuf:"bytes,4,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Errors        int64                         `protobuf:"varint,5,opt,name=errors" json:"errors,omitempty"`
	ReceivedBits  int64                         `protobuf:"varint,6,opt,name=received_bits,json=receivedBits" json:"received_bits,omitempty"`
}

func (m *PerformBERTResponse) Reset()                    { *m = PerformBERTResponse{} }
func (m *PerformBERTResponse) String() string            { return proto.CompactTextString(m) }
func (*PerformBERTResponse) ProtoMessage()               {}
func (*PerformBERTResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PerformBERTResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PerformBERTResponse) GetState() PerformBERTResponse_BERTState {
	if m != nil {
		return m.State
	}
	return PerformBERTResponse_UNKNOWN
}

func (m *PerformBERTResponse) GetElapsedPeriod() int64 {
	if m != nil {
		return m.ElapsedPeriod
	}
	return 0
}

func (m *PerformBERTResponse) GetPattern() []byte {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *PerformBERTResponse) GetErrors() int64 {
	if m != nil {
		return m.Errors
	}
	return 0
}

func (m *PerformBERTResponse) GetReceivedBits() int64 {
	if m != nil {
		return m.ReceivedBits
	}
	return 0
}

type ClearLLDPInterfaceRequest struct {
	Interface *gnoi.Path `protobuf:"bytes,1,opt,name=interface" json:"interface,omitempty"`
}

func (m *ClearLLDPInterfaceRequest) Reset()                    { *m = ClearLLDPInterfaceRequest{} }
func (m *ClearLLDPInterfaceRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearLLDPInterfaceRequest) ProtoMessage()               {}
func (*ClearLLDPInterfaceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ClearLLDPInterfaceRequest) GetInterface() *gnoi.Path {
	if m != nil {
		return m.Interface
	}
	return nil
}

type ClearLLDPInterfaceResponse struct {
}

func (m *ClearLLDPInterfaceResponse) Reset()                    { *m = ClearLLDPInterfaceResponse{} }
func (m *ClearLLDPInterfaceResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearLLDPInterfaceResponse) ProtoMessage()               {}
func (*ClearLLDPInterfaceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type SendWakeOnLANRequest struct {
	Interface  *gnoi.Path `protobuf:"bytes,1,opt,name=interface" json:"interface,omitempty"`
	Address    string     `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	MacAddress []byte     `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
}

func (m *SendWakeOnLANRequest) Reset()                    { *m = SendWakeOnLANRequest{} }
func (m *SendWakeOnLANRequest) String() string            { return proto.CompactTextString(m) }
func (*SendWakeOnLANRequest) ProtoMessage()               {}
func (*SendWakeOnLANRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SendWakeOnLANRequest) GetInterface() *gnoi.Path {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *SendWakeOnLANRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SendWakeOnLANRequest) GetMacAddress() []byte {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

type SendWakeOnLANResponse struct {
}

func (m *SendWakeOnLANResponse) Reset()                    { *m = SendWakeOnLANResponse{} }
func (m *SendWakeOnLANResponse) String() string            { return proto.CompactTextString(m) }
func (*SendWakeOnLANResponse) ProtoMessage()               {}
func (*SendWakeOnLANResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*ClearNeighborDiscoveryRequest)(nil), "gnoi.layer2.ClearNeighborDiscoveryRequest")
	proto.RegisterType((*ClearNeighborDiscoveryResponse)(nil), "gnoi.layer2.ClearNeighborDiscoveryResponse")
	proto.RegisterType((*ClearSpanningTreeRequest)(nil), "gnoi.layer2.ClearSpanningTreeRequest")
	proto.RegisterType((*ClearSpanningTreeResponse)(nil), "gnoi.layer2.ClearSpanningTreeResponse")
	proto.RegisterType((*PerformBERTRequest)(nil), "gnoi.layer2.PerformBERTRequest")
	proto.RegisterType((*PerformBERTResponse)(nil), "gnoi.layer2.PerformBERTResponse")
	proto.RegisterType((*ClearLLDPInterfaceRequest)(nil), "gnoi.layer2.ClearLLDPInterfaceRequest")
	proto.RegisterType((*ClearLLDPInterfaceResponse)(nil), "gnoi.layer2.ClearLLDPInterfaceResponse")
	proto.RegisterType((*SendWakeOnLANRequest)(nil), "gnoi.layer2.SendWakeOnLANRequest")
	proto.RegisterType((*SendWakeOnLANResponse)(nil), "gnoi.layer2.SendWakeOnLANResponse")
	proto.RegisterEnum("gnoi.layer2.PerformBERTResponse_BERTState", PerformBERTResponse_BERTState_name, PerformBERTResponse_BERTState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Layer2 service

type Layer2Client interface {
	// ClearNeighborDiscovery will clear either a specific neighbor entry or
	// clear the entire table based on parameters provided.
	// TODO: This method is subject to deprecation once OpenConfig models this
	// state.
	ClearNeighborDiscovery(ctx context.Context, in *ClearNeighborDiscoveryRequest, opts ...grpc.CallOption) (*ClearNeighborDiscoveryResponse, error)
	// ClearSpanningTree will reset a blocked spanning tree interface.
	// TODO: This method is subject to deprecation once OpenConfig models this
	// state.
	ClearSpanningTree(ctx context.Context, in *ClearSpanningTreeRequest, opts ...grpc.CallOption) (*ClearSpanningTreeResponse, error)
	// PerformBERT will perform a BERT operation on a port. The stream will
	// return the current state of the operation as well as the ID for the
	// operation.
	PerformBERT(ctx context.Context, in *PerformBERTRequest, opts ...grpc.CallOption) (Layer2_PerformBERTClient, error)
	// ClearLLDPInterface will clear all LLDP adjacencies on the provided
	// interface.
	ClearLLDPInterface(ctx context.Context, in *ClearLLDPInterfaceRequest, opts ...grpc.CallOption) (*ClearLLDPInterfaceResponse, error)
	// SendWakeOnLAN will send a WOL event on the requested interface.
	SendWakeOnLAN(ctx context.Context, in *SendWakeOnLANRequest, opts ...grpc.CallOption) (*SendWakeOnLANResponse, error)
}

type layer2Client struct {
	cc *grpc.ClientConn
}

func NewLayer2Client(cc *grpc.ClientConn) Layer2Client {
	return &layer2Client{cc}
}

func (c *layer2Client) ClearNeighborDiscovery(ctx context.Context, in *ClearNeighborDiscoveryRequest, opts ...grpc.CallOption) (*ClearNeighborDiscoveryResponse, error) {
	out := new(ClearNeighborDiscoveryResponse)
	err := grpc.Invoke(ctx, "/gnoi.layer2.Layer2/ClearNeighborDiscovery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layer2Client) ClearSpanningTree(ctx context.Context, in *ClearSpanningTreeRequest, opts ...grpc.CallOption) (*ClearSpanningTreeResponse, error) {
	out := new(ClearSpanningTreeResponse)
	err := grpc.Invoke(ctx, "/gnoi.layer2.Layer2/ClearSpanningTree", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layer2Client) PerformBERT(ctx context.Context, in *PerformBERTRequest, opts ...grpc.CallOption) (Layer2_PerformBERTClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Layer2_serviceDesc.Streams[0], c.cc, "/gnoi.layer2.Layer2/PerformBERT", opts...)
	if err != nil {
		return nil, err
	}
	x := &layer2PerformBERTClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Layer2_PerformBERTClient interface {
	Recv() (*PerformBERTResponse, error)
	grpc.ClientStream
}

type layer2PerformBERTClient struct {
	grpc.ClientStream
}

func (x *layer2PerformBERTClient) Recv() (*PerformBERTResponse, error) {
	m := new(PerformBERTResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *layer2Client) ClearLLDPInterface(ctx context.Context, in *ClearLLDPInterfaceRequest, opts ...grpc.CallOption) (*ClearLLDPInterfaceResponse, error) {
	out := new(ClearLLDPInterfaceResponse)
	err := grpc.Invoke(ctx, "/gnoi.layer2.Layer2/ClearLLDPInterface", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *layer2Client) SendWakeOnLAN(ctx context.Context, in *SendWakeOnLANRequest, opts ...grpc.CallOption) (*SendWakeOnLANResponse, error) {
	out := new(SendWakeOnLANResponse)
	err := grpc.Invoke(ctx, "/gnoi.layer2.Layer2/SendWakeOnLAN", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Layer2 service

type Layer2Server interface {
	// ClearNeighborDiscovery will clear either a specific neighbor entry or
	// clear the entire table based on parameters provided.
	// TODO: This method is subject to deprecation once OpenConfig models this
	// state.
	ClearNeighborDiscovery(context.Context, *ClearNeighborDiscoveryRequest) (*ClearNeighborDiscoveryResponse, error)
	// ClearSpanningTree will reset a blocked spanning tree interface.
	// TODO: This method is subject to deprecation once OpenConfig models this
	// state.
	ClearSpanningTree(context.Context, *ClearSpanningTreeRequest) (*ClearSpanningTreeResponse, error)
	// PerformBERT will perform a BERT operation on a port. The stream will
	// return the current state of the operation as well as the ID for the
	// operation.
	PerformBERT(*PerformBERTRequest, Layer2_PerformBERTServer) error
	// ClearLLDPInterface will clear all LLDP adjacencies on the provided
	// interface.
	ClearLLDPInterface(context.Context, *ClearLLDPInterfaceRequest) (*ClearLLDPInterfaceResponse, error)
	// SendWakeOnLAN will send a WOL event on the requested interface.
	SendWakeOnLAN(context.Context, *SendWakeOnLANRequest) (*SendWakeOnLANResponse, error)
}

func RegisterLayer2Server(s *grpc.Server, srv Layer2Server) {
	s.RegisterService(&_Layer2_serviceDesc, srv)
}

func _Layer2_ClearNeighborDiscovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearNeighborDiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Layer2Server).ClearNeighborDiscovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.layer2.Layer2/ClearNeighborDiscovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Layer2Server).ClearNeighborDiscovery(ctx, req.(*ClearNeighborDiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Layer2_ClearSpanningTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearSpanningTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Layer2Server).ClearSpanningTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.layer2.Layer2/ClearSpanningTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Layer2Server).ClearSpanningTree(ctx, req.(*ClearSpanningTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Layer2_PerformBERT_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PerformBERTRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Layer2Server).PerformBERT(m, &layer2PerformBERTServer{stream})
}

type Layer2_PerformBERTServer interface {
	Send(*PerformBERTResponse) error
	grpc.ServerStream
}

type layer2PerformBERTServer struct {
	grpc.ServerStream
}

func (x *layer2PerformBERTServer) Send(m *PerformBERTResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Layer2_ClearLLDPInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearLLDPInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Layer2Server).ClearLLDPInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.layer2.Layer2/ClearLLDPInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Layer2Server).ClearLLDPInterface(ctx, req.(*ClearLLDPInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Layer2_SendWakeOnLAN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendWakeOnLANRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Layer2Server).SendWakeOnLAN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.layer2.Layer2/SendWakeOnLAN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Layer2Server).SendWakeOnLAN(ctx, req.(*SendWakeOnLANRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Layer2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.layer2.Layer2",
	HandlerType: (*Layer2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClearNeighborDiscovery",
			Handler:    _Layer2_ClearNeighborDiscovery_Handler,
		},
		{
			MethodName: "ClearSpanningTree",
			Handler:    _Layer2_ClearSpanningTree_Handler,
		},
		{
			MethodName: "ClearLLDPInterface",
			Handler:    _Layer2_ClearLLDPInterface_Handler,
		},
		{
			MethodName: "SendWakeOnLAN",
			Handler:    _Layer2_SendWakeOnLAN_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PerformBERT",
			Handler:       _Layer2_PerformBERT_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnoi/layer2/layer2.proto",
}

func init() { proto.RegisterFile("github.com/openconfig/gnoi/layer2/layer2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x6b, 0x6f, 0x12, 0x4d,
	0x14, 0xc7, 0xbb, 0xd0, 0xd2, 0x72, 0x68, 0x09, 0xcf, 0x3c, 0x5a, 0x57, 0xbc, 0x14, 0xd7, 0x54,
	0x89, 0x9a, 0xc5, 0x50, 0x3f, 0x80, 0x6d, 0x21, 0xa6, 0x71, 0x5d, 0xc8, 0x40, 0x53, 0xdf, 0x35,
	0xcb, 0xee, 0x61, 0x3b, 0x11, 0x66, 0xd6, 0x99, 0x69, 0x93, 0xc6, 0x8f, 0xe5, 0xd7, 0xf3, 0x85,
	0x61, 0x2f, 0x58, 0x64, 0x5b, 0xec, 0xab, 0xcd, 0x9c, 0xcb, 0xef, 0x7f, 0x72, 0x2e, 0x0b, 0x76,
	0xc8, 0xf4, 0xc5, 0xe5, 0xc8, 0xf6, 0xc5, 0xb4, 0x25, 0x22, 0xe4, 0xbe, 0xe0, 0x63, 0x16, 0xb6,
	0x42, 0x2e, 0x58, 0x6b, 0xe2, 0x5d, 0xa3, 0x6c, 0xa7, 0x1f, 0x3b, 0x92, 0x42, 0x0b, 0x52, 0x99,
	0x79, 0xec, 0xc4, 0x54, 0xff, 0x90, 0x9f, 0x2c, 0x71, 0x8c, 0x12, 0xb9, 0x8f, 0x2d, 0x19, 0xf9,
	0x09, 0x4a, 0x5f, 0x47, 0xa8, 0x12, 0x84, 0x15, 0xc2, 0xb3, 0xe3, 0x09, 0x7a, 0xd2, 0x45, 0x16,
	0x5e, 0x8c, 0x84, 0xec, 0x30, 0xe5, 0x8b, 0x2b, 0x94, 0xd7, 0x14, 0xbf, 0x5f, 0xa2, 0xd2, 0xe4,
	0x1d, 0x6c, 0xc5, 0x91, 0xbe, 0x98, 0x98, 0x46, 0xc3, 0x68, 0x56, 0xdb, 0x35, 0x3b, 0x96, 0x75,
	0x0e, 0xfa, 0xa9, 0x9d, 0xce, 0x23, 0x88, 0x09, 0x9b, 0x5e, 0x10, 0x48, 0x54, 0xca, 0x2c, 0x34,
	0x8c, 0x66, 0x99, 0x66, 0x4f, 0xab, 0x01, 0xcf, 0x6f, 0x13, 0x52, 0x91, 0xe0, 0x0a, 0xad, 0x0e,
	0x98, 0x71, 0xc4, 0x20, 0xf2, 0x38, 0x67, 0x3c, 0x1c, 0x4a, 0xc4, 0xac, 0x8a, 0x26, 0x94, 0x19,
	0xd7, 0x28, 0xc7, 0x9e, 0x8f, 0x71, 0x19, 0x95, 0x36, 0x24, 0x65, 0xf4, 0x3d, 0x7d, 0x41, 0xff,
	0x38, 0xad, 0x27, 0xf0, 0x38, 0x87, 0x92, 0x4a, 0xb8, 0x40, 0xfa, 0x28, 0xc7, 0x42, 0x4e, 0x8f,
	0xba, 0x74, 0x98, 0xc1, 0xab, 0x50, 0x60, 0x41, 0x4c, 0x2d, 0xd3, 0x02, 0x0b, 0x16, 0xc5, 0x0a,
	0x77, 0x89, 0xfd, 0x2c, 0xc0, 0xff, 0x0b, 0xc0, 0x44, 0x67, 0x89, 0xf8, 0x11, 0x36, 0x94, 0xf6,
	0x74, 0x42, 0xab, 0xb6, 0xdf, 0xd8, 0x37, 0x06, 0x67, 0xe7, 0x00, 0xec, 0xd9, 0x63, 0x30, 0xcb,
	0xa0, 0x49, 0x22, 0xd9, 0x87, 0x2a, 0x4e, 0xbc, 0x48, 0x61, 0x70, 0x1e, 0xa1, 0x64, 0x22, 0x30,
	0x8b, 0x0d, 0xa3, 0x59, 0xa4, 0x3b, 0xa9, 0xb5, 0x1f, 0x1b, 0x67, 0xfd, 0x8f, 0x3c, 0xad, 0x51,
	0x72, 0x73, 0xbd, 0x61, 0x34, 0xb7, 0x69, 0xf6, 0x24, 0xbb, 0x50, 0x42, 0x29, 0x85, 0x54, 0xe6,
	0x46, 0x9c, 0x98, 0xbe, 0xc8, 0x4b, 0xd8, 0x91, 0xe8, 0x23, 0xbb, 0xc2, 0xe0, 0x7c, 0xc4, 0xb4,
	0x32, 0x4b, 0xb1, 0x7b, 0x3b, 0x33, 0x1e, 0x31, 0xad, 0x2c, 0x07, 0xca, 0xf3, 0x8a, 0x48, 0x05,
	0x36, 0x4f, 0xdd, 0xcf, 0x6e, 0xef, 0xcc, 0xad, 0xad, 0x91, 0x6d, 0xd8, 0xea, 0x9c, 0x0c, 0x0e,
	0x8f, 0x9c, 0x6e, 0xa7, 0x66, 0xcc, 0x5c, 0xf4, 0xd4, 0x75, 0x4f, 0xdc, 0x4f, 0xb5, 0xc2, 0xcc,
	0x75, 0xdc, 0xfb, 0xd2, 0x77, 0xba, 0xc3, 0x6e, 0xad, 0x48, 0xca, 0xb0, 0xd1, 0xa5, 0xb4, 0x47,
	0x6b, 0xeb, 0x56, 0x37, 0x1d, 0x91, 0xe3, 0x74, 0xfa, 0x27, 0x59, 0x2f, 0xef, 0x3f, 0xe9, 0xa7,
	0x50, 0xcf, 0xc3, 0xa4, 0xa3, 0xfe, 0x01, 0x0f, 0x06, 0xc8, 0x83, 0x33, 0xef, 0x1b, 0xf6, 0xb8,
	0x73, 0xe8, 0xde, 0x9b, 0x7f, 0xfb, 0x2e, 0x93, 0x3d, 0xa8, 0x4c, 0x3d, 0xff, 0x3c, 0xf3, 0x16,
	0xe3, 0x4e, 0xc3, 0xd4, 0xf3, 0x0f, 0xd3, 0x65, 0x7f, 0x04, 0x0f, 0xff, 0x12, 0x4f, 0xaa, 0x6a,
	0xff, 0x2a, 0x42, 0xc9, 0x89, 0xc7, 0x4e, 0x14, 0xec, 0xe6, 0x1f, 0x04, 0x59, 0x5c, 0x8f, 0x3b,
	0xcf, 0xb3, 0xfe, 0xf6, 0x9f, 0x62, 0xd3, 0x9e, 0xac, 0x91, 0x00, 0xfe, 0x5b, 0xba, 0x0e, 0xb2,
	0xbf, 0xcc, 0xc8, 0xb9, 0xc1, 0xfa, 0xab, 0x55, 0x61, 0x73, 0x95, 0x21, 0x54, 0x6e, 0x2c, 0x35,
	0xd9, 0xbb, 0x7d, 0xdd, 0x13, 0x72, 0x63, 0xd5, 0x3d, 0x58, 0x6b, 0xef, 0x0d, 0x12, 0x02, 0x59,
	0x9e, 0x37, 0xc9, 0xa9, 0x2a, 0x6f, 0xaf, 0xea, 0xaf, 0x57, 0xc6, 0xcd, 0xcb, 0xff, 0x0a, 0x3b,
	0x0b, 0xd3, 0x23, 0x2f, 0x16, 0x72, 0xf3, 0xd6, 0xaa, 0x6e, 0xdd, 0x15, 0x92, 0x91, 0x47, 0xa5,
	0xf8, 0x47, 0x79, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xac, 0xd7, 0xcd, 0xe9, 0x05, 0x00,
	0x00,
}
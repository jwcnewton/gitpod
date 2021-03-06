// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NewuidmapRequest struct {
	Pid                  int64                       `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Gid                  bool                        `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	Mapping              []*NewuidmapRequest_Mapping `protobuf:"bytes,3,rep,name=mapping,proto3" json:"mapping,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *NewuidmapRequest) Reset()         { *m = NewuidmapRequest{} }
func (m *NewuidmapRequest) String() string { return proto.CompactTextString(m) }
func (*NewuidmapRequest) ProtoMessage()    {}
func (*NewuidmapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0}
}

func (m *NewuidmapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewuidmapRequest.Unmarshal(m, b)
}
func (m *NewuidmapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewuidmapRequest.Marshal(b, m, deterministic)
}
func (m *NewuidmapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewuidmapRequest.Merge(m, src)
}
func (m *NewuidmapRequest) XXX_Size() int {
	return xxx_messageInfo_NewuidmapRequest.Size(m)
}
func (m *NewuidmapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewuidmapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewuidmapRequest proto.InternalMessageInfo

func (m *NewuidmapRequest) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *NewuidmapRequest) GetGid() bool {
	if m != nil {
		return m.Gid
	}
	return false
}

func (m *NewuidmapRequest) GetMapping() []*NewuidmapRequest_Mapping {
	if m != nil {
		return m.Mapping
	}
	return nil
}

type NewuidmapRequest_Mapping struct {
	ContainerId          uint32   `protobuf:"varint,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	HostId               uint32   `protobuf:"varint,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	Size                 uint32   `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewuidmapRequest_Mapping) Reset()         { *m = NewuidmapRequest_Mapping{} }
func (m *NewuidmapRequest_Mapping) String() string { return proto.CompactTextString(m) }
func (*NewuidmapRequest_Mapping) ProtoMessage()    {}
func (*NewuidmapRequest_Mapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0, 0}
}

func (m *NewuidmapRequest_Mapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewuidmapRequest_Mapping.Unmarshal(m, b)
}
func (m *NewuidmapRequest_Mapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewuidmapRequest_Mapping.Marshal(b, m, deterministic)
}
func (m *NewuidmapRequest_Mapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewuidmapRequest_Mapping.Merge(m, src)
}
func (m *NewuidmapRequest_Mapping) XXX_Size() int {
	return xxx_messageInfo_NewuidmapRequest_Mapping.Size(m)
}
func (m *NewuidmapRequest_Mapping) XXX_DiscardUnknown() {
	xxx_messageInfo_NewuidmapRequest_Mapping.DiscardUnknown(m)
}

var xxx_messageInfo_NewuidmapRequest_Mapping proto.InternalMessageInfo

func (m *NewuidmapRequest_Mapping) GetContainerId() uint32 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *NewuidmapRequest_Mapping) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *NewuidmapRequest_Mapping) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type NewuidmapResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewuidmapResponse) Reset()         { *m = NewuidmapResponse{} }
func (m *NewuidmapResponse) String() string { return proto.CompactTextString(m) }
func (*NewuidmapResponse) ProtoMessage()    {}
func (*NewuidmapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{1}
}

func (m *NewuidmapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewuidmapResponse.Unmarshal(m, b)
}
func (m *NewuidmapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewuidmapResponse.Marshal(b, m, deterministic)
}
func (m *NewuidmapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewuidmapResponse.Merge(m, src)
}
func (m *NewuidmapResponse) XXX_Size() int {
	return xxx_messageInfo_NewuidmapResponse.Size(m)
}
func (m *NewuidmapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewuidmapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewuidmapResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NewuidmapRequest)(nil), "supervisor.NewuidmapRequest")
	proto.RegisterType((*NewuidmapRequest_Mapping)(nil), "supervisor.NewuidmapRequest.Mapping")
	proto.RegisterType((*NewuidmapResponse)(nil), "supervisor.NewuidmapResponse")
}

func init() {
	proto.RegisterFile("control.proto", fileDescriptor_0c5120591600887d)
}

var fileDescriptor_0c5120591600887d = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x2e, 0x2d, 0x48, 0x2d,
	0x2a, 0xcb, 0x2c, 0xce, 0x2f, 0x52, 0xba, 0xc8, 0xc8, 0x25, 0xe0, 0x97, 0x5a, 0x5e, 0x9a, 0x99,
	0x92, 0x9b, 0x58, 0x10, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc0, 0xc5, 0x5c, 0x90,
	0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0x62, 0x82, 0x44, 0xd2, 0x33, 0x53, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x82, 0x40, 0x4c, 0x21, 0x3b, 0x2e, 0xf6, 0xdc, 0xc4, 0x82, 0x82,
	0xcc, 0xbc, 0x74, 0x09, 0x66, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x15, 0x3d, 0x84, 0xb1, 0x7a, 0xe8,
	0x46, 0xea, 0xf9, 0x42, 0xd4, 0x06, 0xc1, 0x34, 0x49, 0x45, 0x72, 0xb1, 0x43, 0xc5, 0x84, 0x14,
	0xb9, 0x78, 0x40, 0x0e, 0x4c, 0xcc, 0xcc, 0x4b, 0x2d, 0x8a, 0x87, 0xda, 0xcb, 0x1b, 0xc4, 0x0d,
	0x17, 0xf3, 0x4c, 0x11, 0x12, 0xe7, 0x62, 0xcf, 0xc8, 0x2f, 0x2e, 0x89, 0x87, 0xba, 0x81, 0x37,
	0x88, 0x0d, 0xc4, 0xf5, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x29, 0xce, 0xac, 0x4a, 0x95, 0x60, 0x06,
	0x8b, 0x82, 0xd9, 0x4a, 0xc2, 0x5c, 0x82, 0x48, 0xf6, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a,
	0xc5, 0x70, 0xf1, 0x39, 0x43, 0x42, 0x21, 0x18, 0xe4, 0xc8, 0xe4, 0x54, 0x21, 0x2f, 0x2e, 0x4e,
	0xb8, 0x32, 0x21, 0x19, 0x7c, 0xae, 0x97, 0x92, 0xc5, 0x21, 0x0b, 0x31, 0x5b, 0x89, 0xc1, 0x89,
	0x35, 0x8a, 0x39, 0xb1, 0x20, 0x33, 0x89, 0x0d, 0x1c, 0xc0, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x93, 0x20, 0x59, 0xd8, 0x71, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ControlServiceClient is the client API for ControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlServiceClient interface {
	// Newuidmap sets sthe UID/GID map of a user namespace
	Newuidmap(ctx context.Context, in *NewuidmapRequest, opts ...grpc.CallOption) (*NewuidmapResponse, error)
}

type controlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControlServiceClient(cc grpc.ClientConnInterface) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) Newuidmap(ctx context.Context, in *NewuidmapRequest, opts ...grpc.CallOption) (*NewuidmapResponse, error) {
	out := new(NewuidmapResponse)
	err := c.cc.Invoke(ctx, "/supervisor.ControlService/Newuidmap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServiceServer is the server API for ControlService service.
type ControlServiceServer interface {
	// Newuidmap sets sthe UID/GID map of a user namespace
	Newuidmap(context.Context, *NewuidmapRequest) (*NewuidmapResponse, error)
}

// UnimplementedControlServiceServer can be embedded to have forward compatible implementations.
type UnimplementedControlServiceServer struct {
}

func (*UnimplementedControlServiceServer) Newuidmap(ctx context.Context, req *NewuidmapRequest) (*NewuidmapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Newuidmap not implemented")
}

func RegisterControlServiceServer(s *grpc.Server, srv ControlServiceServer) {
	s.RegisterService(&_ControlService_serviceDesc, srv)
}

func _ControlService_Newuidmap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewuidmapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Newuidmap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.ControlService/Newuidmap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Newuidmap(ctx, req.(*NewuidmapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Newuidmap",
			Handler:    _ControlService_Newuidmap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}

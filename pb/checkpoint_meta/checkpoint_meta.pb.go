// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/checkpoint_meta/checkpoint_meta.proto

package checkpoint_meta

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type GetLatestVersionRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLatestVersionRequest) Reset()         { *m = GetLatestVersionRequest{} }
func (m *GetLatestVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestVersionRequest) ProtoMessage()    {}
func (*GetLatestVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be6effb39c1cf1e, []int{0}
}
func (m *GetLatestVersionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestVersionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestVersionRequest.Merge(m, src)
}
func (m *GetLatestVersionRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestVersionRequest proto.InternalMessageInfo

func (m *GetLatestVersionRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type GetLatestVersionResponse struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLatestVersionResponse) Reset()         { *m = GetLatestVersionResponse{} }
func (m *GetLatestVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestVersionResponse) ProtoMessage()    {}
func (*GetLatestVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be6effb39c1cf1e, []int{1}
}
func (m *GetLatestVersionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestVersionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestVersionResponse.Merge(m, src)
}
func (m *GetLatestVersionResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestVersionResponse proto.InternalMessageInfo

func (m *GetLatestVersionResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type UpdatePartitionRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
	PartitionNum         string   `protobuf:"bytes,2,opt,name=partition_num,json=partitionNum,proto3" json:"partition_num,omitempty"`
	Leader               string   `protobuf:"bytes,3,opt,name=leader,proto3" json:"leader,omitempty"`
	Followers            []string `protobuf:"bytes,4,rep,name=followers,proto3" json:"followers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePartitionRequest) Reset()         { *m = UpdatePartitionRequest{} }
func (m *UpdatePartitionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePartitionRequest) ProtoMessage()    {}
func (*UpdatePartitionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be6effb39c1cf1e, []int{2}
}
func (m *UpdatePartitionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatePartitionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatePartitionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatePartitionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePartitionRequest.Merge(m, src)
}
func (m *UpdatePartitionRequest) XXX_Size() int {
	return m.Size()
}
func (m *UpdatePartitionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePartitionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePartitionRequest proto.InternalMessageInfo

func (m *UpdatePartitionRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *UpdatePartitionRequest) GetPartitionNum() string {
	if m != nil {
		return m.PartitionNum
	}
	return ""
}

func (m *UpdatePartitionRequest) GetLeader() string {
	if m != nil {
		return m.Leader
	}
	return ""
}

func (m *UpdatePartitionRequest) GetFollowers() []string {
	if m != nil {
		return m.Followers
	}
	return nil
}

type UpdatePartitionResponse struct {
	Results              string   `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePartitionResponse) Reset()         { *m = UpdatePartitionResponse{} }
func (m *UpdatePartitionResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePartitionResponse) ProtoMessage()    {}
func (*UpdatePartitionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6be6effb39c1cf1e, []int{3}
}
func (m *UpdatePartitionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatePartitionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatePartitionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatePartitionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePartitionResponse.Merge(m, src)
}
func (m *UpdatePartitionResponse) XXX_Size() int {
	return m.Size()
}
func (m *UpdatePartitionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePartitionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePartitionResponse proto.InternalMessageInfo

func (m *UpdatePartitionResponse) GetResults() string {
	if m != nil {
		return m.Results
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLatestVersionRequest)(nil), "checkpoint_meta.GetLatestVersionRequest")
	proto.RegisterType((*GetLatestVersionResponse)(nil), "checkpoint_meta.GetLatestVersionResponse")
	proto.RegisterType((*UpdatePartitionRequest)(nil), "checkpoint_meta.UpdatePartitionRequest")
	proto.RegisterType((*UpdatePartitionResponse)(nil), "checkpoint_meta.UpdatePartitionResponse")
}

func init() {
	proto.RegisterFile("pb/checkpoint_meta/checkpoint_meta.proto", fileDescriptor_6be6effb39c1cf1e)
}

var fileDescriptor_6be6effb39c1cf1e = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x48, 0xd2, 0x4f,
	0xce, 0x48, 0x4d, 0xce, 0x2e, 0xc8, 0xcf, 0xcc, 0x2b, 0x89, 0xcf, 0x4d, 0x2d, 0x49, 0x44, 0xe7,
	0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa3, 0x09, 0x2b, 0xe9, 0x73, 0x89, 0xbb, 0xa7,
	0x96, 0xf8, 0x24, 0x96, 0xa4, 0x16, 0x97, 0x84, 0xa5, 0x16, 0x15, 0x67, 0xe6, 0xe7, 0x05, 0xa5,
	0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x70, 0xb1, 0x66, 0xe5, 0x27, 0x79, 0xa6, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0x26, 0x5c, 0x12, 0x98, 0x1a, 0x8a, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0xcb, 0x20, 0x42, 0x60, 0x3d, 0xac, 0x41, 0x30, 0xae,
	0x52, 0x27, 0x23, 0x97, 0x58, 0x68, 0x41, 0x4a, 0x62, 0x49, 0x6a, 0x40, 0x62, 0x51, 0x49, 0x66,
	0x09, 0x21, 0x6b, 0x84, 0x94, 0xb9, 0x78, 0x0b, 0x60, 0x2a, 0xe3, 0xf3, 0x4a, 0x73, 0x25, 0x98,
	0xc0, 0xb2, 0x3c, 0x70, 0x41, 0xbf, 0xd2, 0x5c, 0x21, 0x31, 0x2e, 0xb6, 0x9c, 0xd4, 0xc4, 0x94,
	0xd4, 0x22, 0x09, 0x66, 0xb0, 0x2c, 0x94, 0x27, 0x24, 0xc3, 0xc5, 0x99, 0x96, 0x9f, 0x93, 0x93,
	0x5f, 0x9e, 0x5a, 0x54, 0x2c, 0xc1, 0xa2, 0xc0, 0xac, 0xc1, 0x19, 0x84, 0x10, 0x50, 0x32, 0xe6,
	0x12, 0xc7, 0x70, 0x0a, 0xc2, 0x03, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0xc5, 0x50, 0xd7, 0xc0,
	0xb8, 0x46, 0xf7, 0x19, 0xb9, 0xb8, 0x7d, 0x53, 0x4b, 0x12, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93,
	0x53, 0x85, 0x32, 0xb8, 0x04, 0xd0, 0x83, 0x41, 0x48, 0x43, 0x0f, 0x3d, 0xd0, 0x71, 0x04, 0xad,
	0x14, 0xd1, 0x2a, 0x95, 0x18, 0x84, 0xd2, 0xb8, 0xf8, 0xd1, 0x9c, 0x2b, 0xa4, 0x8e, 0xa1, 0x1d,
	0x7b, 0xd8, 0x62, 0xb1, 0x07, 0x87, 0xcf, 0x95, 0x18, 0x9c, 0x94, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xa2, 0x04, 0xf5, 0xac,
	0xd1, 0xb4, 0x27, 0xb1, 0x81, 0x93, 0x91, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x92, 0x78,
	0xf1, 0x72, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetaServiceClient is the client API for MetaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetaServiceClient interface {
	GetLatestVersion(ctx context.Context, in *GetLatestVersionRequest, opts ...grpc.CallOption) (*GetLatestVersionRequest, error)
	UpdatePartition(ctx context.Context, in *UpdatePartitionRequest, opts ...grpc.CallOption) (*UpdatePartitionResponse, error)
}

type metaServiceClient struct {
	cc *grpc.ClientConn
}

func NewMetaServiceClient(cc *grpc.ClientConn) MetaServiceClient {
	return &metaServiceClient{cc}
}

func (c *metaServiceClient) GetLatestVersion(ctx context.Context, in *GetLatestVersionRequest, opts ...grpc.CallOption) (*GetLatestVersionRequest, error) {
	out := new(GetLatestVersionRequest)
	err := c.cc.Invoke(ctx, "/checkpoint_meta.MetaService/GetLatestVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaServiceClient) UpdatePartition(ctx context.Context, in *UpdatePartitionRequest, opts ...grpc.CallOption) (*UpdatePartitionResponse, error) {
	out := new(UpdatePartitionResponse)
	err := c.cc.Invoke(ctx, "/checkpoint_meta.MetaService/UpdatePartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetaServiceServer is the server API for MetaService service.
type MetaServiceServer interface {
	GetLatestVersion(context.Context, *GetLatestVersionRequest) (*GetLatestVersionRequest, error)
	UpdatePartition(context.Context, *UpdatePartitionRequest) (*UpdatePartitionResponse, error)
}

// UnimplementedMetaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetaServiceServer struct {
}

func (*UnimplementedMetaServiceServer) GetLatestVersion(ctx context.Context, req *GetLatestVersionRequest) (*GetLatestVersionRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestVersion not implemented")
}
func (*UnimplementedMetaServiceServer) UpdatePartition(ctx context.Context, req *UpdatePartitionRequest) (*UpdatePartitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartition not implemented")
}

func RegisterMetaServiceServer(s *grpc.Server, srv MetaServiceServer) {
	s.RegisterService(&_MetaService_serviceDesc, srv)
}

func _MetaService_GetLatestVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).GetLatestVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkpoint_meta.MetaService/GetLatestVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).GetLatestVersion(ctx, req.(*GetLatestVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaService_UpdatePartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaServiceServer).UpdatePartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/checkpoint_meta.MetaService/UpdatePartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaServiceServer).UpdatePartition(ctx, req.(*UpdatePartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "checkpoint_meta.MetaService",
	HandlerType: (*MetaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestVersion",
			Handler:    _MetaService_GetLatestVersion_Handler,
		},
		{
			MethodName: "UpdatePartition",
			Handler:    _MetaService_UpdatePartition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/checkpoint_meta/checkpoint_meta.proto",
}

func (m *GetLatestVersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestVersionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestVersionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.JobId) > 0 {
		i -= len(m.JobId)
		copy(dAtA[i:], m.JobId)
		i = encodeVarintCheckpointMeta(dAtA, i, uint64(len(m.JobId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetLatestVersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestVersionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestVersionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Version != 0 {
		i = encodeVarintCheckpointMeta(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UpdatePartitionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatePartitionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdatePartitionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Followers) > 0 {
		for iNdEx := len(m.Followers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Followers[iNdEx])
			copy(dAtA[i:], m.Followers[iNdEx])
			i = encodeVarintCheckpointMeta(dAtA, i, uint64(len(m.Followers[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Leader) > 0 {
		i -= len(m.Leader)
		copy(dAtA[i:], m.Leader)
		i = encodeVarintCheckpointMeta(dAtA, i, uint64(len(m.Leader)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PartitionNum) > 0 {
		i -= len(m.PartitionNum)
		copy(dAtA[i:], m.PartitionNum)
		i = encodeVarintCheckpointMeta(dAtA, i, uint64(len(m.PartitionNum)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.JobId) > 0 {
		i -= len(m.JobId)
		copy(dAtA[i:], m.JobId)
		i = encodeVarintCheckpointMeta(dAtA, i, uint64(len(m.JobId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdatePartitionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatePartitionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdatePartitionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Results) > 0 {
		i -= len(m.Results)
		copy(dAtA[i:], m.Results)
		i = encodeVarintCheckpointMeta(dAtA, i, uint64(len(m.Results)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCheckpointMeta(dAtA []byte, offset int, v uint64) int {
	offset -= sovCheckpointMeta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetLatestVersionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JobId)
	if l > 0 {
		n += 1 + l + sovCheckpointMeta(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetLatestVersionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovCheckpointMeta(uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdatePartitionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JobId)
	if l > 0 {
		n += 1 + l + sovCheckpointMeta(uint64(l))
	}
	l = len(m.PartitionNum)
	if l > 0 {
		n += 1 + l + sovCheckpointMeta(uint64(l))
	}
	l = len(m.Leader)
	if l > 0 {
		n += 1 + l + sovCheckpointMeta(uint64(l))
	}
	if len(m.Followers) > 0 {
		for _, s := range m.Followers {
			l = len(s)
			n += 1 + l + sovCheckpointMeta(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdatePartitionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Results)
	if l > 0 {
		n += 1 + l + sovCheckpointMeta(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCheckpointMeta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCheckpointMeta(x uint64) (n int) {
	return sovCheckpointMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetLatestVersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpointMeta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetLatestVersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestVersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpointMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetLatestVersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpointMeta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetLatestVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpointMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdatePartitionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpointMeta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdatePartitionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatePartitionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartitionNum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartitionNum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Followers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Followers = append(m.Followers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpointMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdatePartitionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpointMeta
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UpdatePartitionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatePartitionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpointMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCheckpointMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCheckpointMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCheckpointMeta
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCheckpointMeta
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCheckpointMeta
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCheckpointMeta
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCheckpointMeta
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCheckpointMeta        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCheckpointMeta          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCheckpointMeta = fmt.Errorf("proto: unexpected end of group")
)
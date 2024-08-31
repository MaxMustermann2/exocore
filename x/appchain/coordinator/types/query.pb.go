// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/appchain/coordinator/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/ExocoreNetwork/exocore/x/appchain/common/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query.QueryParams RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c71eb5ef95876a4b, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query.QueryParams RPC method.
type QueryParamsResponse struct {
	// params is the parameters for the appchain coordinator module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c71eb5ef95876a4b, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QuerySubscriberGenesisRequest is the request type for the Query.QuerySubscriberGenesis RPC method.
type QuerySubscriberGenesisRequest struct {
	// chain is the chain ID of the subscriber chain. we intentionally don't use ChainID so that
	// the query can work (it does not support custom names).
	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *QuerySubscriberGenesisRequest) Reset()         { *m = QuerySubscriberGenesisRequest{} }
func (m *QuerySubscriberGenesisRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySubscriberGenesisRequest) ProtoMessage()    {}
func (*QuerySubscriberGenesisRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c71eb5ef95876a4b, []int{2}
}
func (m *QuerySubscriberGenesisRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubscriberGenesisRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubscriberGenesisRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubscriberGenesisRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubscriberGenesisRequest.Merge(m, src)
}
func (m *QuerySubscriberGenesisRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubscriberGenesisRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubscriberGenesisRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubscriberGenesisRequest proto.InternalMessageInfo

func (m *QuerySubscriberGenesisRequest) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

// QuerySubscriberGenesisResponse is the response type for the Query.QuerySubscriberGenesis RPC method.
type QuerySubscriberGenesisResponse struct {
	// subscriber_genesis is the genesis state for the subscriber chain.
	SubscriberGenesis types.SubscriberGenesisState `protobuf:"bytes,1,opt,name=subscriber_genesis,json=subscriberGenesis,proto3" json:"subscriber_genesis"`
}

func (m *QuerySubscriberGenesisResponse) Reset()         { *m = QuerySubscriberGenesisResponse{} }
func (m *QuerySubscriberGenesisResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySubscriberGenesisResponse) ProtoMessage()    {}
func (*QuerySubscriberGenesisResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c71eb5ef95876a4b, []int{3}
}
func (m *QuerySubscriberGenesisResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubscriberGenesisResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubscriberGenesisResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubscriberGenesisResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubscriberGenesisResponse.Merge(m, src)
}
func (m *QuerySubscriberGenesisResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubscriberGenesisResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubscriberGenesisResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubscriberGenesisResponse proto.InternalMessageInfo

func (m *QuerySubscriberGenesisResponse) GetSubscriberGenesis() types.SubscriberGenesisState {
	if m != nil {
		return m.SubscriberGenesis
	}
	return types.SubscriberGenesisState{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "exocore.appchain.coordinator.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "exocore.appchain.coordinator.v1.QueryParamsResponse")
	proto.RegisterType((*QuerySubscriberGenesisRequest)(nil), "exocore.appchain.coordinator.v1.QuerySubscriberGenesisRequest")
	proto.RegisterType((*QuerySubscriberGenesisResponse)(nil), "exocore.appchain.coordinator.v1.QuerySubscriberGenesisResponse")
}

func init() {
	proto.RegisterFile("exocore/appchain/coordinator/v1/query.proto", fileDescriptor_c71eb5ef95876a4b)
}

var fileDescriptor_c71eb5ef95876a4b = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x8b, 0x13, 0x31,
	0x14, 0x9f, 0xa8, 0xbb, 0x60, 0xf6, 0x64, 0x2c, 0x22, 0xc3, 0x9a, 0x95, 0x41, 0x58, 0x41, 0x99,
	0xd0, 0xae, 0x9e, 0xc4, 0x15, 0x56, 0x16, 0x6f, 0xa2, 0xdd, 0x8b, 0x88, 0x20, 0x99, 0x31, 0xcc,
	0x0e, 0xda, 0xbc, 0x69, 0x92, 0xa9, 0x2d, 0xe2, 0xc5, 0x5b, 0x6f, 0x82, 0x77, 0x3f, 0x4f, 0x8f,
	0x05, 0x2f, 0x9e, 0x8a, 0xb4, 0x7e, 0x10, 0x99, 0x24, 0x2d, 0xd5, 0x4e, 0x19, 0xf5, 0x96, 0xc9,
	0x7b, 0xbf, 0x3f, 0xef, 0xf7, 0x32, 0xf8, 0x8e, 0x18, 0x42, 0x0a, 0x4a, 0x30, 0x5e, 0x14, 0xe9,
	0x39, 0xcf, 0x25, 0x4b, 0x01, 0xd4, 0x9b, 0x5c, 0x72, 0x03, 0x8a, 0x0d, 0xda, 0xac, 0x5f, 0x0a,
	0x35, 0x8a, 0x0b, 0x05, 0x06, 0xc8, 0x81, 0x6f, 0x8e, 0x97, 0xcd, 0xf1, 0x5a, 0x73, 0x3c, 0x68,
	0x87, 0x87, 0x35, 0x6c, 0xbd, 0x1e, 0xc8, 0x8a, 0xc8, 0x9d, 0x1c, 0x53, 0x78, 0xb7, 0x49, 0xb6,
	0xe0, 0x8a, 0xf7, 0xb4, 0xef, 0x6e, 0x65, 0x90, 0x81, 0x3d, 0xb2, 0xea, 0xe4, 0x6f, 0xf7, 0x33,
	0x80, 0xec, 0x5d, 0x45, 0x91, 0x33, 0x2e, 0x25, 0x18, 0x6e, 0x72, 0x90, 0x1e, 0x13, 0xb5, 0x30,
	0x79, 0x5e, 0x59, 0x7f, 0x66, 0x89, 0xba, 0xa2, 0x5f, 0x0a, 0x6d, 0xa2, 0x57, 0xf8, 0xea, 0x6f,
	0xb7, 0xba, 0x00, 0xa9, 0x05, 0x39, 0xc5, 0xbb, 0x4e, 0xf0, 0x3a, 0xba, 0x89, 0x6e, 0xef, 0x75,
	0x0e, 0xe3, 0x86, 0x49, 0x63, 0x47, 0x70, 0x72, 0x69, 0x32, 0x3b, 0x08, 0xba, 0x1e, 0x1c, 0xdd,
	0xc7, 0x37, 0x2c, 0xfb, 0x59, 0x99, 0xe8, 0x54, 0xe5, 0x89, 0x50, 0x4f, 0x84, 0x14, 0x3a, 0x5f,
	0xca, 0x93, 0x16, 0xde, 0xb1, 0x6c, 0x56, 0xe6, 0x72, 0xd7, 0x7d, 0x44, 0x63, 0x84, 0xe9, 0x36,
	0x9c, 0x37, 0x98, 0x61, 0xa2, 0x57, 0xc5, 0xd7, 0x99, 0xab, 0x7a, 0xb3, 0x9d, 0x3a, 0xb3, 0x36,
	0xeb, 0x41, 0x3b, 0xde, 0xa0, 0x3c, 0x33, 0xdc, 0x08, 0xef, 0xfb, 0x8a, 0xfe, 0xb3, 0xda, 0x19,
	0x5f, 0xc4, 0x3b, 0xd6, 0x0b, 0xf9, 0x8a, 0xf0, 0xde, 0x5a, 0x56, 0xe4, 0xa8, 0x31, 0x93, 0xcd,
	0xbc, 0xc3, 0x7b, 0xff, 0x06, 0x72, 0xd3, 0x46, 0xb7, 0x3e, 0x7d, 0xfb, 0xf9, 0xe5, 0x02, 0x25,
	0xfb, 0xf5, 0xcf, 0xc3, 0xa5, 0x4d, 0x66, 0x08, 0x5f, 0xab, 0x8f, 0x8d, 0x1c, 0xff, 0x9d, 0xec,
	0xb6, 0x3d, 0x85, 0x8f, 0xfe, 0x1b, 0xef, 0x27, 0x78, 0x6c, 0x27, 0x78, 0x48, 0x1e, 0xb0, 0xa6,
	0x87, 0xbe, 0xb9, 0x56, 0xf6, 0xc1, 0x36, 0x7e, 0x3c, 0x79, 0x31, 0x99, 0x53, 0x34, 0x9d, 0x53,
	0xf4, 0x63, 0x4e, 0xd1, 0xe7, 0x05, 0x0d, 0xa6, 0x0b, 0x1a, 0x7c, 0x5f, 0xd0, 0xe0, 0xe5, 0x71,
	0x96, 0x9b, 0xf3, 0x32, 0xa9, 0xf6, 0xcc, 0x4e, 0x9d, 0xc0, 0x53, 0x61, 0xde, 0x83, 0x7a, 0xbb,
	0xd2, 0x1b, 0xd6, 0x2b, 0x9a, 0x51, 0x21, 0x74, 0xb2, 0x6b, 0xff, 0x91, 0xa3, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0x5d, 0x4a, 0xc3, 0xfe, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// QueryParams returns the appchain coordinator module parameters.
	QueryParams(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// QuerySubscriberGenesis returns the genesis state for a subscriber chain.
	QuerySubscriberGenesis(ctx context.Context, in *QuerySubscriberGenesisRequest, opts ...grpc.CallOption) (*QuerySubscriberGenesisResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) QueryParams(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/exocore.appchain.coordinator.v1.Query/QueryParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QuerySubscriberGenesis(ctx context.Context, in *QuerySubscriberGenesisRequest, opts ...grpc.CallOption) (*QuerySubscriberGenesisResponse, error) {
	out := new(QuerySubscriberGenesisResponse)
	err := c.cc.Invoke(ctx, "/exocore.appchain.coordinator.v1.Query/QuerySubscriberGenesis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// QueryParams returns the appchain coordinator module parameters.
	QueryParams(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// QuerySubscriberGenesis returns the genesis state for a subscriber chain.
	QuerySubscriberGenesis(context.Context, *QuerySubscriberGenesisRequest) (*QuerySubscriberGenesisResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) QueryParams(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryParams not implemented")
}
func (*UnimplementedQueryServer) QuerySubscriberGenesis(ctx context.Context, req *QuerySubscriberGenesisRequest) (*QuerySubscriberGenesisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySubscriberGenesis not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_QueryParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exocore.appchain.coordinator.v1.Query/QueryParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryParams(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QuerySubscriberGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubscriberGenesisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QuerySubscriberGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exocore.appchain.coordinator.v1.Query/QuerySubscriberGenesis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QuerySubscriberGenesis(ctx, req.(*QuerySubscriberGenesisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exocore.appchain.coordinator.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryParams",
			Handler:    _Query_QueryParams_Handler,
		},
		{
			MethodName: "QuerySubscriberGenesis",
			Handler:    _Query_QuerySubscriberGenesis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exocore/appchain/coordinator/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySubscriberGenesisRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubscriberGenesisRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubscriberGenesisRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySubscriberGenesisResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubscriberGenesisResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubscriberGenesisResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SubscriberGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySubscriberGenesisRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySubscriberGenesisResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SubscriberGenesis.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySubscriberGenesisRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QuerySubscriberGenesisRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubscriberGenesisRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySubscriberGenesisResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QuerySubscriberGenesisResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubscriberGenesisResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriberGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubscriberGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

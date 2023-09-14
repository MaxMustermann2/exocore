// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/restaking_assets_manage/v1/tx.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type ClientChainInfo struct {
	ChainName         string `protobuf:"bytes,1,opt,name=ChainName,proto3" json:"ChainName,omitempty"`
	ChainMetaInfo     string `protobuf:"bytes,2,opt,name=ChainMetaInfo,proto3" json:"ChainMetaInfo,omitempty"`
	ChainId           uint64 `protobuf:"varint,3,opt,name=ChainId,proto3" json:"ChainId,omitempty"`
	ExoCoreChainIndex uint64 `protobuf:"varint,4,opt,name=ExoCoreChainIndex,proto3" json:"ExoCoreChainIndex,omitempty"`
}

func (m *ClientChainInfo) Reset()         { *m = ClientChainInfo{} }
func (m *ClientChainInfo) String() string { return proto.CompactTextString(m) }
func (*ClientChainInfo) ProtoMessage()    {}
func (*ClientChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b24e66e530cc30d1, []int{0}
}
func (m *ClientChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientChainInfo.Merge(m, src)
}
func (m *ClientChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *ClientChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientChainInfo proto.InternalMessageInfo

func (m *ClientChainInfo) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *ClientChainInfo) GetChainMetaInfo() string {
	if m != nil {
		return m.ChainMetaInfo
	}
	return ""
}

func (m *ClientChainInfo) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ClientChainInfo) GetExoCoreChainIndex() uint64 {
	if m != nil {
		return m.ExoCoreChainIndex
	}
	return 0
}

type ReStakingAssetInfo struct {
	Name                 string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Symbol               string `protobuf:"bytes,2,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Address              string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Decimals             uint32 `protobuf:"varint,4,opt,name=Decimals,proto3" json:"Decimals,omitempty"`
	TotalSupply          uint64 `protobuf:"varint,5,opt,name=TotalSupply,proto3" json:"TotalSupply,omitempty"`
	ReStakingTotalAmount uint64 `protobuf:"varint,6,opt,name=ReStakingTotalAmount,proto3" json:"ReStakingTotalAmount,omitempty"`
	ExoCoreAssetIndex    uint64 `protobuf:"varint,7,opt,name=ExoCoreAssetIndex,proto3" json:"ExoCoreAssetIndex,omitempty"`
	ExoCoreChainIndex    uint64 `protobuf:"varint,8,opt,name=ExoCoreChainIndex,proto3" json:"ExoCoreChainIndex,omitempty"`
	AssetMetaInfo        string `protobuf:"bytes,9,opt,name=AssetMetaInfo,proto3" json:"AssetMetaInfo,omitempty"`
}

func (m *ReStakingAssetInfo) Reset()         { *m = ReStakingAssetInfo{} }
func (m *ReStakingAssetInfo) String() string { return proto.CompactTextString(m) }
func (*ReStakingAssetInfo) ProtoMessage()    {}
func (*ReStakingAssetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b24e66e530cc30d1, []int{1}
}
func (m *ReStakingAssetInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReStakingAssetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReStakingAssetInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReStakingAssetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReStakingAssetInfo.Merge(m, src)
}
func (m *ReStakingAssetInfo) XXX_Size() int {
	return m.Size()
}
func (m *ReStakingAssetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReStakingAssetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReStakingAssetInfo proto.InternalMessageInfo

func (m *ReStakingAssetInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReStakingAssetInfo) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReStakingAssetInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ReStakingAssetInfo) GetDecimals() uint32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *ReStakingAssetInfo) GetTotalSupply() uint64 {
	if m != nil {
		return m.TotalSupply
	}
	return 0
}

func (m *ReStakingAssetInfo) GetReStakingTotalAmount() uint64 {
	if m != nil {
		return m.ReStakingTotalAmount
	}
	return 0
}

func (m *ReStakingAssetInfo) GetExoCoreAssetIndex() uint64 {
	if m != nil {
		return m.ExoCoreAssetIndex
	}
	return 0
}

func (m *ReStakingAssetInfo) GetExoCoreChainIndex() uint64 {
	if m != nil {
		return m.ExoCoreChainIndex
	}
	return 0
}

func (m *ReStakingAssetInfo) GetAssetMetaInfo() string {
	if m != nil {
		return m.AssetMetaInfo
	}
	return ""
}

type ReStakerSingleAssetInfo struct {
	TotalDepositAmount uint64 `protobuf:"varint,1,opt,name=TotalDepositAmount,proto3" json:"TotalDepositAmount,omitempty"`
	CanWithDrawAmount  uint64 `protobuf:"varint,2,opt,name=CanWithDrawAmount,proto3" json:"CanWithDrawAmount,omitempty"`
}

func (m *ReStakerSingleAssetInfo) Reset()         { *m = ReStakerSingleAssetInfo{} }
func (m *ReStakerSingleAssetInfo) String() string { return proto.CompactTextString(m) }
func (*ReStakerSingleAssetInfo) ProtoMessage()    {}
func (*ReStakerSingleAssetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b24e66e530cc30d1, []int{2}
}
func (m *ReStakerSingleAssetInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReStakerSingleAssetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReStakerSingleAssetInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReStakerSingleAssetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReStakerSingleAssetInfo.Merge(m, src)
}
func (m *ReStakerSingleAssetInfo) XXX_Size() int {
	return m.Size()
}
func (m *ReStakerSingleAssetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReStakerSingleAssetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReStakerSingleAssetInfo proto.InternalMessageInfo

func (m *ReStakerSingleAssetInfo) GetTotalDepositAmount() uint64 {
	if m != nil {
		return m.TotalDepositAmount
	}
	return 0
}

func (m *ReStakerSingleAssetInfo) GetCanWithDrawAmount() uint64 {
	if m != nil {
		return m.CanWithDrawAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientChainInfo)(nil), "exocore.restaking_assets_manage.v1.ClientChainInfo")
	proto.RegisterType((*ReStakingAssetInfo)(nil), "exocore.restaking_assets_manage.v1.ReStakingAssetInfo")
	proto.RegisterType((*ReStakerSingleAssetInfo)(nil), "exocore.restaking_assets_manage.v1.ReStakerSingleAssetInfo")
}

func init() {
	proto.RegisterFile("exocore/restaking_assets_manage/v1/tx.proto", fileDescriptor_b24e66e530cc30d1)
}

var fileDescriptor_b24e66e530cc30d1 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0xaa, 0xda, 0x40,
	0x14, 0xc6, 0x1d, 0x6b, 0xd5, 0x4c, 0x91, 0xd2, 0xa1, 0xb4, 0xa1, 0x94, 0x20, 0xa1, 0x0b, 0xa1,
	0x25, 0x41, 0xfb, 0x04, 0x36, 0x76, 0x51, 0x68, 0xbb, 0x48, 0x0a, 0x85, 0x6e, 0x64, 0x34, 0xa7,
	0x71, 0xb8, 0xc9, 0x4c, 0xc8, 0x8c, 0x1a, 0xdf, 0xe2, 0xbe, 0xc0, 0x7d, 0x9f, 0xbb, 0xba, 0xb8,
	0xbc, 0xcb, 0x8b, 0xbe, 0xc8, 0x25, 0x93, 0xf8, 0x0f, 0x75, 0x97, 0xef, 0x9c, 0x2f, 0x33, 0xbf,
	0x6f, 0xce, 0xc1, 0x9f, 0x21, 0x17, 0x53, 0x91, 0x81, 0x9b, 0x81, 0x54, 0xf4, 0x86, 0xf1, 0x68,
	0x4c, 0xa5, 0x04, 0x25, 0xc7, 0x09, 0xe5, 0x34, 0x02, 0x77, 0xd1, 0x77, 0x55, 0xee, 0xa4, 0x99,
	0x50, 0x82, 0xd8, 0x95, 0xd9, 0xb9, 0x62, 0x76, 0x16, 0x7d, 0xfb, 0x0e, 0xe1, 0xd7, 0x5e, 0xcc,
	0x80, 0x2b, 0x6f, 0x46, 0x19, 0xff, 0xc1, 0xff, 0x0b, 0xf2, 0x11, 0x1b, 0x5a, 0xfc, 0xa6, 0x09,
	0x98, 0xa8, 0x8b, 0x7a, 0x86, 0x7f, 0x28, 0x90, 0x4f, 0xb8, 0xa3, 0xc5, 0x2f, 0x50, 0xb4, 0xb0,
	0x9b, 0x75, 0xed, 0x38, 0x2d, 0x12, 0x13, 0xb7, 0xca, 0x03, 0x43, 0xf3, 0x45, 0x17, 0xf5, 0x1a,
	0xfe, 0x4e, 0x92, 0x2f, 0xf8, 0xcd, 0xf7, 0x5c, 0x78, 0x22, 0x83, 0xea, 0xc6, 0x10, 0x72, 0xb3,
	0xa1, 0x3d, 0xe7, 0x0d, 0xfb, 0xa1, 0x8e, 0x89, 0x0f, 0x41, 0x89, 0x3f, 0x2c, 0xe8, 0xf5, 0xf1,
	0x04, 0x37, 0x8e, 0xe8, 0xf4, 0x37, 0x79, 0x87, 0x9b, 0xc1, 0x2a, 0x99, 0x88, 0xb8, 0x22, 0xaa,
	0x54, 0x81, 0x32, 0x0c, 0xc3, 0x0c, 0xa4, 0xd4, 0x28, 0x86, 0xbf, 0x93, 0xe4, 0x03, 0x6e, 0x8f,
	0x60, 0xca, 0x12, 0x1a, 0x4b, 0x4d, 0xd0, 0xf1, 0xf7, 0x9a, 0x74, 0xf1, 0xab, 0x3f, 0x42, 0xd1,
	0x38, 0x98, 0xa7, 0x69, 0xbc, 0x32, 0x5f, 0x6a, 0xc0, 0xe3, 0x12, 0x19, 0xe0, 0xb7, 0x7b, 0x32,
	0x5d, 0x1f, 0x26, 0x62, 0xce, 0x95, 0xd9, 0xd4, 0xd6, 0x8b, 0xbd, 0xa3, 0xf0, 0x55, 0x96, 0x22,
	0x7c, 0xeb, 0x24, 0xfc, 0xa1, 0x71, 0xf9, 0xa9, 0xda, 0x57, 0x9e, 0xaa, 0x18, 0x8c, 0xfe, 0x77,
	0x3f, 0x18, 0xa3, 0x1c, 0xcc, 0x49, 0xd1, 0x5e, 0xe2, 0xf7, 0x25, 0x19, 0x64, 0x01, 0xe3, 0x51,
	0x0c, 0x87, 0x47, 0x75, 0x30, 0xd1, 0xac, 0x23, 0x48, 0x85, 0x64, 0xaa, 0x8a, 0x83, 0xf4, 0x7d,
	0x17, 0x3a, 0x05, 0x9e, 0x47, 0xf9, 0x5f, 0xa6, 0x66, 0xa3, 0x8c, 0x2e, 0x2b, 0x7b, 0xbd, 0xc4,
	0x3b, 0x6b, 0x7c, 0xfb, 0x79, 0xbf, 0xb1, 0xd0, 0x7a, 0x63, 0xa1, 0xa7, 0x8d, 0x85, 0x6e, 0xb7,
	0x56, 0x6d, 0xbd, 0xb5, 0x6a, 0x8f, 0x5b, 0xab, 0xf6, 0x6f, 0x10, 0x31, 0x35, 0x9b, 0x4f, 0x9c,
	0xa9, 0x48, 0xdc, 0xdd, 0x7e, 0xe7, 0x57, 0x37, 0x5c, 0xad, 0x52, 0x90, 0x93, 0xa6, 0x5e, 0xf1,
	0xaf, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0x91, 0xc9, 0xc3, 0x11, 0x03, 0x00, 0x00,
}

func (m *ClientChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExoCoreChainIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ExoCoreChainIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.ChainId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainMetaInfo) > 0 {
		i -= len(m.ChainMetaInfo)
		copy(dAtA[i:], m.ChainMetaInfo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainMetaInfo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReStakingAssetInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReStakingAssetInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReStakingAssetInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetMetaInfo) > 0 {
		i -= len(m.AssetMetaInfo)
		copy(dAtA[i:], m.AssetMetaInfo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AssetMetaInfo)))
		i--
		dAtA[i] = 0x4a
	}
	if m.ExoCoreChainIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ExoCoreChainIndex))
		i--
		dAtA[i] = 0x40
	}
	if m.ExoCoreAssetIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ExoCoreAssetIndex))
		i--
		dAtA[i] = 0x38
	}
	if m.ReStakingTotalAmount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ReStakingTotalAmount))
		i--
		dAtA[i] = 0x30
	}
	if m.TotalSupply != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TotalSupply))
		i--
		dAtA[i] = 0x28
	}
	if m.Decimals != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReStakerSingleAssetInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReStakerSingleAssetInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReStakerSingleAssetInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CanWithDrawAmount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CanWithDrawAmount))
		i--
		dAtA[i] = 0x10
	}
	if m.TotalDepositAmount != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TotalDepositAmount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainMetaInfo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ChainId != 0 {
		n += 1 + sovTx(uint64(m.ChainId))
	}
	if m.ExoCoreChainIndex != 0 {
		n += 1 + sovTx(uint64(m.ExoCoreChainIndex))
	}
	return n
}

func (m *ReStakingAssetInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovTx(uint64(m.Decimals))
	}
	if m.TotalSupply != 0 {
		n += 1 + sovTx(uint64(m.TotalSupply))
	}
	if m.ReStakingTotalAmount != 0 {
		n += 1 + sovTx(uint64(m.ReStakingTotalAmount))
	}
	if m.ExoCoreAssetIndex != 0 {
		n += 1 + sovTx(uint64(m.ExoCoreAssetIndex))
	}
	if m.ExoCoreChainIndex != 0 {
		n += 1 + sovTx(uint64(m.ExoCoreChainIndex))
	}
	l = len(m.AssetMetaInfo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *ReStakerSingleAssetInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TotalDepositAmount != 0 {
		n += 1 + sovTx(uint64(m.TotalDepositAmount))
	}
	if m.CanWithDrawAmount != 0 {
		n += 1 + sovTx(uint64(m.CanWithDrawAmount))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: ClientChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainMetaInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainMetaInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExoCoreChainIndex", wireType)
			}
			m.ExoCoreChainIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExoCoreChainIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *ReStakingAssetInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: ReStakingAssetInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReStakingAssetInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			m.TotalSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReStakingTotalAmount", wireType)
			}
			m.ReStakingTotalAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReStakingTotalAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExoCoreAssetIndex", wireType)
			}
			m.ExoCoreAssetIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExoCoreAssetIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExoCoreChainIndex", wireType)
			}
			m.ExoCoreChainIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExoCoreChainIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetMetaInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetMetaInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *ReStakerSingleAssetInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: ReStakerSingleAssetInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReStakerSingleAssetInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDepositAmount", wireType)
			}
			m.TotalDepositAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalDepositAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanWithDrawAmount", wireType)
			}
			m.CanWithDrawAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CanWithDrawAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

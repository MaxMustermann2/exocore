// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/assets/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the assets module's state. It needs to encompass
// all of the state that is required to start the chain from the genesis
// or in the event of a restart. At this point, it is only built with
// the former in mind.
// TODO: make this state exportable for the case of chain restarts.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// client_chains is the list of supported client chains,
	// that are supported at chain genesis (or restart).
	ClientChains []ClientChainInfo `protobuf:"bytes,2,rep,name=client_chains,json=clientChains,proto3" json:"client_chains"`
	// tokens is the list of supported client chain tokens and total staked amount
	// that are supported at chain genesis (or restart).
	Tokens []StakingAssetInfo `protobuf:"bytes,3,rep,name=tokens,proto3" json:"tokens"`
	// operator_asset_infos is the list of deposits, indexed by operator
	// address and then the asset id. The struct is the `OperatorAssetInfo`
	// which contains deposits, unbonding, operator own amount, etc.
	OperatorAssetInfos []OpAssetIDAndInfos `protobuf:"bytes,4,rep,name=operator_asset_infos,json=operatorAssetInfos,proto3" json:"operator_asset_infos"`
	// staker_asset_infos is the list of deposits, indexed by staker
	// address and then the asset id. The struct is the `StakerAssetInfo`
	// which contains deposits, withdrawable and unbonding amount.
	StakerAssetInfos []StAssetIDAndInfos `protobuf:"bytes,5,rep,name=staker_asset_infos,json=stakerAssetInfos,proto3" json:"staker_asset_infos"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_caf4f124d39d82ce, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClientChains() []ClientChainInfo {
	if m != nil {
		return m.ClientChains
	}
	return nil
}

func (m *GenesisState) GetTokens() []StakingAssetInfo {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *GenesisState) GetOperatorAssetInfos() []OpAssetIDAndInfos {
	if m != nil {
		return m.OperatorAssetInfos
	}
	return nil
}

func (m *GenesisState) GetStakerAssetInfos() []StAssetIDAndInfos {
	if m != nil {
		return m.StakerAssetInfos
	}
	return nil
}

// OpAssetIDAndInfo is a helper struct to be used in the genesis state.
// It is used to store the asset id and its info for an operator.
type OpAssetIDAndInfo struct {
	// asset_id is the id of the asset.
	AssetID string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// info is the asset info.
	Info *OperatorAssetInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *OpAssetIDAndInfo) Reset()         { *m = OpAssetIDAndInfo{} }
func (m *OpAssetIDAndInfo) String() string { return proto.CompactTextString(m) }
func (*OpAssetIDAndInfo) ProtoMessage()    {}
func (*OpAssetIDAndInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_caf4f124d39d82ce, []int{1}
}
func (m *OpAssetIDAndInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OpAssetIDAndInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OpAssetIDAndInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OpAssetIDAndInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpAssetIDAndInfo.Merge(m, src)
}
func (m *OpAssetIDAndInfo) XXX_Size() int {
	return m.Size()
}
func (m *OpAssetIDAndInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpAssetIDAndInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpAssetIDAndInfo proto.InternalMessageInfo

func (m *OpAssetIDAndInfo) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

func (m *OpAssetIDAndInfo) GetInfo() *OperatorAssetInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

// OpAssetIDAndInfos is a helper struct to be used in the genesis state.
// It is used to store the operator address and its asset id and info.
type OpAssetIDAndInfos struct {
	// operator_address is the address of the operator.
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
	// asset_id_and_infos is the list of asset id and infos that exist
	// for the operator.
	AssetIdAndInfos []*OpAssetIDAndInfo `protobuf:"bytes,2,rep,name=asset_id_and_infos,json=assetIdAndInfos,proto3" json:"asset_id_and_infos,omitempty"`
}

func (m *OpAssetIDAndInfos) Reset()         { *m = OpAssetIDAndInfos{} }
func (m *OpAssetIDAndInfos) String() string { return proto.CompactTextString(m) }
func (*OpAssetIDAndInfos) ProtoMessage()    {}
func (*OpAssetIDAndInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_caf4f124d39d82ce, []int{2}
}
func (m *OpAssetIDAndInfos) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OpAssetIDAndInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OpAssetIDAndInfos.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OpAssetIDAndInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpAssetIDAndInfos.Merge(m, src)
}
func (m *OpAssetIDAndInfos) XXX_Size() int {
	return m.Size()
}
func (m *OpAssetIDAndInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_OpAssetIDAndInfos.DiscardUnknown(m)
}

var xxx_messageInfo_OpAssetIDAndInfos proto.InternalMessageInfo

func (m *OpAssetIDAndInfos) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *OpAssetIDAndInfos) GetAssetIdAndInfos() []*OpAssetIDAndInfo {
	if m != nil {
		return m.AssetIdAndInfos
	}
	return nil
}

// StAssetIDAndInfo is a helper struct to be used in the genesis state.
// It is used to store the asset id and its info for an staker.
type StAssetIDAndInfo struct {
	// asset_id is the id of the asset.
	AssetID string `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// info is the asset info.
	Info *StakerAssetInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *StAssetIDAndInfo) Reset()         { *m = StAssetIDAndInfo{} }
func (m *StAssetIDAndInfo) String() string { return proto.CompactTextString(m) }
func (*StAssetIDAndInfo) ProtoMessage()    {}
func (*StAssetIDAndInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_caf4f124d39d82ce, []int{3}
}
func (m *StAssetIDAndInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StAssetIDAndInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StAssetIDAndInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StAssetIDAndInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StAssetIDAndInfo.Merge(m, src)
}
func (m *StAssetIDAndInfo) XXX_Size() int {
	return m.Size()
}
func (m *StAssetIDAndInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StAssetIDAndInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StAssetIDAndInfo proto.InternalMessageInfo

func (m *StAssetIDAndInfo) GetAssetID() string {
	if m != nil {
		return m.AssetID
	}
	return ""
}

func (m *StAssetIDAndInfo) GetInfo() *StakerAssetInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

// OperatorAssetIDAndInfos is a helper struct to be used in the genesis state.
// It is used to store the operator address and its asset id and info.
type StAssetIDAndInfos struct {
	// staker_id is the address of the operator.
	StakerID string `protobuf:"bytes,1,opt,name=staker_id,json=stakerId,proto3" json:"staker_id,omitempty"`
	// asset_id_and_infos is the list of asset id and infos that exist
	// for the staker.
	AssetIdAndInfos []*StAssetIDAndInfo `protobuf:"bytes,2,rep,name=asset_id_and_infos,json=assetIdAndInfos,proto3" json:"asset_id_and_infos,omitempty"`
}

func (m *StAssetIDAndInfos) Reset()         { *m = StAssetIDAndInfos{} }
func (m *StAssetIDAndInfos) String() string { return proto.CompactTextString(m) }
func (*StAssetIDAndInfos) ProtoMessage()    {}
func (*StAssetIDAndInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_caf4f124d39d82ce, []int{4}
}
func (m *StAssetIDAndInfos) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StAssetIDAndInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StAssetIDAndInfos.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StAssetIDAndInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StAssetIDAndInfos.Merge(m, src)
}
func (m *StAssetIDAndInfos) XXX_Size() int {
	return m.Size()
}
func (m *StAssetIDAndInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_StAssetIDAndInfos.DiscardUnknown(m)
}

var xxx_messageInfo_StAssetIDAndInfos proto.InternalMessageInfo

func (m *StAssetIDAndInfos) GetStakerID() string {
	if m != nil {
		return m.StakerID
	}
	return ""
}

func (m *StAssetIDAndInfos) GetAssetIdAndInfos() []*StAssetIDAndInfo {
	if m != nil {
		return m.AssetIdAndInfos
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "exocore.assets.v1.GenesisState")
	proto.RegisterType((*OpAssetIDAndInfo)(nil), "exocore.assets.v1.OpAssetIDAndInfo")
	proto.RegisterType((*OpAssetIDAndInfos)(nil), "exocore.assets.v1.OpAssetIDAndInfos")
	proto.RegisterType((*StAssetIDAndInfo)(nil), "exocore.assets.v1.StAssetIDAndInfo")
	proto.RegisterType((*StAssetIDAndInfos)(nil), "exocore.assets.v1.StAssetIDAndInfos")
}

func init() { proto.RegisterFile("exocore/assets/v1/genesis.proto", fileDescriptor_caf4f124d39d82ce) }

var fileDescriptor_caf4f124d39d82ce = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0xb6, 0x74, 0xad, 0x5b, 0xb4, 0xce, 0xda, 0x21, 0xf4, 0x90, 0x4e, 0x01, 0xa1,
	0xed, 0x92, 0xb0, 0x21, 0x01, 0xd7, 0x76, 0x43, 0xa8, 0x48, 0xc0, 0x94, 0x5e, 0x10, 0x42, 0xaa,
	0xbc, 0xc4, 0xcb, 0xa2, 0x30, 0x3b, 0xb2, 0xcd, 0x28, 0xdf, 0x62, 0x7c, 0xab, 0x1d, 0x77, 0xe4,
	0x54, 0xa1, 0xf4, 0x8b, 0xa0, 0xd8, 0x4e, 0x4b, 0x93, 0x54, 0x62, 0xb7, 0xc8, 0xef, 0xff, 0x7e,
	0x7f, 0xbf, 0xff, 0x8b, 0xc1, 0x10, 0xcf, 0xa9, 0x4f, 0x19, 0x76, 0x11, 0xe7, 0x58, 0x70, 0xf7,
	0xe6, 0xd8, 0x0d, 0x31, 0xc1, 0x3c, 0xe2, 0x4e, 0xc2, 0xa8, 0xa0, 0x70, 0x4f, 0x0b, 0x1c, 0x25,
	0x70, 0x6e, 0x8e, 0x07, 0xfb, 0x21, 0x0d, 0xa9, 0xac, 0xba, 0xd9, 0x97, 0x12, 0x0e, 0xac, 0x32,
	0x29, 0x41, 0x0c, 0x5d, 0x6b, 0xd0, 0x60, 0x50, 0xae, 0x8b, 0xb9, 0xaa, 0xd9, 0xbf, 0x1a, 0xa0,
	0xf7, 0x4e, 0xd9, 0x4e, 0x05, 0x12, 0x18, 0xbe, 0x06, 0x2d, 0xd5, 0x6c, 0x1a, 0x07, 0xc6, 0x61,
	0xf7, 0xe4, 0x89, 0x53, 0xba, 0x86, 0x73, 0x2e, 0x05, 0xe3, 0xe6, 0xdd, 0x62, 0x58, 0xf3, 0xb4,
	0x1c, 0x7e, 0x00, 0x8f, 0xfd, 0x6f, 0x11, 0x26, 0x62, 0xe6, 0x5f, 0xa1, 0x88, 0x70, 0xb3, 0x7e,
	0xd0, 0x38, 0xec, 0x9e, 0xd8, 0x15, 0xfd, 0xa7, 0x52, 0x77, 0x9a, 0xc9, 0x26, 0xe4, 0x92, 0x6a,
	0x50, 0xcf, 0x5f, 0x1f, 0x73, 0x38, 0x02, 0x2d, 0x41, 0x63, 0x4c, 0xb8, 0xd9, 0x90, 0x9c, 0xa7,
	0x15, 0x9c, 0xa9, 0x40, 0x71, 0x44, 0xc2, 0x51, 0x76, 0xf0, 0x0f, 0x48, 0x37, 0xc2, 0xaf, 0x60,
	0x9f, 0x26, 0x98, 0x21, 0x41, 0xd9, 0x4c, 0x36, 0xcd, 0x22, 0x72, 0x49, 0xb9, 0xd9, 0x94, 0xc0,
	0x67, 0x15, 0xc0, 0x4f, 0x89, 0x62, 0x9d, 0x8d, 0x48, 0x90, 0x01, 0xf3, 0x19, 0x61, 0xce, 0x59,
	0x59, 0x71, 0xf8, 0x19, 0x40, 0x2e, 0x50, 0x8c, 0x37, 0xd9, 0x8f, 0xb6, 0xb2, 0xa7, 0xa2, 0x9a,
	0xdd, 0x57, 0x94, 0x35, 0xd9, 0x16, 0xa0, 0x5f, 0xbc, 0x08, 0x7c, 0x0e, 0xda, 0xda, 0x26, 0x90,
	0x8b, 0xe9, 0x8c, 0xbb, 0xe9, 0x62, 0xb8, 0xa3, 0x55, 0xde, 0x8e, 0x2c, 0x4e, 0x02, 0xf8, 0x06,
	0x34, 0xb3, 0x8b, 0x98, 0x75, 0xb9, 0xbc, 0xea, 0x19, 0x0b, 0xa3, 0x78, 0xb2, 0xc3, 0xbe, 0x35,
	0xc0, 0x5e, 0x69, 0x7e, 0x78, 0x04, 0xfa, 0xeb, 0x0c, 0x83, 0x80, 0x61, 0xae, 0x7e, 0x8c, 0x8e,
	0xb7, 0xbb, 0xca, 0x44, 0x1d, 0xc3, 0x73, 0x00, 0xf3, 0x2b, 0xce, 0x10, 0x09, 0x74, 0x20, 0xf5,
	0xad, 0xdb, 0x2b, 0x9a, 0x79, 0xbb, 0x7a, 0x88, 0xdc, 0xdc, 0x66, 0xa0, 0x5f, 0x4c, 0xed, 0xbf,
	0x83, 0x78, 0xb5, 0x11, 0x84, 0xbd, 0xe5, 0xef, 0xc1, 0x95, 0x31, 0x94, 0x56, 0x05, 0x8f, 0x40,
	0x47, 0x2f, 0x7b, 0x65, 0xdb, 0x4b, 0x17, 0xc3, 0xb6, 0x62, 0x4c, 0xce, 0xbc, 0xb6, 0x2a, 0x4f,
	0x82, 0x07, 0xc7, 0x50, 0x34, 0x2b, 0xc5, 0x30, 0x7e, 0x7f, 0x97, 0x5a, 0xc6, 0x7d, 0x6a, 0x19,
	0x7f, 0x52, 0xcb, 0xb8, 0x5d, 0x5a, 0xb5, 0xfb, 0xa5, 0x55, 0xfb, 0xbd, 0xb4, 0x6a, 0x5f, 0x5e,
	0x84, 0x91, 0xb8, 0xfa, 0x7e, 0xe1, 0xf8, 0xf4, 0xda, 0x7d, 0xab, 0xc8, 0x1f, 0xb1, 0xf8, 0x41,
	0x59, 0xec, 0xe6, 0x6f, 0x7e, 0x9e, 0xbf, 0x7a, 0xf1, 0x33, 0xc1, 0xfc, 0xa2, 0x25, 0x9f, 0xfd,
	0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x36, 0xfd, 0x31, 0x7e, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StakerAssetInfos) > 0 {
		for iNdEx := len(m.StakerAssetInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakerAssetInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.OperatorAssetInfos) > 0 {
		for iNdEx := len(m.OperatorAssetInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OperatorAssetInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClientChains) > 0 {
		for iNdEx := len(m.ClientChains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientChains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *OpAssetIDAndInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OpAssetIDAndInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OpAssetIDAndInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.AssetID) > 0 {
		i -= len(m.AssetID)
		copy(dAtA[i:], m.AssetID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AssetID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OpAssetIDAndInfos) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OpAssetIDAndInfos) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OpAssetIDAndInfos) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetIdAndInfos) > 0 {
		for iNdEx := len(m.AssetIdAndInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetIdAndInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StAssetIDAndInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StAssetIDAndInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StAssetIDAndInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.AssetID) > 0 {
		i -= len(m.AssetID)
		copy(dAtA[i:], m.AssetID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AssetID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StAssetIDAndInfos) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StAssetIDAndInfos) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StAssetIDAndInfos) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssetIdAndInfos) > 0 {
		for iNdEx := len(m.AssetIdAndInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetIdAndInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.StakerID) > 0 {
		i -= len(m.StakerID)
		copy(dAtA[i:], m.StakerID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.StakerID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClientChains) > 0 {
		for _, e := range m.ClientChains {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OperatorAssetInfos) > 0 {
		for _, e := range m.OperatorAssetInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StakerAssetInfos) > 0 {
		for _, e := range m.StakerAssetInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *OpAssetIDAndInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AssetID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *OpAssetIDAndInfos) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AssetIdAndInfos) > 0 {
		for _, e := range m.AssetIdAndInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *StAssetIDAndInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AssetID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *StAssetIDAndInfos) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StakerID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AssetIdAndInfos) > 0 {
		for _, e := range m.AssetIdAndInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientChains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientChains = append(m.ClientChains, ClientChainInfo{})
			if err := m.ClientChains[len(m.ClientChains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, StakingAssetInfo{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAssetInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAssetInfos = append(m.OperatorAssetInfos, OpAssetIDAndInfos{})
			if err := m.OperatorAssetInfos[len(m.OperatorAssetInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerAssetInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerAssetInfos = append(m.StakerAssetInfos, StAssetIDAndInfos{})
			if err := m.StakerAssetInfos[len(m.StakerAssetInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *OpAssetIDAndInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: OpAssetIDAndInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OpAssetIDAndInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &OperatorAssetInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *OpAssetIDAndInfos) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: OpAssetIDAndInfos: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OpAssetIDAndInfos: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetIdAndInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetIdAndInfos = append(m.AssetIdAndInfos, &OpAssetIDAndInfo{})
			if err := m.AssetIdAndInfos[len(m.AssetIdAndInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *StAssetIDAndInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: StAssetIDAndInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StAssetIDAndInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &StakerAssetInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *StAssetIDAndInfos) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: StAssetIDAndInfos: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StAssetIDAndInfos: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetIdAndInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetIdAndInfos = append(m.AssetIdAndInfos, &StAssetIDAndInfo{})
			if err := m.AssetIdAndInfos[len(m.AssetIdAndInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

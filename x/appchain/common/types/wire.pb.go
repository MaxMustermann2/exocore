// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/appchain/common/v1/wire.proto

package types

import (
	fmt "fmt"
	types "github.com/cometbft/cometbft/abci/types"
	_ "github.com/cosmos/cosmos-sdk/x/staking/types"
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

type HandshakeMetadata struct {
	// This address is where the subscriber chain will send the fees proportionally
	CoordinatorFeePoolAddr string `protobuf:"bytes,1,opt,name=coordinator_fee_pool_addr,json=coordinatorFeePoolAddr,proto3" json:"coordinator_fee_pool_addr,omitempty"`
	Version                string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *HandshakeMetadata) Reset()         { *m = HandshakeMetadata{} }
func (m *HandshakeMetadata) String() string { return proto.CompactTextString(m) }
func (*HandshakeMetadata) ProtoMessage()    {}
func (*HandshakeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_646142158918d547, []int{0}
}
func (m *HandshakeMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HandshakeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HandshakeMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HandshakeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandshakeMetadata.Merge(m, src)
}
func (m *HandshakeMetadata) XXX_Size() int {
	return m.Size()
}
func (m *HandshakeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_HandshakeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_HandshakeMetadata proto.InternalMessageInfo

func (m *HandshakeMetadata) GetCoordinatorFeePoolAddr() string {
	if m != nil {
		return m.CoordinatorFeePoolAddr
	}
	return ""
}

func (m *HandshakeMetadata) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// ValidatorSetChangePacketData is sent from the coordinator chain to the subscriber chain
// containing the new validator set and the id of the validator set change.
type ValidatorSetChangePacketData struct {
	// validator_updates is the edits to the existing validator set
	ValidatorUpdates []types.ValidatorUpdate `protobuf:"bytes,1,rep,name=validator_updates,json=validatorUpdates,proto3" json:"validator_updates" yaml:"validator_updates"`
	// valset_update_id is the id of the validator set change
	ValsetUpdateID uint64 `protobuf:"varint,2,opt,name=valset_update_id,json=valsetUpdateId,proto3" json:"valset_update_id,omitempty"`
	// slash_acks is the list of consensus addresses slashed on the coordinator chain,
	// in response to such requests from the subscriber chain.
	SlashAcks []string `protobuf:"bytes,3,rep,name=slash_acks,json=slashAcks,proto3" json:"slash_acks,omitempty"`
}

func (m *ValidatorSetChangePacketData) Reset()         { *m = ValidatorSetChangePacketData{} }
func (m *ValidatorSetChangePacketData) String() string { return proto.CompactTextString(m) }
func (*ValidatorSetChangePacketData) ProtoMessage()    {}
func (*ValidatorSetChangePacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_646142158918d547, []int{1}
}
func (m *ValidatorSetChangePacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSetChangePacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSetChangePacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSetChangePacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSetChangePacketData.Merge(m, src)
}
func (m *ValidatorSetChangePacketData) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSetChangePacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSetChangePacketData.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSetChangePacketData proto.InternalMessageInfo

func (m *ValidatorSetChangePacketData) GetValidatorUpdates() []types.ValidatorUpdate {
	if m != nil {
		return m.ValidatorUpdates
	}
	return nil
}

func (m *ValidatorSetChangePacketData) GetValsetUpdateID() uint64 {
	if m != nil {
		return m.ValsetUpdateID
	}
	return 0
}

func (m *ValidatorSetChangePacketData) GetSlashAcks() []string {
	if m != nil {
		return m.SlashAcks
	}
	return nil
}

func init() {
	proto.RegisterType((*HandshakeMetadata)(nil), "exocore.appchain.common.v1.HandshakeMetadata")
	proto.RegisterType((*ValidatorSetChangePacketData)(nil), "exocore.appchain.common.v1.ValidatorSetChangePacketData")
}

func init() {
	proto.RegisterFile("exocore/appchain/common/v1/wire.proto", fileDescriptor_646142158918d547)
}

var fileDescriptor_646142158918d547 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xe3, 0x37, 0xaf, 0x40, 0x59, 0xa4, 0x8a, 0x5a, 0x08, 0x99, 0x00, 0x8e, 0x15, 0x81,
	0x94, 0x93, 0x57, 0x81, 0x53, 0x11, 0x97, 0x86, 0x82, 0xe0, 0x00, 0xaa, 0x52, 0xd1, 0x03, 0x17,
	0x6b, 0xb2, 0x3b, 0xd8, 0x2b, 0xdb, 0x3b, 0xd6, 0xee, 0xd6, 0x6d, 0xbf, 0x05, 0x1f, 0xab, 0xc7,
	0x1e, 0x39, 0x45, 0x28, 0xb9, 0x72, 0xe2, 0x13, 0x20, 0xdb, 0x31, 0x14, 0xf5, 0x36, 0x7f, 0x7e,
	0xf3, 0xec, 0x8e, 0x9e, 0x61, 0xcf, 0xf1, 0x82, 0x04, 0x19, 0xe4, 0x50, 0x55, 0x22, 0x03, 0xa5,
	0xb9, 0xa0, 0xb2, 0x24, 0xcd, 0xeb, 0x39, 0x3f, 0x57, 0x06, 0xe3, 0xca, 0x90, 0x23, 0x7f, 0xbc,
	0xc3, 0xe2, 0x1e, 0x8b, 0x3b, 0x2c, 0xae, 0xe7, 0xe3, 0x67, 0x82, 0x6c, 0x49, 0x96, 0x5b, 0x07,
	0xb9, 0xd2, 0x29, 0xaf, 0xe7, 0x2b, 0x74, 0x30, 0xef, 0xf3, 0x4e, 0x61, 0xfc, 0x20, 0xa5, 0x94,
	0xda, 0x90, 0x37, 0xd1, 0xae, 0xfa, 0xd8, 0xa1, 0x96, 0x68, 0x4a, 0xa5, 0x1d, 0x87, 0x95, 0x50,
	0xdc, 0x5d, 0x56, 0x68, 0xbb, 0xe6, 0x34, 0x63, 0xfb, 0xef, 0x41, 0x4b, 0x9b, 0x41, 0x8e, 0x1f,
	0xd1, 0x81, 0x04, 0x07, 0xfe, 0x01, 0x7b, 0x24, 0x88, 0x8c, 0x54, 0x1a, 0x1c, 0x99, 0xe4, 0x2b,
	0x62, 0x52, 0x11, 0x15, 0x09, 0x48, 0x69, 0x02, 0x2f, 0xf2, 0x66, 0xa3, 0xe5, 0xc3, 0x1b, 0xc0,
	0x3b, 0xc4, 0x63, 0xa2, 0xe2, 0x50, 0x4a, 0xe3, 0x07, 0xec, 0x6e, 0x8d, 0xc6, 0x2a, 0xd2, 0xc1,
	0x7f, 0x2d, 0xd8, 0xa7, 0xd3, 0x9f, 0x1e, 0x7b, 0x72, 0x0a, 0x85, 0x92, 0xcd, 0xc8, 0x09, 0xba,
	0x37, 0x19, 0xe8, 0x14, 0x8f, 0x41, 0xe4, 0xe8, 0x8e, 0x9a, 0x57, 0x89, 0xed, 0xd7, 0x7d, 0x3f,
	0x39, 0xab, 0x24, 0x38, 0xb4, 0x81, 0x17, 0x0d, 0x67, 0xf7, 0x5e, 0x44, 0xf1, 0xdf, 0x1d, 0xe2,
	0x66, 0x87, 0xf8, 0x8f, 0xd2, 0xe7, 0x16, 0x5c, 0x44, 0x57, 0xeb, 0xc9, 0xe0, 0xd7, 0x7a, 0x12,
	0x5c, 0x42, 0x59, 0xbc, 0x9a, 0xde, 0x12, 0x9a, 0x2e, 0xef, 0xd7, 0xff, 0x8e, 0x58, 0xff, 0x35,
	0x6b, 0x6a, 0x16, 0xdd, 0x0e, 0x4a, 0x94, 0x6c, 0x3f, 0xfd, 0xff, 0xc2, 0xdf, 0xac, 0x27, 0x7b,
	0xa7, 0x6d, 0xaf, 0x83, 0x3f, 0x1c, 0x2d, 0xf7, 0xea, 0x9b, 0xb9, 0xf4, 0x9f, 0x32, 0x66, 0x0b,
	0xb0, 0x59, 0x02, 0x22, 0xb7, 0xc1, 0x30, 0x1a, 0xce, 0x46, 0xcb, 0x51, 0x5b, 0x39, 0x14, 0xb9,
	0x5d, 0x9c, 0x5c, 0x6d, 0x42, 0xef, 0x7a, 0x13, 0x7a, 0x3f, 0x36, 0xa1, 0xf7, 0x6d, 0x1b, 0x0e,
	0xae, 0xb7, 0xe1, 0xe0, 0xfb, 0x36, 0x1c, 0x7c, 0x39, 0x48, 0x95, 0xcb, 0xce, 0x56, 0x8d, 0xbb,
	0xfc, 0x6d, 0x67, 0xf9, 0x27, 0x74, 0xe7, 0x64, 0x72, 0xde, 0x1f, 0xca, 0xc5, 0xad, 0x53, 0x69,
	0x3d, 0x5b, 0xdd, 0x69, 0x4d, 0x7b, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x29, 0x74, 0xa9, 0x8a,
	0x52, 0x02, 0x00, 0x00,
}

func (m *HandshakeMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandshakeMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HandshakeMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintWire(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CoordinatorFeePoolAddr) > 0 {
		i -= len(m.CoordinatorFeePoolAddr)
		copy(dAtA[i:], m.CoordinatorFeePoolAddr)
		i = encodeVarintWire(dAtA, i, uint64(len(m.CoordinatorFeePoolAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSetChangePacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSetChangePacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSetChangePacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SlashAcks) > 0 {
		for iNdEx := len(m.SlashAcks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SlashAcks[iNdEx])
			copy(dAtA[i:], m.SlashAcks[iNdEx])
			i = encodeVarintWire(dAtA, i, uint64(len(m.SlashAcks[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ValsetUpdateID != 0 {
		i = encodeVarintWire(dAtA, i, uint64(m.ValsetUpdateID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ValidatorUpdates) > 0 {
		for iNdEx := len(m.ValidatorUpdates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorUpdates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWire(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintWire(dAtA []byte, offset int, v uint64) int {
	offset -= sovWire(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HandshakeMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CoordinatorFeePoolAddr)
	if l > 0 {
		n += 1 + l + sovWire(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovWire(uint64(l))
	}
	return n
}

func (m *ValidatorSetChangePacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorUpdates) > 0 {
		for _, e := range m.ValidatorUpdates {
			l = e.Size()
			n += 1 + l + sovWire(uint64(l))
		}
	}
	if m.ValsetUpdateID != 0 {
		n += 1 + sovWire(uint64(m.ValsetUpdateID))
	}
	if len(m.SlashAcks) > 0 {
		for _, s := range m.SlashAcks {
			l = len(s)
			n += 1 + l + sovWire(uint64(l))
		}
	}
	return n
}

func sovWire(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWire(x uint64) (n int) {
	return sovWire(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HandshakeMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
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
			return fmt.Errorf("proto: HandshakeMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandshakeMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorFeePoolAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoordinatorFeePoolAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWire
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
func (m *ValidatorSetChangePacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWire
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
			return fmt.Errorf("proto: ValidatorSetChangePacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSetChangePacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorUpdates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorUpdates = append(m.ValidatorUpdates, types.ValidatorUpdate{})
			if err := m.ValidatorUpdates[len(m.ValidatorUpdates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValsetUpdateID", wireType)
			}
			m.ValsetUpdateID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ValsetUpdateID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashAcks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWire
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
				return ErrInvalidLengthWire
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWire
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashAcks = append(m.SlashAcks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWire(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWire
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
func skipWire(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWire
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
					return 0, ErrIntOverflowWire
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
					return 0, ErrIntOverflowWire
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
				return 0, ErrInvalidLengthWire
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWire
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWire
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWire        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWire          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWire = fmt.Errorf("proto: unexpected end of group")
)

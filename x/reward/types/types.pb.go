// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/reward/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Pool struct {
	Name    string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rewards []Pool_Reward `protobuf:"bytes,2,rep,name=rewards,proto3" json:"rewards"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_620cd6dbeff3c5e2, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

type Pool_Reward struct {
	Validator github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=validator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"validator,omitempty"`
	Coins     github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Pool_Reward) Reset()         { *m = Pool_Reward{} }
func (m *Pool_Reward) String() string { return proto.CompactTextString(m) }
func (*Pool_Reward) ProtoMessage()    {}
func (*Pool_Reward) Descriptor() ([]byte, []int) {
	return fileDescriptor_620cd6dbeff3c5e2, []int{0, 0}
}
func (m *Pool_Reward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool_Reward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool_Reward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool_Reward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool_Reward.Merge(m, src)
}
func (m *Pool_Reward) XXX_Size() int {
	return m.Size()
}
func (m *Pool_Reward) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool_Reward.DiscardUnknown(m)
}

var xxx_messageInfo_Pool_Reward proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Pool)(nil), "exocore.reward.Pool")
	proto.RegisterType((*Pool_Reward)(nil), "exocore.reward.Pool.Reward")
}

func init() { proto.RegisterFile("exocore/reward/types.proto", fileDescriptor_620cd6dbeff3c5e2) }

var fileDescriptor_620cd6dbeff3c5e2 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x3d, 0x4e, 0xc3, 0x30,
	0x18, 0x8d, 0x4b, 0x29, 0xaa, 0x41, 0x0c, 0x16, 0x43, 0x09, 0x92, 0x5b, 0x98, 0xba, 0xd4, 0xa6,
	0x30, 0x32, 0x20, 0xc2, 0x01, 0x40, 0x19, 0x18, 0xd8, 0x9c, 0xc4, 0x2a, 0x11, 0x4d, 0xbe, 0xca,
	0x0e, 0xa5, 0xdc, 0x82, 0x89, 0x43, 0x70, 0x06, 0x0e, 0x90, 0xb1, 0x23, 0x53, 0x81, 0xe4, 0x16,
	0x4c, 0xc8, 0x71, 0x22, 0xe8, 0xc6, 0x94, 0x27, 0xbd, 0xbc, 0x3f, 0x7f, 0xd8, 0x95, 0x0b, 0x08,
	0x41, 0x49, 0xae, 0xe4, 0xa3, 0x50, 0x11, 0xcf, 0x9e, 0x66, 0x52, 0xb3, 0x99, 0x82, 0x0c, 0xc8,
	0x6e, 0xcd, 0x31, 0xcb, 0xb9, 0x7b, 0x13, 0x98, 0x40, 0x45, 0x71, 0x83, 0xec, 0x5f, 0x2e, 0x0d,
	0x41, 0x27, 0xa0, 0x79, 0x20, 0xb4, 0xe4, 0xf3, 0x71, 0x20, 0x33, 0x31, 0xe6, 0x21, 0xc4, 0xa9,
	0xe5, 0x8f, 0x5e, 0x5a, 0xb8, 0x7d, 0x0d, 0x30, 0x25, 0x04, 0xb7, 0x53, 0x91, 0xc8, 0x1e, 0x1a,
	0xa0, 0x61, 0xd7, 0xaf, 0x30, 0x39, 0xc3, 0x5b, 0xd6, 0x5c, 0xf7, 0x5a, 0x83, 0x8d, 0xe1, 0xf6,
	0xc9, 0x01, 0x5b, 0x0f, 0x65, 0x46, 0xca, 0xfc, 0x0a, 0x7b, 0xed, 0x7c, 0xd5, 0x77, 0xfc, 0x46,
	0xe1, 0xbe, 0x21, 0xdc, 0xb1, 0x0c, 0xb9, 0xc2, 0xdd, 0xb9, 0x98, 0xc6, 0x91, 0xc8, 0x40, 0x55,
	0x01, 0x3b, 0xde, 0xf8, 0x7b, 0xd5, 0x1f, 0x4d, 0xe2, 0xec, 0xee, 0x21, 0x60, 0x21, 0x24, 0xbc,
	0xae, 0x69, 0x3f, 0x23, 0x1d, 0xdd, 0xd7, 0x5b, 0x6f, 0xc4, 0xf4, 0x22, 0x8a, 0x94, 0xd4, 0xda,
	0xff, 0xf5, 0x20, 0x02, 0x6f, 0x9a, 0x0d, 0x4d, 0xad, 0x7d, 0x66, 0x75, 0xcc, 0xac, 0x64, 0xf5,
	0x4a, 0x76, 0x09, 0x71, 0xea, 0x1d, 0x9b, 0x52, 0xaf, 0x1f, 0xfd, 0xe1, 0x3f, 0xb2, 0x8c, 0x40,
	0xfb, 0xd6, 0xd9, 0x3b, 0xcf, 0xbf, 0xa8, 0x93, 0x17, 0x14, 0x2d, 0x0b, 0x8a, 0x3e, 0x0b, 0x8a,
	0x9e, 0x4b, 0xea, 0x2c, 0x4b, 0xea, 0xbc, 0x97, 0xd4, 0xb9, 0x3d, 0xfc, 0x63, 0xd7, 0xdc, 0x68,
	0xb1, 0x76, 0xa5, 0xa0, 0x53, 0x3d, 0xf0, 0xe9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xe4,
	0xb5, 0x9c, 0xc4, 0x01, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pool_Reward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool_Reward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool_Reward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Pool_Reward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, Pool_Reward{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Pool_Reward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Reward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = append(m.Validator[:0], dAtA[iNdEx:postIndex]...)
			if m.Validator == nil {
				m.Validator = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)

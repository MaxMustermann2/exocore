// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: exocore/appchain/coordinator/v1/params.proto

package types

import (
	fmt "fmt"
	types "github.com/ExocoreNetwork/exocore/x/epochs/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_07_tendermint "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params is the parameters for the appchain coordinator module.
type Params struct {
	// template_client is the IBC template client.
	TemplateClient *_07_tendermint.ClientState `protobuf:"bytes,1,opt,name=template_client,json=templateClient,proto3" json:"template_client,omitempty"`
	// trusting_period_fraction is the multiplier applied on the subscriber's unbonding duration to determine the IBC trusting period.
	TrustingPeriodFraction string `protobuf:"bytes,2,opt,name=trusting_period_fraction,json=trustingPeriodFraction,proto3" json:"trusting_period_fraction,omitempty"`
	// ibc_timeout_period is the timeout period for IBC packets. While our system is largely created with epochs as a unit of time (and not standard durations), this is an exception since it is used directly by the IBC codebase.
	IBCTimeoutPeriod time.Duration `protobuf:"bytes,3,opt,name=ibc_timeout_period,json=ibcTimeoutPeriod,proto3,stdduration" json:"ibc_timeout_period"`
	// init_timeout_period is the period within which the subscriber chain must make a connection with the coordinator, after being spawned.
	InitTimeoutPeriod types.Epoch `protobuf:"bytes,4,opt,name=init_timeout_period,json=initTimeoutPeriod,proto3" json:"init_timeout_period"`
	// vsc_timeout_period is the period within which the subscriber chain must respond to a VSC request, after it is sent.
	VSCTimeoutPeriod types.Epoch `protobuf:"bytes,5,opt,name=vsc_timeout_period,json=vscTimeoutPeriod,proto3" json:"vsc_timeout_period"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_153952e3bb2d49c4, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetTemplateClient() *_07_tendermint.ClientState {
	if m != nil {
		return m.TemplateClient
	}
	return nil
}

func (m *Params) GetTrustingPeriodFraction() string {
	if m != nil {
		return m.TrustingPeriodFraction
	}
	return ""
}

func (m *Params) GetIBCTimeoutPeriod() time.Duration {
	if m != nil {
		return m.IBCTimeoutPeriod
	}
	return 0
}

func (m *Params) GetInitTimeoutPeriod() types.Epoch {
	if m != nil {
		return m.InitTimeoutPeriod
	}
	return types.Epoch{}
}

func (m *Params) GetVSCTimeoutPeriod() types.Epoch {
	if m != nil {
		return m.VSCTimeoutPeriod
	}
	return types.Epoch{}
}

func init() {
	proto.RegisterType((*Params)(nil), "exocore.appchain.coordinator.v1.Params")
}

func init() {
	proto.RegisterFile("exocore/appchain/coordinator/v1/params.proto", fileDescriptor_153952e3bb2d49c4)
}

var fileDescriptor_153952e3bb2d49c4 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x14, 0xdc, 0xd0, 0x52, 0x89, 0x20, 0xc1, 0x12, 0x10, 0x0a, 0x15, 0xca, 0x56, 0x9c, 0x2a, 0x81,
	0x6c, 0x2d, 0x5c, 0x38, 0x71, 0xd8, 0x52, 0x24, 0x2e, 0x55, 0x95, 0x56, 0x08, 0x71, 0x59, 0x1c,
	0xaf, 0x9b, 0x7d, 0x62, 0xe3, 0x67, 0x39, 0x2f, 0xa1, 0xfc, 0x05, 0x47, 0x7e, 0x83, 0xbf, 0xe8,
	0xb1, 0x47, 0x4e, 0x05, 0xed, 0xfe, 0x08, 0x8a, 0x1d, 0x8b, 0x5d, 0x40, 0xea, 0xcd, 0x7e, 0x33,
	0x6f, 0x66, 0x32, 0x71, 0xfc, 0x4c, 0x9d, 0xa3, 0x44, 0xab, 0xb8, 0x30, 0x46, 0xce, 0x05, 0x68,
	0x2e, 0x11, 0xed, 0x0c, 0xb4, 0x20, 0xb4, 0xbc, 0x1d, 0x73, 0x23, 0xac, 0xa8, 0x6a, 0x66, 0x2c,
	0x12, 0x26, 0xa3, 0x9e, 0xcd, 0x02, 0x9b, 0xad, 0xb1, 0x59, 0x3b, 0xde, 0x7d, 0x50, 0x62, 0x89,
	0x8e, 0xcb, 0xbb, 0x93, 0x5f, 0xdb, 0xcd, 0x4a, 0xc4, 0x72, 0xa1, 0xb8, 0xbb, 0x15, 0xcd, 0x19,
	0x9f, 0x35, 0x56, 0x10, 0xa0, 0xee, 0x71, 0x0e, 0x85, 0xe4, 0x0b, 0x28, 0xe7, 0x24, 0x17, 0xa0,
	0x34, 0xd5, 0x9c, 0x94, 0x9e, 0x29, 0x5b, 0x81, 0xa6, 0x2e, 0xc3, 0x9f, 0x5b, 0x10, 0x0c, 0xa9,
	0x95, 0x41, 0x39, 0xaf, 0x3b, 0x8e, 0x3f, 0x79, 0xfc, 0xc9, 0xf7, 0xad, 0x78, 0xe7, 0xd8, 0x05,
	0x4f, 0x4e, 0xe3, 0xbb, 0xa4, 0x2a, 0xb3, 0x10, 0xa4, 0xa6, 0x5e, 0x3d, 0x8d, 0xf6, 0xa2, 0xfd,
	0xdb, 0xcf, 0x9f, 0x32, 0x28, 0x24, 0x5b, 0x77, 0x65, 0x6b, 0x3e, 0xed, 0x98, 0x1d, 0xb8, 0xe9,
	0x09, 0x09, 0x52, 0xf9, 0x9d, 0xa0, 0xe1, 0x87, 0xc9, 0xcb, 0x38, 0x25, 0xdb, 0xd4, 0x04, 0xba,
	0x9c, 0x1a, 0x65, 0x01, 0x67, 0xd3, 0x33, 0x2b, 0x64, 0xf7, 0x4d, 0xe9, 0x8d, 0xbd, 0x68, 0xff,
	0x56, 0xfe, 0x30, 0xe0, 0xc7, 0x0e, 0x7e, 0xd3, 0xa3, 0x89, 0x8a, 0x13, 0x28, 0xe4, 0x94, 0xa0,
	0x52, 0xd8, 0x50, 0xbf, 0x9c, 0x6e, 0xb9, 0x48, 0x8f, 0x98, 0x2f, 0x8a, 0x85, 0xa2, 0xd8, 0xeb,
	0xbe, 0xa8, 0xc9, 0xe3, 0x8b, 0xab, 0xd1, 0x60, 0x79, 0x35, 0x1a, 0xbe, 0x9d, 0x1c, 0x9c, 0xfa,
	0x5d, 0x2f, 0xfc, 0xed, 0xe7, 0x28, 0xca, 0x87, 0x50, 0xc8, 0x8d, 0x69, 0x72, 0x14, 0xdf, 0x07,
	0x0d, 0xf4, 0xb7, 0xcf, 0xb6, 0xf3, 0x49, 0x59, 0xf8, 0x8f, 0x7d, 0x6b, 0xed, 0x98, 0x1d, 0x76,
	0xa7, 0xc9, 0x76, 0x67, 0x93, 0xdf, 0xeb, 0x56, 0x37, 0xf5, 0x3e, 0xc6, 0x49, 0x5b, 0xff, 0x13,
	0xfb, 0xe6, 0x35, 0x72, 0x69, 0x48, 0xfd, 0xee, 0x64, 0x33, 0x75, 0x3e, 0x6c, 0xeb, 0xcd, 0xc4,
	0x93, 0xf7, 0x17, 0xcb, 0x2c, 0xba, 0x5c, 0x66, 0xd1, 0xaf, 0x65, 0x16, 0x7d, 0x5d, 0x65, 0x83,
	0xcb, 0x55, 0x36, 0xf8, 0xb1, 0xca, 0x06, 0x1f, 0x5e, 0x95, 0x40, 0xf3, 0xa6, 0x60, 0x12, 0x2b,
	0x7e, 0xe8, 0x9d, 0x8e, 0x14, 0x7d, 0x46, 0xfb, 0x89, 0x87, 0x77, 0x70, 0xfe, 0xff, 0xf7, 0x4b,
	0x5f, 0x8c, 0xaa, 0x8b, 0x1d, 0x57, 0xe7, 0x8b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xa3,
	0x9c, 0xaa, 0xec, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.VSCTimeoutPeriod.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.InitTimeoutPeriod.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.IBCTimeoutPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.IBCTimeoutPeriod):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	if len(m.TrustingPeriodFraction) > 0 {
		i -= len(m.TrustingPeriodFraction)
		copy(dAtA[i:], m.TrustingPeriodFraction)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TrustingPeriodFraction)))
		i--
		dAtA[i] = 0x12
	}
	if m.TemplateClient != nil {
		{
			size, err := m.TemplateClient.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TemplateClient != nil {
		l = m.TemplateClient.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TrustingPeriodFraction)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.IBCTimeoutPeriod)
	n += 1 + l + sovParams(uint64(l))
	l = m.InitTimeoutPeriod.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.VSCTimeoutPeriod.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TemplateClient", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TemplateClient == nil {
				m.TemplateClient = &_07_tendermint.ClientState{}
			}
			if err := m.TemplateClient.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustingPeriodFraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustingPeriodFraction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCTimeoutPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.IBCTimeoutPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitTimeoutPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitTimeoutPeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VSCTimeoutPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VSCTimeoutPeriod.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)

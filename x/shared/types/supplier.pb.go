// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/shared/supplier.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Supplier is the type defining the actor in Pocket Network that provides RPC services.
type Supplier struct {
	Address                       string                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stake                         *types.Coin              `protobuf:"bytes,2,opt,name=stake,proto3" json:"stake,omitempty"`
	Services                      []*SupplierServiceConfig `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
	UnstakeCommitSessionEndHeight int64                    `protobuf:"varint,4,opt,name=unstake_commit_session_end_height,json=unstakeCommitSessionEndHeight,proto3" json:"unstake_commit_session_end_height,omitempty"`
}

func (m *Supplier) Reset()         { *m = Supplier{} }
func (m *Supplier) String() string { return proto.CompactTextString(m) }
func (*Supplier) ProtoMessage()    {}
func (*Supplier) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a189b52ba503cf2, []int{0}
}
func (m *Supplier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Supplier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Supplier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Supplier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Supplier.Merge(m, src)
}
func (m *Supplier) XXX_Size() int {
	return m.Size()
}
func (m *Supplier) XXX_DiscardUnknown() {
	xxx_messageInfo_Supplier.DiscardUnknown(m)
}

var xxx_messageInfo_Supplier proto.InternalMessageInfo

func (m *Supplier) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Supplier) GetStake() *types.Coin {
	if m != nil {
		return m.Stake
	}
	return nil
}

func (m *Supplier) GetServices() []*SupplierServiceConfig {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Supplier) GetUnstakeCommitSessionEndHeight() int64 {
	if m != nil {
		return m.UnstakeCommitSessionEndHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Supplier)(nil), "poktroll.shared.Supplier")
}

func init() { proto.RegisterFile("poktroll/shared/supplier.proto", fileDescriptor_4a189b52ba503cf2) }

var fileDescriptor_4a189b52ba503cf2 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbf, 0x4e, 0xeb, 0x30,
	0x18, 0xc5, 0xeb, 0xdb, 0x7b, 0x2f, 0x25, 0x1d, 0x90, 0x22, 0x86, 0xb4, 0x52, 0xad, 0xc0, 0x80,
	0xb2, 0xd4, 0x56, 0xcb, 0x13, 0xd0, 0x0a, 0xa9, 0xac, 0xc9, 0xc6, 0x12, 0xe5, 0x8f, 0x49, 0xac,
	0x36, 0x76, 0xe4, 0xcf, 0x2d, 0xf0, 0x16, 0x3c, 0x0c, 0x0f, 0xc1, 0x58, 0x31, 0x31, 0xa2, 0xf6,
	0x39, 0x90, 0x50, 0x63, 0xa7, 0x03, 0x8c, 0x9f, 0x7e, 0xbf, 0xe4, 0x1c, 0x1f, 0x07, 0xd7, 0x72,
	0xa9, 0x95, 0x5c, 0xad, 0x28, 0x94, 0x89, 0x62, 0x39, 0x85, 0x75, 0x5d, 0xaf, 0x38, 0x53, 0xa4,
	0x56, 0x52, 0x4b, 0xf7, 0xac, 0xe5, 0xc4, 0xf0, 0xe1, 0x20, 0x93, 0x50, 0x49, 0x88, 0x1b, 0x4c,
	0xcd, 0x61, 0xdc, 0x21, 0x36, 0x17, 0x4d, 0x13, 0x60, 0x74, 0x33, 0x49, 0x99, 0x4e, 0x26, 0x34,
	0x93, 0x5c, 0x58, 0x3e, 0xfa, 0x95, 0xc5, 0xd4, 0x86, 0x67, 0xcc, 0xe0, 0xcb, 0x2f, 0xe4, 0xf4,
	0x22, 0x9b, 0xee, 0x4e, 0x9d, 0x93, 0x24, 0xcf, 0x15, 0x03, 0xf0, 0x90, 0x8f, 0x82, 0xd3, 0x99,
	0xf7, 0xfe, 0x3a, 0x3e, 0xb7, 0x71, 0x37, 0x86, 0x44, 0x5a, 0x71, 0x51, 0x84, 0xad, 0xe8, 0x52,
	0xe7, 0x1f, 0xe8, 0x64, 0xc9, 0xbc, 0x3f, 0x3e, 0x0a, 0xfa, 0xd3, 0x01, 0xb1, 0xfa, 0xa1, 0x0f,
	0xb1, 0x7d, 0xc8, 0x5c, 0x72, 0x11, 0x1a, 0xcf, 0x9d, 0x39, 0x3d, 0x5b, 0x01, 0xbc, 0xae, 0xdf,
	0x0d, 0xfa, 0xd3, 0x2b, 0xf2, 0xe3, 0xbd, 0xa4, 0x6d, 0x14, 0x19, 0x71, 0x2e, 0xc5, 0x03, 0x2f,
	0xc2, 0xe3, 0x77, 0xee, 0xc2, 0xb9, 0x58, 0x8b, 0xe6, 0x77, 0x71, 0x26, 0xab, 0x8a, 0xeb, 0x18,
	0x18, 0x00, 0x97, 0x22, 0x66, 0x22, 0x8f, 0x4b, 0xc6, 0x8b, 0x52, 0x7b, 0x7f, 0x7d, 0x14, 0x74,
	0xc3, 0x91, 0x15, 0xe7, 0x8d, 0x17, 0x19, 0xed, 0x56, 0xe4, 0x8b, 0x46, 0x9a, 0xdd, 0xbd, 0xed,
	0x30, 0xda, 0xee, 0x30, 0xfa, 0xdc, 0x61, 0xf4, 0xb2, 0xc7, 0x9d, 0xed, 0x1e, 0x77, 0x3e, 0xf6,
	0xb8, 0x73, 0x4f, 0x0b, 0xae, 0xcb, 0x75, 0x4a, 0x32, 0x59, 0xd1, 0x43, 0xbf, 0xb1, 0x60, 0xfa,
	0x51, 0xaa, 0x25, 0x3d, 0x0e, 0xfa, 0xd4, 0x4e, 0xaa, 0x9f, 0x6b, 0x06, 0xe9, 0xff, 0x66, 0xd1,
	0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x3c, 0xfd, 0xfe, 0xde, 0x01, 0x00, 0x00,
}

func (m *Supplier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Supplier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Supplier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnstakeCommitSessionEndHeight != 0 {
		i = encodeVarintSupplier(dAtA, i, uint64(m.UnstakeCommitSessionEndHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Services[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSupplier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Stake != nil {
		{
			size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSupplier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSupplier(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSupplier(dAtA []byte, offset int, v uint64) int {
	offset -= sovSupplier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Supplier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSupplier(uint64(l))
	}
	if m.Stake != nil {
		l = m.Stake.Size()
		n += 1 + l + sovSupplier(uint64(l))
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovSupplier(uint64(l))
		}
	}
	if m.UnstakeCommitSessionEndHeight != 0 {
		n += 1 + sovSupplier(uint64(m.UnstakeCommitSessionEndHeight))
	}
	return n
}

func sovSupplier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSupplier(x uint64) (n int) {
	return sovSupplier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Supplier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSupplier
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
			return fmt.Errorf("proto: Supplier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Supplier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stake == nil {
				m.Stake = &types.Coin{}
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
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
				return ErrInvalidLengthSupplier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSupplier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &SupplierServiceConfig{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnstakeCommitSessionEndHeight", wireType)
			}
			m.UnstakeCommitSessionEndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSupplier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnstakeCommitSessionEndHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSupplier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSupplier
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
func skipSupplier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSupplier
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
					return 0, ErrIntOverflowSupplier
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
					return 0, ErrIntOverflowSupplier
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
				return 0, ErrInvalidLengthSupplier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSupplier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSupplier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSupplier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSupplier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSupplier = fmt.Errorf("proto: unexpected end of group")
)

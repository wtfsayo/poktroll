// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/tokenomics/event.proto

package tokenomics

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	proof "github.com/pokt-network/poktroll/proto/types/proof"
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

// EventClaimExpired is an event emitted during settlement whenever a claim requiring
// an on-chain proof doesn't have one. The claim cannot be settled, leading to that work
// never being rewarded.
type EventClaimExpired struct {
	Claim           *proof.Claim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim"`
	NumRelays       uint64       `protobuf:"varint,2,opt,name=num_relays,json=numRelays,proto3" json:"num_relays"`
	NumComputeUnits uint64       `protobuf:"varint,3,opt,name=num_compute_units,json=numComputeUnits,proto3" json:"num_compute_units"`
}

func (m *EventClaimExpired) Reset()         { *m = EventClaimExpired{} }
func (m *EventClaimExpired) String() string { return proto.CompactTextString(m) }
func (*EventClaimExpired) ProtoMessage()    {}
func (*EventClaimExpired) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{0}
}
func (m *EventClaimExpired) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClaimExpired) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClaimExpired.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClaimExpired) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClaimExpired.Merge(m, src)
}
func (m *EventClaimExpired) XXX_Size() int {
	return m.Size()
}
func (m *EventClaimExpired) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClaimExpired.DiscardUnknown(m)
}

var xxx_messageInfo_EventClaimExpired proto.InternalMessageInfo

func (m *EventClaimExpired) GetClaim() *proof.Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (m *EventClaimExpired) GetNumRelays() uint64 {
	if m != nil {
		return m.NumRelays
	}
	return 0
}

func (m *EventClaimExpired) GetNumComputeUnits() uint64 {
	if m != nil {
		return m.NumComputeUnits
	}
	return 0
}

// EventClaimSettled is an event emitted whenever a claim is settled.
// The proof_required determines whether the claim requires a proof that has been submitted or not
type EventClaimSettled struct {
	Claim            *proof.Claim                 `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim"`
	NumRelays        uint64                       `protobuf:"varint,2,opt,name=num_relays,json=numRelays,proto3" json:"num_relays"`
	NumComputeUnits  uint64                       `protobuf:"varint,3,opt,name=num_compute_units,json=numComputeUnits,proto3" json:"num_compute_units"`
	ProofRequirement proof.ProofRequirementReason `protobuf:"varint,4,opt,name=proof_requirement,json=proofRequirement,proto3,enum=poktroll.proof.ProofRequirementReason" json:"proof_requirement"`
}

func (m *EventClaimSettled) Reset()         { *m = EventClaimSettled{} }
func (m *EventClaimSettled) String() string { return proto.CompactTextString(m) }
func (*EventClaimSettled) ProtoMessage()    {}
func (*EventClaimSettled) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{1}
}
func (m *EventClaimSettled) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClaimSettled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClaimSettled.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClaimSettled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClaimSettled.Merge(m, src)
}
func (m *EventClaimSettled) XXX_Size() int {
	return m.Size()
}
func (m *EventClaimSettled) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClaimSettled.DiscardUnknown(m)
}

var xxx_messageInfo_EventClaimSettled proto.InternalMessageInfo

func (m *EventClaimSettled) GetClaim() *proof.Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (m *EventClaimSettled) GetNumRelays() uint64 {
	if m != nil {
		return m.NumRelays
	}
	return 0
}

func (m *EventClaimSettled) GetNumComputeUnits() uint64 {
	if m != nil {
		return m.NumComputeUnits
	}
	return 0
}

func (m *EventClaimSettled) GetProofRequirement() proof.ProofRequirementReason {
	if m != nil {
		return m.ProofRequirement
	}
	return proof.ProofRequirementReason_NOT_REQUIRED
}

// EventRelayMiningDifficultyUpdated is an event emitted whenever the relay mining difficulty is updated
// for a given service.
type EventRelayMiningDifficultyUpdated struct {
	ServiceId                string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	PrevTargetHashHexEncoded string `protobuf:"bytes,2,opt,name=prev_target_hash_hex_encoded,json=prevTargetHashHexEncoded,proto3" json:"prev_target_hash_hex_encoded,omitempty"`
	NewTargetHashHexEncoded  string `protobuf:"bytes,3,opt,name=new_target_hash_hex_encoded,json=newTargetHashHexEncoded,proto3" json:"new_target_hash_hex_encoded,omitempty"`
	PrevNumRelaysEma         uint64 `protobuf:"varint,4,opt,name=prev_num_relays_ema,json=prevNumRelaysEma,proto3" json:"prev_num_relays_ema,omitempty"`
	NewNumRelaysEma          uint64 `protobuf:"varint,5,opt,name=new_num_relays_ema,json=newNumRelaysEma,proto3" json:"new_num_relays_ema,omitempty"`
}

func (m *EventRelayMiningDifficultyUpdated) Reset()         { *m = EventRelayMiningDifficultyUpdated{} }
func (m *EventRelayMiningDifficultyUpdated) String() string { return proto.CompactTextString(m) }
func (*EventRelayMiningDifficultyUpdated) ProtoMessage()    {}
func (*EventRelayMiningDifficultyUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{2}
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRelayMiningDifficultyUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRelayMiningDifficultyUpdated.Merge(m, src)
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Size() int {
	return m.Size()
}
func (m *EventRelayMiningDifficultyUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRelayMiningDifficultyUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_EventRelayMiningDifficultyUpdated proto.InternalMessageInfo

func (m *EventRelayMiningDifficultyUpdated) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetPrevTargetHashHexEncoded() string {
	if m != nil {
		return m.PrevTargetHashHexEncoded
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetNewTargetHashHexEncoded() string {
	if m != nil {
		return m.NewTargetHashHexEncoded
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetPrevNumRelaysEma() uint64 {
	if m != nil {
		return m.PrevNumRelaysEma
	}
	return 0
}

func (m *EventRelayMiningDifficultyUpdated) GetNewNumRelaysEma() uint64 {
	if m != nil {
		return m.NewNumRelaysEma
	}
	return 0
}

// EventApplicationOverserviced is emitted when an application has less stake
// than the expected burn.
type EventApplicationOverserviced struct {
	ApplicationAddr string      `protobuf:"bytes,1,opt,name=application_addr,json=applicationAddr,proto3" json:"application_addr,omitempty"`
	ExpectedBurn    *types.Coin `protobuf:"bytes,2,opt,name=expected_burn,json=expectedBurn,proto3" json:"expected_burn,omitempty"`
	EffectiveBurn   *types.Coin `protobuf:"bytes,3,opt,name=effective_burn,json=effectiveBurn,proto3" json:"effective_burn,omitempty"`
}

func (m *EventApplicationOverserviced) Reset()         { *m = EventApplicationOverserviced{} }
func (m *EventApplicationOverserviced) String() string { return proto.CompactTextString(m) }
func (*EventApplicationOverserviced) ProtoMessage()    {}
func (*EventApplicationOverserviced) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{3}
}
func (m *EventApplicationOverserviced) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventApplicationOverserviced) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventApplicationOverserviced.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventApplicationOverserviced) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventApplicationOverserviced.Merge(m, src)
}
func (m *EventApplicationOverserviced) XXX_Size() int {
	return m.Size()
}
func (m *EventApplicationOverserviced) XXX_DiscardUnknown() {
	xxx_messageInfo_EventApplicationOverserviced.DiscardUnknown(m)
}

var xxx_messageInfo_EventApplicationOverserviced proto.InternalMessageInfo

func (m *EventApplicationOverserviced) GetApplicationAddr() string {
	if m != nil {
		return m.ApplicationAddr
	}
	return ""
}

func (m *EventApplicationOverserviced) GetExpectedBurn() *types.Coin {
	if m != nil {
		return m.ExpectedBurn
	}
	return nil
}

func (m *EventApplicationOverserviced) GetEffectiveBurn() *types.Coin {
	if m != nil {
		return m.EffectiveBurn
	}
	return nil
}

func init() {
	proto.RegisterType((*EventClaimExpired)(nil), "poktroll.tokenomics.EventClaimExpired")
	proto.RegisterType((*EventClaimSettled)(nil), "poktroll.tokenomics.EventClaimSettled")
	proto.RegisterType((*EventRelayMiningDifficultyUpdated)(nil), "poktroll.tokenomics.EventRelayMiningDifficultyUpdated")
	proto.RegisterType((*EventApplicationOverserviced)(nil), "poktroll.tokenomics.EventApplicationOverserviced")
}

func init() { proto.RegisterFile("poktroll/tokenomics/event.proto", fileDescriptor_a78874bbf91a58c7) }

var fileDescriptor_a78874bbf91a58c7 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0xcb, 0x6e, 0x13, 0x3d,
	0x14, 0xc7, 0x3b, 0xbd, 0x7c, 0x52, 0xdc, 0xaf, 0x37, 0x97, 0x8a, 0x50, 0xca, 0x24, 0x64, 0x81,
	0x8a, 0x50, 0x67, 0xd4, 0x22, 0xc1, 0x06, 0x55, 0x34, 0x25, 0x52, 0x59, 0x70, 0x1b, 0xe8, 0x86,
	0xcd, 0xc8, 0x19, 0x9f, 0x24, 0xa6, 0x33, 0xb6, 0xf1, 0x78, 0x26, 0xe9, 0x5b, 0xf0, 0x00, 0xbc,
	0x00, 0x0f, 0xc1, 0x1e, 0x89, 0x4d, 0x97, 0xac, 0x2a, 0xd4, 0xee, 0xfa, 0x14, 0xc8, 0x9e, 0x5c,
	0x46, 0x29, 0x88, 0x35, 0x9b, 0xc4, 0x3a, 0xff, 0xdf, 0x39, 0xe7, 0xef, 0xe3, 0xb1, 0x51, 0x4d,
	0x8a, 0x13, 0xad, 0x44, 0x1c, 0xfb, 0x5a, 0x9c, 0x00, 0x17, 0x09, 0x8b, 0x52, 0x1f, 0x72, 0xe0,
	0xda, 0x93, 0x4a, 0x68, 0x81, 0xd7, 0x47, 0x80, 0x37, 0x01, 0x36, 0x6f, 0x74, 0x45, 0x57, 0x58,
	0xdd, 0x37, 0xab, 0x02, 0xdd, 0x74, 0x23, 0x91, 0x26, 0x22, 0xf5, 0xdb, 0x24, 0x05, 0x3f, 0xdf,
	0x6d, 0x83, 0x26, 0xbb, 0x7e, 0x24, 0x18, 0x1f, 0xea, 0x9b, 0xe3, 0x5e, 0x52, 0x09, 0xd1, 0xf1,
	0xa3, 0x98, 0xb0, 0x64, 0xa8, 0xd5, 0xa7, 0x34, 0x05, 0x1f, 0x33, 0xa6, 0x20, 0x19, 0x1b, 0x69,
	0x7c, 0x75, 0xd0, 0x5a, 0xcb, 0x18, 0x3b, 0x34, 0x69, 0xad, 0x81, 0x64, 0x0a, 0x28, 0x7e, 0x84,
	0x16, 0x6c, 0x99, 0xaa, 0x53, 0x77, 0xb6, 0x17, 0xf7, 0x36, 0xbc, 0xb1, 0x5d, 0x5b, 0xc7, 0xb3,
	0x70, 0xb3, 0x72, 0x75, 0x5e, 0x2b, 0xb8, 0xa0, 0xf8, 0xc3, 0x3b, 0x08, 0xf1, 0x2c, 0x09, 0x15,
	0xc4, 0xe4, 0x34, 0xad, 0xce, 0xd6, 0x9d, 0xed, 0xf9, 0xe6, 0xf2, 0xd5, 0x79, 0xad, 0x14, 0x0d,
	0x2a, 0x3c, 0x4b, 0x02, 0xbb, 0xc4, 0x07, 0x68, 0xcd, 0x08, 0x91, 0x48, 0x64, 0xa6, 0x21, 0xcc,
	0x38, 0xd3, 0x69, 0x75, 0xce, 0x66, 0x6d, 0x5c, 0x9d, 0xd7, 0xae, 0x8b, 0xc1, 0x0a, 0xcf, 0x92,
	0xc3, 0x22, 0x72, 0x6c, 0x02, 0x8d, 0x2f, 0xb3, 0x65, 0xff, 0x6f, 0x41, 0xeb, 0xf8, 0x5f, 0xf2,
	0x8f, 0x3f, 0xa0, 0x35, 0x6b, 0x29, 0x2c, 0x1d, 0x4d, 0x75, 0xbe, 0xee, 0x6c, 0x2f, 0xef, 0xdd,
	0x9b, 0x76, 0xfd, 0xda, 0xfc, 0x06, 0x13, 0x2e, 0x00, 0x92, 0x0a, 0x5e, 0xb4, 0xba, 0x56, 0x24,
	0x58, 0x95, 0x53, 0x78, 0xe3, 0xf3, 0x2c, 0xba, 0x6b, 0x67, 0x65, 0xed, 0xbf, 0x60, 0x9c, 0xf1,
	0xee, 0x33, 0xd6, 0xe9, 0xb0, 0x28, 0x8b, 0xf5, 0xe9, 0xb1, 0xa4, 0x44, 0x03, 0xc5, 0x77, 0x10,
	0x4a, 0x41, 0xe5, 0x2c, 0x82, 0x90, 0x51, 0x3b, 0xc0, 0x4a, 0x50, 0x19, 0x46, 0x9e, 0x53, 0xbc,
	0x8f, 0xb6, 0xa4, 0x82, 0x3c, 0xd4, 0x44, 0x75, 0x41, 0x87, 0x3d, 0x92, 0xf6, 0xc2, 0x1e, 0x0c,
	0x42, 0xe0, 0x91, 0xa0, 0x40, 0xed, 0xd0, 0x2a, 0x41, 0xd5, 0x30, 0xef, 0x2c, 0x72, 0x44, 0xd2,
	0xde, 0x11, 0x0c, 0x5a, 0x85, 0x8e, 0x9f, 0xa0, 0xdb, 0x1c, 0xfa, 0x7f, 0x4c, 0x9f, 0xb3, 0xe9,
	0x37, 0x39, 0xf4, 0x7f, 0x9b, 0xbd, 0x83, 0xd6, 0x6d, 0xf7, 0xc9, 0x79, 0x84, 0x90, 0x10, 0x3b,
	0xb0, 0x79, 0xb3, 0x63, 0xc8, 0x5f, 0x8e, 0x4e, 0xa7, 0x95, 0x10, 0xfc, 0x00, 0x61, 0xd3, 0x6c,
	0x8a, 0x5e, 0xb0, 0xf4, 0x0a, 0x87, 0x7e, 0x19, 0x6e, 0x7c, 0x77, 0xd0, 0x96, 0x1d, 0xcf, 0x81,
	0x94, 0x31, 0x8b, 0x88, 0x66, 0x82, 0xbf, 0xca, 0x41, 0x0d, 0xf7, 0x4e, 0xf1, 0x7d, 0xb4, 0x4a,
	0x26, 0x52, 0x48, 0x28, 0x55, 0xc3, 0xf9, 0xac, 0x94, 0xe2, 0x07, 0x94, 0x2a, 0xbc, 0x8f, 0x96,
	0x60, 0x20, 0x21, 0xd2, 0x40, 0xc3, 0x76, 0xa6, 0xb8, 0x1d, 0xcb, 0xe2, 0xde, 0x2d, 0xaf, 0xb8,
	0xcc, 0x9e, 0xb9, 0xcc, 0xde, 0xf0, 0x32, 0x7b, 0x87, 0x82, 0xf1, 0xe0, 0xff, 0x11, 0xdf, 0xcc,
	0x14, 0xc7, 0x4f, 0xd1, 0x32, 0x74, 0x3a, 0x10, 0x69, 0x96, 0x43, 0x51, 0x60, 0xee, 0x6f, 0x05,
	0x96, 0xc6, 0x09, 0xa6, 0x42, 0xf3, 0xcd, 0xb7, 0x0b, 0xd7, 0x39, 0xbb, 0x70, 0x9d, 0x9f, 0x17,
	0xae, 0xf3, 0xe9, 0xd2, 0x9d, 0x39, 0xbb, 0x74, 0x67, 0x7e, 0x5c, 0xba, 0x33, 0xef, 0x1f, 0x77,
	0x99, 0xee, 0x65, 0x6d, 0x2f, 0x12, 0x89, 0x6f, 0xbe, 0xb0, 0x1d, 0x0e, 0xba, 0x2f, 0xd4, 0x89,
	0x5f, 0x7e, 0x2c, 0xb4, 0xf0, 0xf5, 0xa9, 0x84, 0xb4, 0xf4, 0x80, 0xb5, 0xff, 0xb3, 0xf1, 0x87,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x5d, 0x1c, 0x61, 0xde, 0x04, 0x00, 0x00,
}

func (m *EventClaimExpired) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClaimExpired) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClaimExpired) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumComputeUnits != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumComputeUnits))
		i--
		dAtA[i] = 0x18
	}
	if m.NumRelays != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumRelays))
		i--
		dAtA[i] = 0x10
	}
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventClaimSettled) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClaimSettled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClaimSettled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProofRequirement != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.ProofRequirement))
		i--
		dAtA[i] = 0x20
	}
	if m.NumComputeUnits != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumComputeUnits))
		i--
		dAtA[i] = 0x18
	}
	if m.NumRelays != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumRelays))
		i--
		dAtA[i] = 0x10
	}
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRelayMiningDifficultyUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRelayMiningDifficultyUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRelayMiningDifficultyUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewNumRelaysEma != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NewNumRelaysEma))
		i--
		dAtA[i] = 0x28
	}
	if m.PrevNumRelaysEma != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.PrevNumRelaysEma))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NewTargetHashHexEncoded) > 0 {
		i -= len(m.NewTargetHashHexEncoded)
		copy(dAtA[i:], m.NewTargetHashHexEncoded)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.NewTargetHashHexEncoded)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PrevTargetHashHexEncoded) > 0 {
		i -= len(m.PrevTargetHashHexEncoded)
		copy(dAtA[i:], m.PrevTargetHashHexEncoded)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PrevTargetHashHexEncoded)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServiceId) > 0 {
		i -= len(m.ServiceId)
		copy(dAtA[i:], m.ServiceId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ServiceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventApplicationOverserviced) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventApplicationOverserviced) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventApplicationOverserviced) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EffectiveBurn != nil {
		{
			size, err := m.EffectiveBurn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ExpectedBurn != nil {
		{
			size, err := m.ExpectedBurn.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ApplicationAddr) > 0 {
		i -= len(m.ApplicationAddr)
		copy(dAtA[i:], m.ApplicationAddr)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ApplicationAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventClaimExpired) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.NumRelays != 0 {
		n += 1 + sovEvent(uint64(m.NumRelays))
	}
	if m.NumComputeUnits != 0 {
		n += 1 + sovEvent(uint64(m.NumComputeUnits))
	}
	return n
}

func (m *EventClaimSettled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.NumRelays != 0 {
		n += 1 + sovEvent(uint64(m.NumRelays))
	}
	if m.NumComputeUnits != 0 {
		n += 1 + sovEvent(uint64(m.NumComputeUnits))
	}
	if m.ProofRequirement != 0 {
		n += 1 + sovEvent(uint64(m.ProofRequirement))
	}
	return n
}

func (m *EventRelayMiningDifficultyUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.PrevTargetHashHexEncoded)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.NewTargetHashHexEncoded)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.PrevNumRelaysEma != 0 {
		n += 1 + sovEvent(uint64(m.PrevNumRelaysEma))
	}
	if m.NewNumRelaysEma != 0 {
		n += 1 + sovEvent(uint64(m.NewNumRelaysEma))
	}
	return n
}

func (m *EventApplicationOverserviced) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ApplicationAddr)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.ExpectedBurn != nil {
		l = m.ExpectedBurn.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.EffectiveBurn != nil {
		l = m.EffectiveBurn.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventClaimExpired) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventClaimExpired: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClaimExpired: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &proof.Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRelays", wireType)
			}
			m.NumRelays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRelays |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumComputeUnits", wireType)
			}
			m.NumComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventClaimSettled) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventClaimSettled: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClaimSettled: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &proof.Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRelays", wireType)
			}
			m.NumRelays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRelays |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumComputeUnits", wireType)
			}
			m.NumComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofRequirement", wireType)
			}
			m.ProofRequirement = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProofRequirement |= proof.ProofRequirementReason(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventRelayMiningDifficultyUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventRelayMiningDifficultyUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRelayMiningDifficultyUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevTargetHashHexEncoded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevTargetHashHexEncoded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewTargetHashHexEncoded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewTargetHashHexEncoded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevNumRelaysEma", wireType)
			}
			m.PrevNumRelaysEma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevNumRelaysEma |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewNumRelaysEma", wireType)
			}
			m.NewNumRelaysEma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewNumRelaysEma |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventApplicationOverserviced) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventApplicationOverserviced: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventApplicationOverserviced: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedBurn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpectedBurn == nil {
				m.ExpectedBurn = &types.Coin{}
			}
			if err := m.ExpectedBurn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EffectiveBurn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EffectiveBurn == nil {
				m.EffectiveBurn = &types.Coin{}
			}
			if err := m.EffectiveBurn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)

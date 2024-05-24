// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package tokenomics

import (
	_ "cosmossdk.io/api/amino"
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_RelayMiningDifficulty                protoreflect.MessageDescriptor
	fd_RelayMiningDifficulty_service_id     protoreflect.FieldDescriptor
	fd_RelayMiningDifficulty_block_height   protoreflect.FieldDescriptor
	fd_RelayMiningDifficulty_num_relays_ema protoreflect.FieldDescriptor
	fd_RelayMiningDifficulty_difficulty     protoreflect.FieldDescriptor
)

func init() {
	file_poktroll_tokenomics_relay_mining_difficulty_proto_init()
	md_RelayMiningDifficulty = File_poktroll_tokenomics_relay_mining_difficulty_proto.Messages().ByName("RelayMiningDifficulty")
	fd_RelayMiningDifficulty_service_id = md_RelayMiningDifficulty.Fields().ByName("service_id")
	fd_RelayMiningDifficulty_block_height = md_RelayMiningDifficulty.Fields().ByName("block_height")
	fd_RelayMiningDifficulty_num_relays_ema = md_RelayMiningDifficulty.Fields().ByName("num_relays_ema")
	fd_RelayMiningDifficulty_difficulty = md_RelayMiningDifficulty.Fields().ByName("difficulty")
}

var _ protoreflect.Message = (*fastReflection_RelayMiningDifficulty)(nil)

type fastReflection_RelayMiningDifficulty RelayMiningDifficulty

func (x *RelayMiningDifficulty) ProtoReflect() protoreflect.Message {
	return (*fastReflection_RelayMiningDifficulty)(x)
}

func (x *RelayMiningDifficulty) slowProtoReflect() protoreflect.Message {
	mi := &file_poktroll_tokenomics_relay_mining_difficulty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_RelayMiningDifficulty_messageType fastReflection_RelayMiningDifficulty_messageType
var _ protoreflect.MessageType = fastReflection_RelayMiningDifficulty_messageType{}

type fastReflection_RelayMiningDifficulty_messageType struct{}

func (x fastReflection_RelayMiningDifficulty_messageType) Zero() protoreflect.Message {
	return (*fastReflection_RelayMiningDifficulty)(nil)
}
func (x fastReflection_RelayMiningDifficulty_messageType) New() protoreflect.Message {
	return new(fastReflection_RelayMiningDifficulty)
}
func (x fastReflection_RelayMiningDifficulty_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_RelayMiningDifficulty
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_RelayMiningDifficulty) Descriptor() protoreflect.MessageDescriptor {
	return md_RelayMiningDifficulty
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_RelayMiningDifficulty) Type() protoreflect.MessageType {
	return _fastReflection_RelayMiningDifficulty_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_RelayMiningDifficulty) New() protoreflect.Message {
	return new(fastReflection_RelayMiningDifficulty)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_RelayMiningDifficulty) Interface() protoreflect.ProtoMessage {
	return (*RelayMiningDifficulty)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_RelayMiningDifficulty) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ServiceId != "" {
		value := protoreflect.ValueOfString(x.ServiceId)
		if !f(fd_RelayMiningDifficulty_service_id, value) {
			return
		}
	}
	if x.BlockHeight != int64(0) {
		value := protoreflect.ValueOfInt64(x.BlockHeight)
		if !f(fd_RelayMiningDifficulty_block_height, value) {
			return
		}
	}
	if x.NumRelaysEma != uint64(0) {
		value := protoreflect.ValueOfUint64(x.NumRelaysEma)
		if !f(fd_RelayMiningDifficulty_num_relays_ema, value) {
			return
		}
	}
	if len(x.Difficulty) != 0 {
		value := protoreflect.ValueOfBytes(x.Difficulty)
		if !f(fd_RelayMiningDifficulty_difficulty, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_RelayMiningDifficulty) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "poktroll.tokenomics.RelayMiningDifficulty.service_id":
		return x.ServiceId != ""
	case "poktroll.tokenomics.RelayMiningDifficulty.block_height":
		return x.BlockHeight != int64(0)
	case "poktroll.tokenomics.RelayMiningDifficulty.num_relays_ema":
		return x.NumRelaysEma != uint64(0)
	case "poktroll.tokenomics.RelayMiningDifficulty.difficulty":
		return len(x.Difficulty) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.tokenomics.RelayMiningDifficulty"))
		}
		panic(fmt.Errorf("message poktroll.tokenomics.RelayMiningDifficulty does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RelayMiningDifficulty) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "poktroll.tokenomics.RelayMiningDifficulty.service_id":
		x.ServiceId = ""
	case "poktroll.tokenomics.RelayMiningDifficulty.block_height":
		x.BlockHeight = int64(0)
	case "poktroll.tokenomics.RelayMiningDifficulty.num_relays_ema":
		x.NumRelaysEma = uint64(0)
	case "poktroll.tokenomics.RelayMiningDifficulty.difficulty":
		x.Difficulty = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.tokenomics.RelayMiningDifficulty"))
		}
		panic(fmt.Errorf("message poktroll.tokenomics.RelayMiningDifficulty does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_RelayMiningDifficulty) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "poktroll.tokenomics.RelayMiningDifficulty.service_id":
		value := x.ServiceId
		return protoreflect.ValueOfString(value)
	case "poktroll.tokenomics.RelayMiningDifficulty.block_height":
		value := x.BlockHeight
		return protoreflect.ValueOfInt64(value)
	case "poktroll.tokenomics.RelayMiningDifficulty.num_relays_ema":
		value := x.NumRelaysEma
		return protoreflect.ValueOfUint64(value)
	case "poktroll.tokenomics.RelayMiningDifficulty.difficulty":
		value := x.Difficulty
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.tokenomics.RelayMiningDifficulty"))
		}
		panic(fmt.Errorf("message poktroll.tokenomics.RelayMiningDifficulty does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RelayMiningDifficulty) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "poktroll.tokenomics.RelayMiningDifficulty.service_id":
		x.ServiceId = value.Interface().(string)
	case "poktroll.tokenomics.RelayMiningDifficulty.block_height":
		x.BlockHeight = value.Int()
	case "poktroll.tokenomics.RelayMiningDifficulty.num_relays_ema":
		x.NumRelaysEma = value.Uint()
	case "poktroll.tokenomics.RelayMiningDifficulty.difficulty":
		x.Difficulty = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.tokenomics.RelayMiningDifficulty"))
		}
		panic(fmt.Errorf("message poktroll.tokenomics.RelayMiningDifficulty does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RelayMiningDifficulty) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "poktroll.tokenomics.RelayMiningDifficulty.service_id":
		panic(fmt.Errorf("field service_id of message poktroll.tokenomics.RelayMiningDifficulty is not mutable"))
	case "poktroll.tokenomics.RelayMiningDifficulty.block_height":
		panic(fmt.Errorf("field block_height of message poktroll.tokenomics.RelayMiningDifficulty is not mutable"))
	case "poktroll.tokenomics.RelayMiningDifficulty.num_relays_ema":
		panic(fmt.Errorf("field num_relays_ema of message poktroll.tokenomics.RelayMiningDifficulty is not mutable"))
	case "poktroll.tokenomics.RelayMiningDifficulty.difficulty":
		panic(fmt.Errorf("field difficulty of message poktroll.tokenomics.RelayMiningDifficulty is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.tokenomics.RelayMiningDifficulty"))
		}
		panic(fmt.Errorf("message poktroll.tokenomics.RelayMiningDifficulty does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_RelayMiningDifficulty) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "poktroll.tokenomics.RelayMiningDifficulty.service_id":
		return protoreflect.ValueOfString("")
	case "poktroll.tokenomics.RelayMiningDifficulty.block_height":
		return protoreflect.ValueOfInt64(int64(0))
	case "poktroll.tokenomics.RelayMiningDifficulty.num_relays_ema":
		return protoreflect.ValueOfUint64(uint64(0))
	case "poktroll.tokenomics.RelayMiningDifficulty.difficulty":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.tokenomics.RelayMiningDifficulty"))
		}
		panic(fmt.Errorf("message poktroll.tokenomics.RelayMiningDifficulty does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_RelayMiningDifficulty) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in poktroll.tokenomics.RelayMiningDifficulty", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_RelayMiningDifficulty) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_RelayMiningDifficulty) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_RelayMiningDifficulty) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_RelayMiningDifficulty) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*RelayMiningDifficulty)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ServiceId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BlockHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.BlockHeight))
		}
		if x.NumRelaysEma != 0 {
			n += 1 + runtime.Sov(uint64(x.NumRelaysEma))
		}
		l = len(x.Difficulty)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*RelayMiningDifficulty)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Difficulty) > 0 {
			i -= len(x.Difficulty)
			copy(dAtA[i:], x.Difficulty)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Difficulty)))
			i--
			dAtA[i] = 0x2a
		}
		if x.NumRelaysEma != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.NumRelaysEma))
			i--
			dAtA[i] = 0x20
		}
		if x.BlockHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BlockHeight))
			i--
			dAtA[i] = 0x10
		}
		if len(x.ServiceId) > 0 {
			i -= len(x.ServiceId)
			copy(dAtA[i:], x.ServiceId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ServiceId)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*RelayMiningDifficulty)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RelayMiningDifficulty: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: RelayMiningDifficulty: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ServiceId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
				}
				x.BlockHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BlockHeight |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field NumRelaysEma", wireType)
				}
				x.NumRelaysEma = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.NumRelaysEma |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Difficulty", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Difficulty = append(x.Difficulty[:0], dAtA[iNdEx:postIndex]...)
				if x.Difficulty == nil {
					x.Difficulty = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: poktroll/tokenomics/relay_mining_difficulty.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RelayMiningDifficulty is a message used to compute, store, update and access
// the latest relay mining difficulty for each service.
type RelayMiningDifficulty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The service ID this relay mining difficulty is associated with.
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// The block height at which this relay mining difficulty was computed.
	BlockHeight int64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// The latest exponential moving average of the number of relays for this service.
	NumRelaysEma uint64 `protobuf:"varint,4,opt,name=num_relays_ema,json=numRelaysEma,proto3" json:"num_relays_ema,omitempty"`
	// The latest relay mining difficulty for this service.
	Difficulty []byte `protobuf:"bytes,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
}

func (x *RelayMiningDifficulty) Reset() {
	*x = RelayMiningDifficulty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poktroll_tokenomics_relay_mining_difficulty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayMiningDifficulty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayMiningDifficulty) ProtoMessage() {}

// Deprecated: Use RelayMiningDifficulty.ProtoReflect.Descriptor instead.
func (*RelayMiningDifficulty) Descriptor() ([]byte, []int) {
	return file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDescGZIP(), []int{0}
}

func (x *RelayMiningDifficulty) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *RelayMiningDifficulty) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *RelayMiningDifficulty) GetNumRelaysEma() uint64 {
	if x != nil {
		return x.NumRelaysEma
	}
	return 0
}

func (x *RelayMiningDifficulty) GetDifficulty() []byte {
	if x != nil {
		return x.Difficulty
	}
	return nil
}

var File_poktroll_tokenomics_relay_mining_difficulty_proto protoreflect.FileDescriptor

var file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f,
	0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6f,
	0x6d, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x5f, 0x65, 0x6d, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x73,
	0x45, 0x6d, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x42, 0xc8, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6b, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x42,
	0x1a, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0xa2, 0x02, 0x03, 0x50, 0x54, 0x58, 0xaa, 0x02, 0x13, 0x50, 0x6f, 0x6b, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0xca,
	0x02, 0x13, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x6f, 0x6d, 0x69, 0x63, 0x73, 0xe2, 0x02, 0x1f, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x5c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x3a, 0x3a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDescOnce sync.Once
	file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDescData = file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDesc
)

func file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDescGZIP() []byte {
	file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDescOnce.Do(func() {
		file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDescData = protoimpl.X.CompressGZIP(file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDescData)
	})
	return file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDescData
}

var file_poktroll_tokenomics_relay_mining_difficulty_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_poktroll_tokenomics_relay_mining_difficulty_proto_goTypes = []interface{}{
	(*RelayMiningDifficulty)(nil), // 0: poktroll.tokenomics.RelayMiningDifficulty
}
var file_poktroll_tokenomics_relay_mining_difficulty_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_poktroll_tokenomics_relay_mining_difficulty_proto_init() }
func file_poktroll_tokenomics_relay_mining_difficulty_proto_init() {
	if File_poktroll_tokenomics_relay_mining_difficulty_proto != nil {
		return
	}
	file_poktroll_tokenomics_params_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_poktroll_tokenomics_relay_mining_difficulty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayMiningDifficulty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_poktroll_tokenomics_relay_mining_difficulty_proto_goTypes,
		DependencyIndexes: file_poktroll_tokenomics_relay_mining_difficulty_proto_depIdxs,
		MessageInfos:      file_poktroll_tokenomics_relay_mining_difficulty_proto_msgTypes,
	}.Build()
	File_poktroll_tokenomics_relay_mining_difficulty_proto = out.File
	file_poktroll_tokenomics_relay_mining_difficulty_proto_rawDesc = nil
	file_poktroll_tokenomics_relay_mining_difficulty_proto_goTypes = nil
	file_poktroll_tokenomics_relay_mining_difficulty_proto_depIdxs = nil
}

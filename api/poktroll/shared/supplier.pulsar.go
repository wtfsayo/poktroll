// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package shared

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Supplier_3_list)(nil)

type _Supplier_3_list struct {
	list *[]*SupplierServiceConfig
}

func (x *_Supplier_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Supplier_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Supplier_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SupplierServiceConfig)
	(*x.list)[i] = concreteValue
}

func (x *_Supplier_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*SupplierServiceConfig)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Supplier_3_list) AppendMutable() protoreflect.Value {
	v := new(SupplierServiceConfig)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Supplier_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Supplier_3_list) NewElement() protoreflect.Value {
	v := new(SupplierServiceConfig)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Supplier_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Supplier                                   protoreflect.MessageDescriptor
	fd_Supplier_address                           protoreflect.FieldDescriptor
	fd_Supplier_stake                             protoreflect.FieldDescriptor
	fd_Supplier_services                          protoreflect.FieldDescriptor
	fd_Supplier_unstake_commit_session_end_height protoreflect.FieldDescriptor
)

func init() {
	file_poktroll_shared_supplier_proto_init()
	md_Supplier = File_poktroll_shared_supplier_proto.Messages().ByName("Supplier")
	fd_Supplier_address = md_Supplier.Fields().ByName("address")
	fd_Supplier_stake = md_Supplier.Fields().ByName("stake")
	fd_Supplier_services = md_Supplier.Fields().ByName("services")
	fd_Supplier_unstake_commit_session_end_height = md_Supplier.Fields().ByName("unstake_commit_session_end_height")
}

var _ protoreflect.Message = (*fastReflection_Supplier)(nil)

type fastReflection_Supplier Supplier

func (x *Supplier) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Supplier)(x)
}

func (x *Supplier) slowProtoReflect() protoreflect.Message {
	mi := &file_poktroll_shared_supplier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Supplier_messageType fastReflection_Supplier_messageType
var _ protoreflect.MessageType = fastReflection_Supplier_messageType{}

type fastReflection_Supplier_messageType struct{}

func (x fastReflection_Supplier_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Supplier)(nil)
}
func (x fastReflection_Supplier_messageType) New() protoreflect.Message {
	return new(fastReflection_Supplier)
}
func (x fastReflection_Supplier_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Supplier
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Supplier) Descriptor() protoreflect.MessageDescriptor {
	return md_Supplier
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Supplier) Type() protoreflect.MessageType {
	return _fastReflection_Supplier_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Supplier) New() protoreflect.Message {
	return new(fastReflection_Supplier)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Supplier) Interface() protoreflect.ProtoMessage {
	return (*Supplier)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Supplier) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_Supplier_address, value) {
			return
		}
	}
	if x.Stake != nil {
		value := protoreflect.ValueOfMessage(x.Stake.ProtoReflect())
		if !f(fd_Supplier_stake, value) {
			return
		}
	}
	if len(x.Services) != 0 {
		value := protoreflect.ValueOfList(&_Supplier_3_list{list: &x.Services})
		if !f(fd_Supplier_services, value) {
			return
		}
	}
	if x.UnstakeCommitSessionEndHeight != uint64(0) {
		value := protoreflect.ValueOfUint64(x.UnstakeCommitSessionEndHeight)
		if !f(fd_Supplier_unstake_commit_session_end_height, value) {
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
func (x *fastReflection_Supplier) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "poktroll.shared.Supplier.address":
		return x.Address != ""
	case "poktroll.shared.Supplier.stake":
		return x.Stake != nil
	case "poktroll.shared.Supplier.services":
		return len(x.Services) != 0
	case "poktroll.shared.Supplier.unstake_commit_session_end_height":
		return x.UnstakeCommitSessionEndHeight != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.shared.Supplier"))
		}
		panic(fmt.Errorf("message poktroll.shared.Supplier does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Supplier) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "poktroll.shared.Supplier.address":
		x.Address = ""
	case "poktroll.shared.Supplier.stake":
		x.Stake = nil
	case "poktroll.shared.Supplier.services":
		x.Services = nil
	case "poktroll.shared.Supplier.unstake_commit_session_end_height":
		x.UnstakeCommitSessionEndHeight = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.shared.Supplier"))
		}
		panic(fmt.Errorf("message poktroll.shared.Supplier does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Supplier) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "poktroll.shared.Supplier.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "poktroll.shared.Supplier.stake":
		value := x.Stake
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "poktroll.shared.Supplier.services":
		if len(x.Services) == 0 {
			return protoreflect.ValueOfList(&_Supplier_3_list{})
		}
		listValue := &_Supplier_3_list{list: &x.Services}
		return protoreflect.ValueOfList(listValue)
	case "poktroll.shared.Supplier.unstake_commit_session_end_height":
		value := x.UnstakeCommitSessionEndHeight
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.shared.Supplier"))
		}
		panic(fmt.Errorf("message poktroll.shared.Supplier does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Supplier) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "poktroll.shared.Supplier.address":
		x.Address = value.Interface().(string)
	case "poktroll.shared.Supplier.stake":
		x.Stake = value.Message().Interface().(*v1beta1.Coin)
	case "poktroll.shared.Supplier.services":
		lv := value.List()
		clv := lv.(*_Supplier_3_list)
		x.Services = *clv.list
	case "poktroll.shared.Supplier.unstake_commit_session_end_height":
		x.UnstakeCommitSessionEndHeight = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.shared.Supplier"))
		}
		panic(fmt.Errorf("message poktroll.shared.Supplier does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Supplier) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "poktroll.shared.Supplier.stake":
		if x.Stake == nil {
			x.Stake = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.Stake.ProtoReflect())
	case "poktroll.shared.Supplier.services":
		if x.Services == nil {
			x.Services = []*SupplierServiceConfig{}
		}
		value := &_Supplier_3_list{list: &x.Services}
		return protoreflect.ValueOfList(value)
	case "poktroll.shared.Supplier.address":
		panic(fmt.Errorf("field address of message poktroll.shared.Supplier is not mutable"))
	case "poktroll.shared.Supplier.unstake_commit_session_end_height":
		panic(fmt.Errorf("field unstake_commit_session_end_height of message poktroll.shared.Supplier is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.shared.Supplier"))
		}
		panic(fmt.Errorf("message poktroll.shared.Supplier does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Supplier) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "poktroll.shared.Supplier.address":
		return protoreflect.ValueOfString("")
	case "poktroll.shared.Supplier.stake":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "poktroll.shared.Supplier.services":
		list := []*SupplierServiceConfig{}
		return protoreflect.ValueOfList(&_Supplier_3_list{list: &list})
	case "poktroll.shared.Supplier.unstake_commit_session_end_height":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: poktroll.shared.Supplier"))
		}
		panic(fmt.Errorf("message poktroll.shared.Supplier does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Supplier) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in poktroll.shared.Supplier", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Supplier) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Supplier) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Supplier) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Supplier) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Supplier)
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
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Stake != nil {
			l = options.Size(x.Stake)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Services) > 0 {
			for _, e := range x.Services {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.UnstakeCommitSessionEndHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.UnstakeCommitSessionEndHeight))
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
		x := input.Message.Interface().(*Supplier)
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
		if x.UnstakeCommitSessionEndHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.UnstakeCommitSessionEndHeight))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Services) > 0 {
			for iNdEx := len(x.Services) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Services[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if x.Stake != nil {
			encoded, err := options.Marshal(x.Stake)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
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
		x := input.Message.Interface().(*Supplier)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Supplier: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Supplier: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Stake == nil {
					x.Stake = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Stake); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Services = append(x.Services, &SupplierServiceConfig{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Services[len(x.Services)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UnstakeCommitSessionEndHeight", wireType)
				}
				x.UnstakeCommitSessionEndHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.UnstakeCommitSessionEndHeight |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: poktroll/shared/supplier.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Supplier is the type defining the actor in Pocket Network that provides RPC services.
type Supplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`   // The Bech32 address of the supplier using cosmos' ScalarDescriptor to ensure deterministic encoding
	Stake    *v1beta1.Coin            `protobuf:"bytes,2,opt,name=stake,proto3" json:"stake,omitempty"`       // The total amount of uPOKT the supplier has staked
	Services []*SupplierServiceConfig `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"` // The service configs this supplier can support
	// The session end height corresponding the the unstake tx commit height.
	// If the supplier is not unstaking, this value will be 0.
	UnstakeCommitSessionEndHeight uint64 `protobuf:"varint,4,opt,name=unstake_commit_session_end_height,json=unstakeCommitSessionEndHeight,proto3" json:"unstake_commit_session_end_height,omitempty"`
}

func (x *Supplier) Reset() {
	*x = Supplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_poktroll_shared_supplier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supplier) ProtoMessage() {}

// Deprecated: Use Supplier.ProtoReflect.Descriptor instead.
func (*Supplier) Descriptor() ([]byte, []int) {
	return file_poktroll_shared_supplier_proto_rawDescGZIP(), []int{0}
}

func (x *Supplier) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Supplier) GetStake() *v1beta1.Coin {
	if x != nil {
		return x.Stake
	}
	return nil
}

func (x *Supplier) GetServices() []*SupplierServiceConfig {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Supplier) GetUnstakeCommitSessionEndHeight() uint64 {
	if x != nil {
		return x.UnstakeCommitSessionEndHeight
	}
	return 0
}

var File_poktroll_shared_supplier_proto protoreflect.FileDescriptor

var file_poktroll_shared_supplier_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6f,
	0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x08,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x42, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x48, 0x0a, 0x21, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1d, 0x75, 0x6e,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0xa3, 0x01, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x42, 0x0d, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e,
	0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0xa2, 0x02, 0x03, 0x50, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x50,
	0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0xca, 0x02,
	0x0f, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0xe2, 0x02, 0x1b, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x10, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x3a, 0x3a, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_poktroll_shared_supplier_proto_rawDescOnce sync.Once
	file_poktroll_shared_supplier_proto_rawDescData = file_poktroll_shared_supplier_proto_rawDesc
)

func file_poktroll_shared_supplier_proto_rawDescGZIP() []byte {
	file_poktroll_shared_supplier_proto_rawDescOnce.Do(func() {
		file_poktroll_shared_supplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_poktroll_shared_supplier_proto_rawDescData)
	})
	return file_poktroll_shared_supplier_proto_rawDescData
}

var file_poktroll_shared_supplier_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_poktroll_shared_supplier_proto_goTypes = []interface{}{
	(*Supplier)(nil),              // 0: poktroll.shared.Supplier
	(*v1beta1.Coin)(nil),          // 1: cosmos.base.v1beta1.Coin
	(*SupplierServiceConfig)(nil), // 2: poktroll.shared.SupplierServiceConfig
}
var file_poktroll_shared_supplier_proto_depIdxs = []int32{
	1, // 0: poktroll.shared.Supplier.stake:type_name -> cosmos.base.v1beta1.Coin
	2, // 1: poktroll.shared.Supplier.services:type_name -> poktroll.shared.SupplierServiceConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_poktroll_shared_supplier_proto_init() }
func file_poktroll_shared_supplier_proto_init() {
	if File_poktroll_shared_supplier_proto != nil {
		return
	}
	file_poktroll_shared_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_poktroll_shared_supplier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supplier); i {
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
			RawDescriptor: file_poktroll_shared_supplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_poktroll_shared_supplier_proto_goTypes,
		DependencyIndexes: file_poktroll_shared_supplier_proto_depIdxs,
		MessageInfos:      file_poktroll_shared_supplier_proto_msgTypes,
	}.Build()
	File_poktroll_shared_supplier_proto = out.File
	file_poktroll_shared_supplier_proto_rawDesc = nil
	file_poktroll_shared_supplier_proto_goTypes = nil
	file_poktroll_shared_supplier_proto_depIdxs = nil
}

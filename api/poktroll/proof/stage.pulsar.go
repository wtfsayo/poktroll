// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package proof

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: poktroll/proof/stage.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClaimProofStage int32

const (
	ClaimProofStage_CLAIMED ClaimProofStage = 0
	ClaimProofStage_PROVEN  ClaimProofStage = 1
	ClaimProofStage_SETTLED ClaimProofStage = 2
	ClaimProofStage_EXPIRED ClaimProofStage = 3
)

// Enum value maps for ClaimProofStage.
var (
	ClaimProofStage_name = map[int32]string{
		0: "CLAIMED",
		1: "PROVEN",
		2: "SETTLED",
		3: "EXPIRED",
	}
	ClaimProofStage_value = map[string]int32{
		"CLAIMED": 0,
		"PROVEN":  1,
		"SETTLED": 2,
		"EXPIRED": 3,
	}
)

func (x ClaimProofStage) Enum() *ClaimProofStage {
	p := new(ClaimProofStage)
	*p = x
	return p
}

func (x ClaimProofStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClaimProofStage) Descriptor() protoreflect.EnumDescriptor {
	return file_poktroll_proof_stage_proto_enumTypes[0].Descriptor()
}

func (ClaimProofStage) Type() protoreflect.EnumType {
	return &file_poktroll_proof_stage_proto_enumTypes[0]
}

func (x ClaimProofStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClaimProofStage.Descriptor instead.
func (ClaimProofStage) EnumDescriptor() ([]byte, []int) {
	return file_poktroll_proof_stage_proto_rawDescGZIP(), []int{0}
}

var File_poktroll_proof_stage_proto protoreflect.FileDescriptor

var file_poktroll_proof_stage_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x2f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x6f,
	0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x2a, 0x44, 0x0a, 0x0f,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x52, 0x4f, 0x56, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x54, 0x54,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44,
	0x10, 0x03, 0x42, 0x9a, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6b, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73,
	0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x6b, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0xa2, 0x02, 0x03, 0x50, 0x50, 0x58, 0xaa, 0x02,
	0x0e, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0xca,
	0x02, 0x0e, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0xe2, 0x02, 0x1a, 0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x5c, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f,
	0x50, 0x6f, 0x6b, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_poktroll_proof_stage_proto_rawDescOnce sync.Once
	file_poktroll_proof_stage_proto_rawDescData = file_poktroll_proof_stage_proto_rawDesc
)

func file_poktroll_proof_stage_proto_rawDescGZIP() []byte {
	file_poktroll_proof_stage_proto_rawDescOnce.Do(func() {
		file_poktroll_proof_stage_proto_rawDescData = protoimpl.X.CompressGZIP(file_poktroll_proof_stage_proto_rawDescData)
	})
	return file_poktroll_proof_stage_proto_rawDescData
}

var file_poktroll_proof_stage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_poktroll_proof_stage_proto_goTypes = []interface{}{
	(ClaimProofStage)(0), // 0: poktroll.proof.ClaimProofStage
}
var file_poktroll_proof_stage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_poktroll_proof_stage_proto_init() }
func file_poktroll_proof_stage_proto_init() {
	if File_poktroll_proof_stage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_poktroll_proof_stage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_poktroll_proof_stage_proto_goTypes,
		DependencyIndexes: file_poktroll_proof_stage_proto_depIdxs,
		EnumInfos:         file_poktroll_proof_stage_proto_enumTypes,
	}.Build()
	File_poktroll_proof_stage_proto = out.File
	file_poktroll_proof_stage_proto_rawDesc = nil
	file_poktroll_proof_stage_proto_goTypes = nil
	file_poktroll_proof_stage_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: onboarding/onboarding_enums.proto

package onboarding

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OnboardingStage int32

const (
	OnboardingStage_ONBOARDING_STAGE_NEW_USER            OnboardingStage = 0
	OnboardingStage_ONBOARDING_STAGE_EMAIL_VERIFICATION  OnboardingStage = 1
	OnboardingStage_ONBOARDING_STAGE_PHONE_VERIFICATION  OnboardingStage = 2
	OnboardingStage_ONBOARDING_STAGE_PAN_VERIFICATION    OnboardingStage = 3
	OnboardingStage_ONBOARDING_STAGE_AADHAR_VERIFICATION OnboardingStage = 4
	OnboardingStage_ONBOARDING_STAGE_LIVENESS            OnboardingStage = 5
	OnboardingStage_ONBOARDING_STAGE_EXTRA_DETAILS       OnboardingStage = 6
	OnboardingStage_ONBOARDING_STAGE_HOME_STRETCH        OnboardingStage = 7
	OnboardingStage_ONBOARDING_COMPLETED                 OnboardingStage = 8
)

// Enum value maps for OnboardingStage.
var (
	OnboardingStage_name = map[int32]string{
		0: "ONBOARDING_STAGE_NEW_USER",
		1: "ONBOARDING_STAGE_EMAIL_VERIFICATION",
		2: "ONBOARDING_STAGE_PHONE_VERIFICATION",
		3: "ONBOARDING_STAGE_PAN_VERIFICATION",
		4: "ONBOARDING_STAGE_AADHAR_VERIFICATION",
		5: "ONBOARDING_STAGE_LIVENESS",
		6: "ONBOARDING_STAGE_EXTRA_DETAILS",
		7: "ONBOARDING_STAGE_HOME_STRETCH",
		8: "ONBOARDING_COMPLETED",
	}
	OnboardingStage_value = map[string]int32{
		"ONBOARDING_STAGE_NEW_USER":            0,
		"ONBOARDING_STAGE_EMAIL_VERIFICATION":  1,
		"ONBOARDING_STAGE_PHONE_VERIFICATION":  2,
		"ONBOARDING_STAGE_PAN_VERIFICATION":    3,
		"ONBOARDING_STAGE_AADHAR_VERIFICATION": 4,
		"ONBOARDING_STAGE_LIVENESS":            5,
		"ONBOARDING_STAGE_EXTRA_DETAILS":       6,
		"ONBOARDING_STAGE_HOME_STRETCH":        7,
		"ONBOARDING_COMPLETED":                 8,
	}
)

func (x OnboardingStage) Enum() *OnboardingStage {
	p := new(OnboardingStage)
	*p = x
	return p
}

func (x OnboardingStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OnboardingStage) Descriptor() protoreflect.EnumDescriptor {
	return file_onboarding_onboarding_enums_proto_enumTypes[0].Descriptor()
}

func (OnboardingStage) Type() protoreflect.EnumType {
	return &file_onboarding_onboarding_enums_proto_enumTypes[0]
}

func (x OnboardingStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OnboardingStage.Descriptor instead.
func (OnboardingStage) EnumDescriptor() ([]byte, []int) {
	return file_onboarding_onboarding_enums_proto_rawDescGZIP(), []int{0}
}

var File_onboarding_onboarding_enums_proto protoreflect.FileDescriptor

var file_onboarding_onboarding_enums_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2a,
	0xd3, 0x02, 0x0a, 0x0f, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x4f, 0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x56, 0x45, 0x52,
	0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x4f,
	0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f,
	0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21, 0x4f, 0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x41, 0x4e, 0x5f, 0x56, 0x45, 0x52,
	0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x28, 0x0a, 0x24, 0x4f,
	0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f,
	0x41, 0x41, 0x44, 0x48, 0x41, 0x52, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x4c, 0x49, 0x56, 0x45, 0x4e, 0x45,
	0x53, 0x53, 0x10, 0x05, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x45, 0x58, 0x54, 0x52, 0x41, 0x5f, 0x44,
	0x45, 0x54, 0x41, 0x49, 0x4c, 0x53, 0x10, 0x06, 0x12, 0x21, 0x0a, 0x1d, 0x4f, 0x4e, 0x42, 0x4f,
	0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x48, 0x4f, 0x4d,
	0x45, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x54, 0x43, 0x48, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x4f,
	0x4e, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x44, 0x10, 0x08, 0x42, 0x15, 0x5a, 0x13, 0x52, 0x69, 0x6e, 0x6e, 0x65, 0x67, 0x61,
	0x6e, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_onboarding_onboarding_enums_proto_rawDescOnce sync.Once
	file_onboarding_onboarding_enums_proto_rawDescData = file_onboarding_onboarding_enums_proto_rawDesc
)

func file_onboarding_onboarding_enums_proto_rawDescGZIP() []byte {
	file_onboarding_onboarding_enums_proto_rawDescOnce.Do(func() {
		file_onboarding_onboarding_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_onboarding_onboarding_enums_proto_rawDescData)
	})
	return file_onboarding_onboarding_enums_proto_rawDescData
}

var file_onboarding_onboarding_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_onboarding_onboarding_enums_proto_goTypes = []interface{}{
	(OnboardingStage)(0), // 0: onboarding.OnboardingStage
}
var file_onboarding_onboarding_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_onboarding_onboarding_enums_proto_init() }
func file_onboarding_onboarding_enums_proto_init() {
	if File_onboarding_onboarding_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_onboarding_onboarding_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_onboarding_onboarding_enums_proto_goTypes,
		DependencyIndexes: file_onboarding_onboarding_enums_proto_depIdxs,
		EnumInfos:         file_onboarding_onboarding_enums_proto_enumTypes,
	}.Build()
	File_onboarding_onboarding_enums_proto = out.File
	file_onboarding_onboarding_enums_proto_rawDesc = nil
	file_onboarding_onboarding_enums_proto_goTypes = nil
	file_onboarding_onboarding_enums_proto_depIdxs = nil
}

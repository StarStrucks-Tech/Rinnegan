// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: onboarding/onboarding_service.proto

package onboarding

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_onboarding_onboarding_service_proto protoreflect.FileDescriptor

var file_onboarding_onboarding_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x1a, 0x24, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x7a,
	0x0a, 0x11, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x6f, 0x6e, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x52, 0x69,
	0x6e, 0x6e, 0x65, 0x67, 0x61, 0x6e, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_onboarding_onboarding_service_proto_goTypes = []interface{}{
	(*UpdateCurrentStageRequest)(nil),  // 0: onboarding.UpdateCurrentStageRequest
	(*UpdateCurrentStageResponse)(nil), // 1: onboarding.UpdateCurrentStageResponse
}
var file_onboarding_onboarding_service_proto_depIdxs = []int32{
	0, // 0: onboarding.OnboardingService.UpdateCurrentStage:input_type -> onboarding.UpdateCurrentStageRequest
	1, // 1: onboarding.OnboardingService.UpdateCurrentStage:output_type -> onboarding.UpdateCurrentStageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_onboarding_onboarding_service_proto_init() }
func file_onboarding_onboarding_service_proto_init() {
	if File_onboarding_onboarding_service_proto != nil {
		return
	}
	file_onboarding_onboarding_requests_proto_init()
	file_onboarding_onboarding_responses_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_onboarding_onboarding_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_onboarding_onboarding_service_proto_goTypes,
		DependencyIndexes: file_onboarding_onboarding_service_proto_depIdxs,
	}.Build()
	File_onboarding_onboarding_service_proto = out.File
	file_onboarding_onboarding_service_proto_rawDesc = nil
	file_onboarding_onboarding_service_proto_goTypes = nil
	file_onboarding_onboarding_service_proto_depIdxs = nil
}

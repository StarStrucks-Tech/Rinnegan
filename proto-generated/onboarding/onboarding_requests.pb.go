// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: onboarding/onboarding_requests.proto

package onboarding

import (
	core "Rinnegan/proto-generated/core"
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

type UpdateCurrentStageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqHeader *core.RequestHeader `protobuf:"bytes,1,opt,name=req_header,json=reqHeader,proto3" json:"req_header,omitempty"`
	ActorId   string              `protobuf:"bytes,2,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
}

func (x *UpdateCurrentStageRequest) Reset() {
	*x = UpdateCurrentStageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onboarding_onboarding_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCurrentStageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentStageRequest) ProtoMessage() {}

func (x *UpdateCurrentStageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onboarding_onboarding_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentStageRequest.ProtoReflect.Descriptor instead.
func (*UpdateCurrentStageRequest) Descriptor() ([]byte, []int) {
	return file_onboarding_onboarding_requests_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCurrentStageRequest) GetReqHeader() *core.RequestHeader {
	if x != nil {
		return x.ReqHeader
	}
	return nil
}

func (x *UpdateCurrentStageRequest) GetActorId() string {
	if x != nil {
		return x.ActorId
	}
	return ""
}

var File_onboarding_onboarding_requests_proto protoreflect.FileDescriptor

var file_onboarding_onboarding_requests_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x1a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x42, 0x15, 0x5a, 0x13, 0x52, 0x69, 0x6e, 0x6e, 0x65, 0x67, 0x61, 0x6e, 0x2f, 0x6f,
	0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_onboarding_onboarding_requests_proto_rawDescOnce sync.Once
	file_onboarding_onboarding_requests_proto_rawDescData = file_onboarding_onboarding_requests_proto_rawDesc
)

func file_onboarding_onboarding_requests_proto_rawDescGZIP() []byte {
	file_onboarding_onboarding_requests_proto_rawDescOnce.Do(func() {
		file_onboarding_onboarding_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_onboarding_onboarding_requests_proto_rawDescData)
	})
	return file_onboarding_onboarding_requests_proto_rawDescData
}

var file_onboarding_onboarding_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_onboarding_onboarding_requests_proto_goTypes = []interface{}{
	(*UpdateCurrentStageRequest)(nil), // 0: onboarding.UpdateCurrentStageRequest
	(*core.RequestHeader)(nil),        // 1: core.RequestHeader
}
var file_onboarding_onboarding_requests_proto_depIdxs = []int32{
	1, // 0: onboarding.UpdateCurrentStageRequest.req_header:type_name -> core.RequestHeader
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_onboarding_onboarding_requests_proto_init() }
func file_onboarding_onboarding_requests_proto_init() {
	if File_onboarding_onboarding_requests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_onboarding_onboarding_requests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCurrentStageRequest); i {
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
			RawDescriptor: file_onboarding_onboarding_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_onboarding_onboarding_requests_proto_goTypes,
		DependencyIndexes: file_onboarding_onboarding_requests_proto_depIdxs,
		MessageInfos:      file_onboarding_onboarding_requests_proto_msgTypes,
	}.Build()
	File_onboarding_onboarding_requests_proto = out.File
	file_onboarding_onboarding_requests_proto_rawDesc = nil
	file_onboarding_onboarding_requests_proto_goTypes = nil
	file_onboarding_onboarding_requests_proto_depIdxs = nil
}

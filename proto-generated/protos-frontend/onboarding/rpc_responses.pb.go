// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: protos-frontend/onboarding/rpc_responses.proto

package onboarding

import (
	core "Rinnegan/proto-generated/core"
	deeplink "Rinnegan/proto-generated/protos-frontend/deeplink"
	enums "Rinnegan/proto-generated/protos-frontend/enums"
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

type TriggerPhoneVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespHeader   *core.ResponseHeader           `protobuf:"bytes,1,opt,name=resp_header,json=respHeader,proto3" json:"resp_header,omitempty"`
	ResponseType enums.VerificationResponseType `protobuf:"varint,2,opt,name=response_type,json=responseType,proto3,enum=enums.VerificationResponseType" json:"response_type,omitempty"`
}

func (x *TriggerPhoneVerificationResponse) Reset() {
	*x = TriggerPhoneVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerPhoneVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerPhoneVerificationResponse) ProtoMessage() {}

func (x *TriggerPhoneVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerPhoneVerificationResponse.ProtoReflect.Descriptor instead.
func (*TriggerPhoneVerificationResponse) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_responses_proto_rawDescGZIP(), []int{0}
}

func (x *TriggerPhoneVerificationResponse) GetRespHeader() *core.ResponseHeader {
	if x != nil {
		return x.RespHeader
	}
	return nil
}

func (x *TriggerPhoneVerificationResponse) GetResponseType() enums.VerificationResponseType {
	if x != nil {
		return x.ResponseType
	}
	return enums.VerificationResponseType(0)
}

type TriggerPhoneOTPVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespHeader   *core.ResponseHeader           `protobuf:"bytes,1,opt,name=resp_header,json=respHeader,proto3" json:"resp_header,omitempty"`
	ResponseType enums.VerificationResponseType `protobuf:"varint,2,opt,name=response_type,json=responseType,proto3,enum=enums.VerificationResponseType" json:"response_type,omitempty"`
}

func (x *TriggerPhoneOTPVerificationResponse) Reset() {
	*x = TriggerPhoneOTPVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerPhoneOTPVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerPhoneOTPVerificationResponse) ProtoMessage() {}

func (x *TriggerPhoneOTPVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerPhoneOTPVerificationResponse.ProtoReflect.Descriptor instead.
func (*TriggerPhoneOTPVerificationResponse) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_responses_proto_rawDescGZIP(), []int{1}
}

func (x *TriggerPhoneOTPVerificationResponse) GetRespHeader() *core.ResponseHeader {
	if x != nil {
		return x.RespHeader
	}
	return nil
}

func (x *TriggerPhoneOTPVerificationResponse) GetResponseType() enums.VerificationResponseType {
	if x != nil {
		return x.ResponseType
	}
	return enums.VerificationResponseType(0)
}

type TriggerEmailVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespHeader   *core.ResponseHeader           `protobuf:"bytes,1,opt,name=resp_header,json=respHeader,proto3" json:"resp_header,omitempty"`
	ResponseType enums.VerificationResponseType `protobuf:"varint,2,opt,name=response_type,json=responseType,proto3,enum=enums.VerificationResponseType" json:"response_type,omitempty"`
}

func (x *TriggerEmailVerificationResponse) Reset() {
	*x = TriggerEmailVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerEmailVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerEmailVerificationResponse) ProtoMessage() {}

func (x *TriggerEmailVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerEmailVerificationResponse.ProtoReflect.Descriptor instead.
func (*TriggerEmailVerificationResponse) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_responses_proto_rawDescGZIP(), []int{2}
}

func (x *TriggerEmailVerificationResponse) GetRespHeader() *core.ResponseHeader {
	if x != nil {
		return x.RespHeader
	}
	return nil
}

func (x *TriggerEmailVerificationResponse) GetResponseType() enums.VerificationResponseType {
	if x != nil {
		return x.ResponseType
	}
	return enums.VerificationResponseType(0)
}

type TriggerPanVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespHeader   *core.ResponseHeader           `protobuf:"bytes,1,opt,name=resp_header,json=respHeader,proto3" json:"resp_header,omitempty"`
	ResponseType enums.VerificationResponseType `protobuf:"varint,2,opt,name=response_type,json=responseType,proto3,enum=enums.VerificationResponseType" json:"response_type,omitempty"`
	Otp          string                         `protobuf:"bytes,3,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *TriggerPanVerificationResponse) Reset() {
	*x = TriggerPanVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerPanVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerPanVerificationResponse) ProtoMessage() {}

func (x *TriggerPanVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerPanVerificationResponse.ProtoReflect.Descriptor instead.
func (*TriggerPanVerificationResponse) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_responses_proto_rawDescGZIP(), []int{3}
}

func (x *TriggerPanVerificationResponse) GetRespHeader() *core.ResponseHeader {
	if x != nil {
		return x.RespHeader
	}
	return nil
}

func (x *TriggerPanVerificationResponse) GetResponseType() enums.VerificationResponseType {
	if x != nil {
		return x.ResponseType
	}
	return enums.VerificationResponseType(0)
}

func (x *TriggerPanVerificationResponse) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type TriggerLivenessCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespHeader          *core.ResponseHeader      `protobuf:"bytes,1,opt,name=resp_header,json=respHeader,proto3" json:"resp_header,omitempty"`
	LivenessCheckStatus enums.LivenessCheckStatus `protobuf:"varint,2,opt,name=liveness_check_status,json=livenessCheckStatus,proto3,enum=enums.LivenessCheckStatus" json:"liveness_check_status,omitempty"`
}

func (x *TriggerLivenessCheckResponse) Reset() {
	*x = TriggerLivenessCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerLivenessCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerLivenessCheckResponse) ProtoMessage() {}

func (x *TriggerLivenessCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerLivenessCheckResponse.ProtoReflect.Descriptor instead.
func (*TriggerLivenessCheckResponse) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_responses_proto_rawDescGZIP(), []int{4}
}

func (x *TriggerLivenessCheckResponse) GetRespHeader() *core.ResponseHeader {
	if x != nil {
		return x.RespHeader
	}
	return nil
}

func (x *TriggerLivenessCheckResponse) GetLivenessCheckStatus() enums.LivenessCheckStatus {
	if x != nil {
		return x.LivenessCheckStatus
	}
	return enums.LivenessCheckStatus(0)
}

type GetCurrentOnboardingStageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespHeader   *core.ResponseHeader  `protobuf:"bytes,1,opt,name=resp_header,json=respHeader,proto3" json:"resp_header,omitempty"`
	CurrentStage enums.OnboardingStage `protobuf:"varint,2,opt,name=current_stage,json=currentStage,proto3,enum=enums.OnboardingStage" json:"current_stage,omitempty"`
	NextAction   *deeplink.Deeplink    `protobuf:"bytes,3,opt,name=next_action,json=nextAction,proto3" json:"next_action,omitempty"`
}

func (x *GetCurrentOnboardingStageResponse) Reset() {
	*x = GetCurrentOnboardingStageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentOnboardingStageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentOnboardingStageResponse) ProtoMessage() {}

func (x *GetCurrentOnboardingStageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentOnboardingStageResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentOnboardingStageResponse) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_responses_proto_rawDescGZIP(), []int{5}
}

func (x *GetCurrentOnboardingStageResponse) GetRespHeader() *core.ResponseHeader {
	if x != nil {
		return x.RespHeader
	}
	return nil
}

func (x *GetCurrentOnboardingStageResponse) GetCurrentStage() enums.OnboardingStage {
	if x != nil {
		return x.CurrentStage
	}
	return enums.OnboardingStage(0)
}

func (x *GetCurrentOnboardingStageResponse) GetNextAction() *deeplink.Deeplink {
	if x != nil {
		return x.NextAction
	}
	return nil
}

var File_protos_frontend_onboarding_rpc_responses_proto protoreflect.FileDescriptor

var file_protos_frontend_onboarding_rpc_responses_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x2a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x64, 0x65, 0x65,
	0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x20, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x23, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4f, 0x54, 0x50, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x20,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xaf, 0x01,
	0x0a, 0x1e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x6e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6f, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22,
	0xa5, 0x01, 0x0a, 0x1c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x15, 0x6c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4c,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x13, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x33, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x65, 0x65, 0x70, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x35, 0x5a, 0x33, 0x52, 0x69, 0x6e, 0x6e, 0x65, 0x67,
	0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_frontend_onboarding_rpc_responses_proto_rawDescOnce sync.Once
	file_protos_frontend_onboarding_rpc_responses_proto_rawDescData = file_protos_frontend_onboarding_rpc_responses_proto_rawDesc
)

func file_protos_frontend_onboarding_rpc_responses_proto_rawDescGZIP() []byte {
	file_protos_frontend_onboarding_rpc_responses_proto_rawDescOnce.Do(func() {
		file_protos_frontend_onboarding_rpc_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_frontend_onboarding_rpc_responses_proto_rawDescData)
	})
	return file_protos_frontend_onboarding_rpc_responses_proto_rawDescData
}

var file_protos_frontend_onboarding_rpc_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_frontend_onboarding_rpc_responses_proto_goTypes = []interface{}{
	(*TriggerPhoneVerificationResponse)(nil),    // 0: onboarding.TriggerPhoneVerificationResponse
	(*TriggerPhoneOTPVerificationResponse)(nil), // 1: onboarding.TriggerPhoneOTPVerificationResponse
	(*TriggerEmailVerificationResponse)(nil),    // 2: onboarding.TriggerEmailVerificationResponse
	(*TriggerPanVerificationResponse)(nil),      // 3: onboarding.TriggerPanVerificationResponse
	(*TriggerLivenessCheckResponse)(nil),        // 4: onboarding.TriggerLivenessCheckResponse
	(*GetCurrentOnboardingStageResponse)(nil),   // 5: onboarding.GetCurrentOnboardingStageResponse
	(*core.ResponseHeader)(nil),                 // 6: core.ResponseHeader
	(enums.VerificationResponseType)(0),         // 7: enums.VerificationResponseType
	(enums.LivenessCheckStatus)(0),              // 8: enums.LivenessCheckStatus
	(enums.OnboardingStage)(0),                  // 9: enums.OnboardingStage
	(*deeplink.Deeplink)(nil),                   // 10: deeplink.Deeplink
}
var file_protos_frontend_onboarding_rpc_responses_proto_depIdxs = []int32{
	6,  // 0: onboarding.TriggerPhoneVerificationResponse.resp_header:type_name -> core.ResponseHeader
	7,  // 1: onboarding.TriggerPhoneVerificationResponse.response_type:type_name -> enums.VerificationResponseType
	6,  // 2: onboarding.TriggerPhoneOTPVerificationResponse.resp_header:type_name -> core.ResponseHeader
	7,  // 3: onboarding.TriggerPhoneOTPVerificationResponse.response_type:type_name -> enums.VerificationResponseType
	6,  // 4: onboarding.TriggerEmailVerificationResponse.resp_header:type_name -> core.ResponseHeader
	7,  // 5: onboarding.TriggerEmailVerificationResponse.response_type:type_name -> enums.VerificationResponseType
	6,  // 6: onboarding.TriggerPanVerificationResponse.resp_header:type_name -> core.ResponseHeader
	7,  // 7: onboarding.TriggerPanVerificationResponse.response_type:type_name -> enums.VerificationResponseType
	6,  // 8: onboarding.TriggerLivenessCheckResponse.resp_header:type_name -> core.ResponseHeader
	8,  // 9: onboarding.TriggerLivenessCheckResponse.liveness_check_status:type_name -> enums.LivenessCheckStatus
	6,  // 10: onboarding.GetCurrentOnboardingStageResponse.resp_header:type_name -> core.ResponseHeader
	9,  // 11: onboarding.GetCurrentOnboardingStageResponse.current_stage:type_name -> enums.OnboardingStage
	10, // 12: onboarding.GetCurrentOnboardingStageResponse.next_action:type_name -> deeplink.Deeplink
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_protos_frontend_onboarding_rpc_responses_proto_init() }
func file_protos_frontend_onboarding_rpc_responses_proto_init() {
	if File_protos_frontend_onboarding_rpc_responses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerPhoneVerificationResponse); i {
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
		file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerPhoneOTPVerificationResponse); i {
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
		file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerEmailVerificationResponse); i {
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
		file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerPanVerificationResponse); i {
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
		file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerLivenessCheckResponse); i {
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
		file_protos_frontend_onboarding_rpc_responses_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentOnboardingStageResponse); i {
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
			RawDescriptor: file_protos_frontend_onboarding_rpc_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_frontend_onboarding_rpc_responses_proto_goTypes,
		DependencyIndexes: file_protos_frontend_onboarding_rpc_responses_proto_depIdxs,
		MessageInfos:      file_protos_frontend_onboarding_rpc_responses_proto_msgTypes,
	}.Build()
	File_protos_frontend_onboarding_rpc_responses_proto = out.File
	file_protos_frontend_onboarding_rpc_responses_proto_rawDesc = nil
	file_protos_frontend_onboarding_rpc_responses_proto_goTypes = nil
	file_protos_frontend_onboarding_rpc_responses_proto_depIdxs = nil
}

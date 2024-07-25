// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: protos-frontend/onboarding/rpc_requests.proto

package onboarding

import (
	core "Rinnegan/proto-generated/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TriggerPhoneOTPVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqHeader *core.RequestHeader `protobuf:"bytes,1,opt,name=req_header,json=reqHeader,proto3" json:"req_header,omitempty"`
	Otp       string              `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *TriggerPhoneOTPVerificationRequest) Reset() {
	*x = TriggerPhoneOTPVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerPhoneOTPVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerPhoneOTPVerificationRequest) ProtoMessage() {}

func (x *TriggerPhoneOTPVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerPhoneOTPVerificationRequest.ProtoReflect.Descriptor instead.
func (*TriggerPhoneOTPVerificationRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_requests_proto_rawDescGZIP(), []int{0}
}

func (x *TriggerPhoneOTPVerificationRequest) GetReqHeader() *core.RequestHeader {
	if x != nil {
		return x.ReqHeader
	}
	return nil
}

func (x *TriggerPhoneOTPVerificationRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type TriggerEmailVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqHeader *core.RequestHeader `protobuf:"bytes,1,opt,name=req_header,json=reqHeader,proto3" json:"req_header,omitempty"`
	EmailId   string              `protobuf:"bytes,2,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	UserName  string              `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *TriggerEmailVerificationRequest) Reset() {
	*x = TriggerEmailVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerEmailVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerEmailVerificationRequest) ProtoMessage() {}

func (x *TriggerEmailVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerEmailVerificationRequest.ProtoReflect.Descriptor instead.
func (*TriggerEmailVerificationRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_requests_proto_rawDescGZIP(), []int{1}
}

func (x *TriggerEmailVerificationRequest) GetReqHeader() *core.RequestHeader {
	if x != nil {
		return x.ReqHeader
	}
	return nil
}

func (x *TriggerEmailVerificationRequest) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

func (x *TriggerEmailVerificationRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type TriggerPanVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqHeader   *core.RequestHeader    `protobuf:"bytes,1,opt,name=req_header,json=reqHeader,proto3" json:"req_header,omitempty"`
	PanNumber   string                 `protobuf:"bytes,2,opt,name=pan_number,json=panNumber,proto3" json:"pan_number,omitempty"`
	DateOfBirth *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
}

func (x *TriggerPanVerificationRequest) Reset() {
	*x = TriggerPanVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerPanVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerPanVerificationRequest) ProtoMessage() {}

func (x *TriggerPanVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerPanVerificationRequest.ProtoReflect.Descriptor instead.
func (*TriggerPanVerificationRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_requests_proto_rawDescGZIP(), []int{2}
}

func (x *TriggerPanVerificationRequest) GetReqHeader() *core.RequestHeader {
	if x != nil {
		return x.ReqHeader
	}
	return nil
}

func (x *TriggerPanVerificationRequest) GetPanNumber() string {
	if x != nil {
		return x.PanNumber
	}
	return ""
}

func (x *TriggerPanVerificationRequest) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

type TriggerLivenessCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqHeader *core.RequestHeader `protobuf:"bytes,1,opt,name=req_header,json=reqHeader,proto3" json:"req_header,omitempty"`
	OtpSpoken string              `protobuf:"bytes,2,opt,name=otp_spoken,json=otpSpoken,proto3" json:"otp_spoken,omitempty"`
}

func (x *TriggerLivenessCheckRequest) Reset() {
	*x = TriggerLivenessCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerLivenessCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerLivenessCheckRequest) ProtoMessage() {}

func (x *TriggerLivenessCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerLivenessCheckRequest.ProtoReflect.Descriptor instead.
func (*TriggerLivenessCheckRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_requests_proto_rawDescGZIP(), []int{3}
}

func (x *TriggerLivenessCheckRequest) GetReqHeader() *core.RequestHeader {
	if x != nil {
		return x.ReqHeader
	}
	return nil
}

func (x *TriggerLivenessCheckRequest) GetOtpSpoken() string {
	if x != nil {
		return x.OtpSpoken
	}
	return ""
}

// Define the Aadhar verification request message
type AadharVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqHeader         *core.RequestHeader `protobuf:"bytes,1,opt,name=req_header,json=reqHeader,proto3" json:"req_header,omitempty"`
	AadharNumber      string              `protobuf:"bytes,2,opt,name=aadhar_number,json=aadharNumber,proto3" json:"aadhar_number,omitempty"`
	Otp               string              `protobuf:"bytes,3,opt,name=otp,proto3" json:"otp,omitempty"`
	LinkedPhoneNumber string              `protobuf:"bytes,4,opt,name=linked_phone_number,json=linkedPhoneNumber,proto3" json:"linked_phone_number,omitempty"`
}

func (x *AadharVerificationRequest) Reset() {
	*x = AadharVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AadharVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AadharVerificationRequest) ProtoMessage() {}

func (x *AadharVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AadharVerificationRequest.ProtoReflect.Descriptor instead.
func (*AadharVerificationRequest) Descriptor() ([]byte, []int) {
	return file_protos_frontend_onboarding_rpc_requests_proto_rawDescGZIP(), []int{4}
}

func (x *AadharVerificationRequest) GetReqHeader() *core.RequestHeader {
	if x != nil {
		return x.ReqHeader
	}
	return nil
}

func (x *AadharVerificationRequest) GetAadharNumber() string {
	if x != nil {
		return x.AadharNumber
	}
	return ""
}

func (x *AadharVerificationRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *AadharVerificationRequest) GetLinkedPhoneNumber() string {
	if x != nil {
		return x.LinkedPhoneNumber
	}
	return ""
}

var File_protos_frontend_onboarding_rpc_requests_proto protoreflect.FileDescriptor

var file_protos_frontend_onboarding_rpc_requests_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x70, 0x63,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6a, 0x0a, 0x22, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4f, 0x54, 0x50, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x8d, 0x01, 0x0a,
	0x1f, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x71, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb2, 0x01, 0x0a,
	0x1d, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x50, 0x61, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x71, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x22, 0x70, 0x0a, 0x1b, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x76, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x71, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x74, 0x70, 0x5f, 0x73, 0x70, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x74, 0x70, 0x53, 0x70, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xb6, 0x01, 0x0a, 0x19, 0x41, 0x61, 0x64, 0x68, 0x61, 0x72, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x09, 0x72, 0x65, 0x71, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x61, 0x64, 0x68, 0x61, 0x72, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x61,
	0x64, 0x68, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x2e, 0x0a, 0x13,
	0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x35, 0x5a, 0x33,
	0x52, 0x69, 0x6e, 0x6e, 0x65, 0x67, 0x61, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_frontend_onboarding_rpc_requests_proto_rawDescOnce sync.Once
	file_protos_frontend_onboarding_rpc_requests_proto_rawDescData = file_protos_frontend_onboarding_rpc_requests_proto_rawDesc
)

func file_protos_frontend_onboarding_rpc_requests_proto_rawDescGZIP() []byte {
	file_protos_frontend_onboarding_rpc_requests_proto_rawDescOnce.Do(func() {
		file_protos_frontend_onboarding_rpc_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_frontend_onboarding_rpc_requests_proto_rawDescData)
	})
	return file_protos_frontend_onboarding_rpc_requests_proto_rawDescData
}

var file_protos_frontend_onboarding_rpc_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_frontend_onboarding_rpc_requests_proto_goTypes = []any{
	(*TriggerPhoneOTPVerificationRequest)(nil), // 0: onboarding.TriggerPhoneOTPVerificationRequest
	(*TriggerEmailVerificationRequest)(nil),    // 1: onboarding.TriggerEmailVerificationRequest
	(*TriggerPanVerificationRequest)(nil),      // 2: onboarding.TriggerPanVerificationRequest
	(*TriggerLivenessCheckRequest)(nil),        // 3: onboarding.TriggerLivenessCheckRequest
	(*AadharVerificationRequest)(nil),          // 4: onboarding.AadharVerificationRequest
	(*core.RequestHeader)(nil),                 // 5: core.RequestHeader
	(*timestamppb.Timestamp)(nil),              // 6: google.protobuf.Timestamp
}
var file_protos_frontend_onboarding_rpc_requests_proto_depIdxs = []int32{
	5, // 0: onboarding.TriggerPhoneOTPVerificationRequest.req_header:type_name -> core.RequestHeader
	5, // 1: onboarding.TriggerEmailVerificationRequest.req_header:type_name -> core.RequestHeader
	5, // 2: onboarding.TriggerPanVerificationRequest.req_header:type_name -> core.RequestHeader
	6, // 3: onboarding.TriggerPanVerificationRequest.date_of_birth:type_name -> google.protobuf.Timestamp
	5, // 4: onboarding.TriggerLivenessCheckRequest.req_header:type_name -> core.RequestHeader
	5, // 5: onboarding.AadharVerificationRequest.req_header:type_name -> core.RequestHeader
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protos_frontend_onboarding_rpc_requests_proto_init() }
func file_protos_frontend_onboarding_rpc_requests_proto_init() {
	if File_protos_frontend_onboarding_rpc_requests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TriggerPhoneOTPVerificationRequest); i {
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
		file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TriggerEmailVerificationRequest); i {
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
		file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TriggerPanVerificationRequest); i {
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
		file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TriggerLivenessCheckRequest); i {
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
		file_protos_frontend_onboarding_rpc_requests_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AadharVerificationRequest); i {
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
			RawDescriptor: file_protos_frontend_onboarding_rpc_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_frontend_onboarding_rpc_requests_proto_goTypes,
		DependencyIndexes: file_protos_frontend_onboarding_rpc_requests_proto_depIdxs,
		MessageInfos:      file_protos_frontend_onboarding_rpc_requests_proto_msgTypes,
	}.Build()
	File_protos_frontend_onboarding_rpc_requests_proto = out.File
	file_protos_frontend_onboarding_rpc_requests_proto_rawDesc = nil
	file_protos_frontend_onboarding_rpc_requests_proto_goTypes = nil
	file_protos_frontend_onboarding_rpc_requests_proto_depIdxs = nil
}

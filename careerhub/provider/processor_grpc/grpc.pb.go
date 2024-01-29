// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: careerhub/provider/processor_grpc/grpc.proto

package processor_grpc

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

type JobPostings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPostingIds []*JobPostingId `protobuf:"bytes,1,rep,name=jobPostingIds,proto3" json:"jobPostingIds,omitempty"`
}

func (x *JobPostings) Reset() {
	*x = JobPostings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobPostings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobPostings) ProtoMessage() {}

func (x *JobPostings) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobPostings.ProtoReflect.Descriptor instead.
func (*JobPostings) Descriptor() ([]byte, []int) {
	return file_careerhub_provider_processor_grpc_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *JobPostings) GetJobPostingIds() []*JobPostingId {
	if x != nil {
		return x.JobPostingIds
	}
	return nil
}

type JobPostingId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site      string `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	PostingId string `protobuf:"bytes,2,opt,name=postingId,proto3" json:"postingId,omitempty"`
}

func (x *JobPostingId) Reset() {
	*x = JobPostingId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobPostingId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobPostingId) ProtoMessage() {}

func (x *JobPostingId) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobPostingId.ProtoReflect.Descriptor instead.
func (*JobPostingId) Descriptor() ([]byte, []int) {
	return file_careerhub_provider_processor_grpc_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *JobPostingId) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *JobPostingId) GetPostingId() string {
	if x != nil {
		return x.PostingId
	}
	return ""
}

type JobPostingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobPostingId   *JobPostingId `protobuf:"bytes,1,opt,name=jobPostingId,proto3" json:"jobPostingId,omitempty"`
	CompanyId      string        `protobuf:"bytes,2,opt,name=companyId,proto3" json:"companyId,omitempty"`
	CompanyName    string        `protobuf:"bytes,3,opt,name=companyName,proto3" json:"companyName,omitempty"`
	JobCategory    []string      `protobuf:"bytes,4,rep,name=jobCategory,proto3" json:"jobCategory,omitempty"`
	MainContent    *MainContent  `protobuf:"bytes,5,opt,name=mainContent,proto3" json:"mainContent,omitempty"`
	RequiredSkill  []string      `protobuf:"bytes,6,rep,name=requiredSkill,proto3" json:"requiredSkill,omitempty"`
	Tags           []string      `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	RequiredCareer *Career       `protobuf:"bytes,8,opt,name=requiredCareer,proto3" json:"requiredCareer,omitempty"`
	PublishedAt    *int64        `protobuf:"varint,9,opt,name=publishedAt,proto3,oneof" json:"publishedAt,omitempty"`
	ClosedAt       *int64        `protobuf:"varint,10,opt,name=closedAt,proto3,oneof" json:"closedAt,omitempty"`
	Address        []string      `protobuf:"bytes,11,rep,name=address,proto3" json:"address,omitempty"`
	CreatedAt      int64         `protobuf:"varint,12,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *JobPostingInfo) Reset() {
	*x = JobPostingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobPostingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobPostingInfo) ProtoMessage() {}

func (x *JobPostingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobPostingInfo.ProtoReflect.Descriptor instead.
func (*JobPostingInfo) Descriptor() ([]byte, []int) {
	return file_careerhub_provider_processor_grpc_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *JobPostingInfo) GetJobPostingId() *JobPostingId {
	if x != nil {
		return x.JobPostingId
	}
	return nil
}

func (x *JobPostingInfo) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *JobPostingInfo) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *JobPostingInfo) GetJobCategory() []string {
	if x != nil {
		return x.JobCategory
	}
	return nil
}

func (x *JobPostingInfo) GetMainContent() *MainContent {
	if x != nil {
		return x.MainContent
	}
	return nil
}

func (x *JobPostingInfo) GetRequiredSkill() []string {
	if x != nil {
		return x.RequiredSkill
	}
	return nil
}

func (x *JobPostingInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *JobPostingInfo) GetRequiredCareer() *Career {
	if x != nil {
		return x.RequiredCareer
	}
	return nil
}

func (x *JobPostingInfo) GetPublishedAt() int64 {
	if x != nil && x.PublishedAt != nil {
		return *x.PublishedAt
	}
	return 0
}

func (x *JobPostingInfo) GetClosedAt() int64 {
	if x != nil && x.ClosedAt != nil {
		return *x.ClosedAt
	}
	return 0
}

func (x *JobPostingInfo) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *JobPostingInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type MainContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostUrl        string  `protobuf:"bytes,1,opt,name=postUrl,proto3" json:"postUrl,omitempty"`
	Title          string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Intro          string  `protobuf:"bytes,3,opt,name=intro,proto3" json:"intro,omitempty"`
	MainTask       string  `protobuf:"bytes,4,opt,name=mainTask,proto3" json:"mainTask,omitempty"`
	Qualifications string  `protobuf:"bytes,5,opt,name=qualifications,proto3" json:"qualifications,omitempty"`
	Preferred      string  `protobuf:"bytes,6,opt,name=preferred,proto3" json:"preferred,omitempty"`
	Benefits       string  `protobuf:"bytes,7,opt,name=benefits,proto3" json:"benefits,omitempty"`
	RecruitProcess *string `protobuf:"bytes,8,opt,name=recruitProcess,proto3,oneof" json:"recruitProcess,omitempty"`
}

func (x *MainContent) Reset() {
	*x = MainContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainContent) ProtoMessage() {}

func (x *MainContent) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainContent.ProtoReflect.Descriptor instead.
func (*MainContent) Descriptor() ([]byte, []int) {
	return file_careerhub_provider_processor_grpc_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *MainContent) GetPostUrl() string {
	if x != nil {
		return x.PostUrl
	}
	return ""
}

func (x *MainContent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MainContent) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *MainContent) GetMainTask() string {
	if x != nil {
		return x.MainTask
	}
	return ""
}

func (x *MainContent) GetQualifications() string {
	if x != nil {
		return x.Qualifications
	}
	return ""
}

func (x *MainContent) GetPreferred() string {
	if x != nil {
		return x.Preferred
	}
	return ""
}

func (x *MainContent) GetBenefits() string {
	if x != nil {
		return x.Benefits
	}
	return ""
}

func (x *MainContent) GetRecruitProcess() string {
	if x != nil && x.RecruitProcess != nil {
		return *x.RecruitProcess
	}
	return ""
}

type Career struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min *int32 `protobuf:"varint,1,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max *int32 `protobuf:"varint,2,opt,name=max,proto3,oneof" json:"max,omitempty"`
}

func (x *Career) Reset() {
	*x = Career{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Career) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Career) ProtoMessage() {}

func (x *Career) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Career.ProtoReflect.Descriptor instead.
func (*Career) Descriptor() ([]byte, []int) {
	return file_careerhub_provider_processor_grpc_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *Career) GetMin() int32 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *Career) GetMax() int32 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site          string   `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	CompanyId     string   `protobuf:"bytes,2,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Name          string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CompanyUrl    *string  `protobuf:"bytes,4,opt,name=companyUrl,proto3,oneof" json:"companyUrl,omitempty"`
	CompanyImages []string `protobuf:"bytes,5,rep,name=companyImages,proto3" json:"companyImages,omitempty"`
	Description   string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	CompanyLogo   string   `protobuf:"bytes,7,opt,name=companyLogo,proto3" json:"companyLogo,omitempty"`
	CreatedAt     int64    `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_careerhub_provider_processor_grpc_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *Company) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *Company) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetCompanyUrl() string {
	if x != nil && x.CompanyUrl != nil {
		return *x.CompanyUrl
	}
	return ""
}

func (x *Company) GetCompanyImages() []string {
	if x != nil {
		return x.CompanyImages
	}
	return nil
}

func (x *Company) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Company) GetCompanyLogo() string {
	if x != nil {
		return x.CompanyLogo
	}
	return ""
}

func (x *Company) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type BoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BoolResponse) Reset() {
	*x = BoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolResponse) ProtoMessage() {}

func (x *BoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolResponse.ProtoReflect.Descriptor instead.
func (*BoolResponse) Descriptor() ([]byte, []int) {
	return file_careerhub_provider_processor_grpc_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *BoolResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_careerhub_provider_processor_grpc_grpc_proto protoreflect.FileDescriptor

var file_careerhub_provider_processor_grpc_grpc_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x65, 0x0a, 0x0b, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x56, 0x0a, 0x0d, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65,
	0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f,
	0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x0d, 0x6a, 0x6f, 0x62, 0x50,
	0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x22, 0x40, 0x0a, 0x0c, 0x4a, 0x6f, 0x62,
	0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0xc6, 0x04, 0x0a, 0x0e,
	0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x54,
	0x0a, 0x0c, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x0c, 0x6a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x51, 0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x6d, 0x61,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x52, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x43,
	0x61, 0x72, 0x65, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0b,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x01, 0x52, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x02, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61,
	0x69, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61,
	0x69, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x72,
	0x75, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x46, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x65,
	0x65, 0x72, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88, 0x01, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78,
	0x22, 0x8b, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x55, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x55, 0x72, 0x6c, 0x22, 0x28,
	0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xf8, 0x02, 0x0a, 0x0d, 0x44, 0x61, 0x74,
	0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x75, 0x0a, 0x10, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2f,
	0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a,
	0x30, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x7e, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4a, 0x6f, 0x62,
	0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x2e, 0x63, 0x61,
	0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x30, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x70, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0x2b, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x1a, 0x30, 0x2e, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_careerhub_provider_processor_grpc_grpc_proto_rawDescOnce sync.Once
	file_careerhub_provider_processor_grpc_grpc_proto_rawDescData = file_careerhub_provider_processor_grpc_grpc_proto_rawDesc
)

func file_careerhub_provider_processor_grpc_grpc_proto_rawDescGZIP() []byte {
	file_careerhub_provider_processor_grpc_grpc_proto_rawDescOnce.Do(func() {
		file_careerhub_provider_processor_grpc_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_careerhub_provider_processor_grpc_grpc_proto_rawDescData)
	})
	return file_careerhub_provider_processor_grpc_grpc_proto_rawDescData
}

var file_careerhub_provider_processor_grpc_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_careerhub_provider_processor_grpc_grpc_proto_goTypes = []interface{}{
	(*JobPostings)(nil),    // 0: careerhub.processor.processor_grpc.JobPostings
	(*JobPostingId)(nil),   // 1: careerhub.processor.processor_grpc.JobPostingId
	(*JobPostingInfo)(nil), // 2: careerhub.processor.processor_grpc.JobPostingInfo
	(*MainContent)(nil),    // 3: careerhub.processor.processor_grpc.MainContent
	(*Career)(nil),         // 4: careerhub.processor.processor_grpc.Career
	(*Company)(nil),        // 5: careerhub.processor.processor_grpc.Company
	(*BoolResponse)(nil),   // 6: careerhub.processor.processor_grpc.BoolResponse
}
var file_careerhub_provider_processor_grpc_grpc_proto_depIdxs = []int32{
	1, // 0: careerhub.processor.processor_grpc.JobPostings.jobPostingIds:type_name -> careerhub.processor.processor_grpc.JobPostingId
	1, // 1: careerhub.processor.processor_grpc.JobPostingInfo.jobPostingId:type_name -> careerhub.processor.processor_grpc.JobPostingId
	3, // 2: careerhub.processor.processor_grpc.JobPostingInfo.mainContent:type_name -> careerhub.processor.processor_grpc.MainContent
	4, // 3: careerhub.processor.processor_grpc.JobPostingInfo.requiredCareer:type_name -> careerhub.processor.processor_grpc.Career
	0, // 4: careerhub.processor.processor_grpc.DataProcessor.CloseJobPostings:input_type -> careerhub.processor.processor_grpc.JobPostings
	2, // 5: careerhub.processor.processor_grpc.DataProcessor.RegisterJobPostingInfo:input_type -> careerhub.processor.processor_grpc.JobPostingInfo
	5, // 6: careerhub.processor.processor_grpc.DataProcessor.RegisterCompany:input_type -> careerhub.processor.processor_grpc.Company
	6, // 7: careerhub.processor.processor_grpc.DataProcessor.CloseJobPostings:output_type -> careerhub.processor.processor_grpc.BoolResponse
	6, // 8: careerhub.processor.processor_grpc.DataProcessor.RegisterJobPostingInfo:output_type -> careerhub.processor.processor_grpc.BoolResponse
	6, // 9: careerhub.processor.processor_grpc.DataProcessor.RegisterCompany:output_type -> careerhub.processor.processor_grpc.BoolResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_careerhub_provider_processor_grpc_grpc_proto_init() }
func file_careerhub_provider_processor_grpc_grpc_proto_init() {
	if File_careerhub_provider_processor_grpc_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobPostings); i {
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
		file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobPostingId); i {
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
		file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobPostingInfo); i {
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
		file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainContent); i {
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
		file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Career); i {
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
		file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
		file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolResponse); i {
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
	file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_careerhub_provider_processor_grpc_grpc_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_careerhub_provider_processor_grpc_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_careerhub_provider_processor_grpc_grpc_proto_goTypes,
		DependencyIndexes: file_careerhub_provider_processor_grpc_grpc_proto_depIdxs,
		MessageInfos:      file_careerhub_provider_processor_grpc_grpc_proto_msgTypes,
	}.Build()
	File_careerhub_provider_processor_grpc_grpc_proto = out.File
	file_careerhub_provider_processor_grpc_grpc_proto_rawDesc = nil
	file_careerhub_provider_processor_grpc_grpc_proto_goTypes = nil
	file_careerhub_provider_processor_grpc_grpc_proto_depIdxs = nil
}

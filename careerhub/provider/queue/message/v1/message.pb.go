// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: careerhub/provider/queue/message/v1/message.proto

package message_v1

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

type JobPostingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site           string       `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
	PostingId      string       `protobuf:"bytes,2,opt,name=postingId,proto3" json:"postingId,omitempty"`
	Company        *Company     `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	JobCategory    []string     `protobuf:"bytes,4,rep,name=jobCategory,proto3" json:"jobCategory,omitempty"`
	MainContent    *MainContent `protobuf:"bytes,5,opt,name=mainContent,proto3" json:"mainContent,omitempty"`
	RequiredSkill  []string     `protobuf:"bytes,6,rep,name=requiredSkill,proto3" json:"requiredSkill,omitempty"`
	Tags           []string     `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	RequiredCareer *Career      `protobuf:"bytes,8,opt,name=requiredCareer,proto3" json:"requiredCareer,omitempty"`
	PublishedAt    *int64       `protobuf:"varint,9,opt,name=publishedAt,proto3,oneof" json:"publishedAt,omitempty"`
	ClosedAt       *int64       `protobuf:"varint,10,opt,name=closedAt,proto3,oneof" json:"closedAt,omitempty"`
	Address        []string     `protobuf:"bytes,11,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *JobPostingInfo) Reset() {
	*x = JobPostingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_queue_message_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobPostingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobPostingInfo) ProtoMessage() {}

func (x *JobPostingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_queue_message_v1_message_proto_msgTypes[0]
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
	return file_careerhub_provider_queue_message_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *JobPostingInfo) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *JobPostingInfo) GetPostingId() string {
	if x != nil {
		return x.PostingId
	}
	return ""
}

func (x *JobPostingInfo) GetCompany() *Company {
	if x != nil {
		return x.Company
	}
	return nil
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
		mi := &file_careerhub_provider_queue_message_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainContent) ProtoMessage() {}

func (x *MainContent) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_queue_message_v1_message_proto_msgTypes[1]
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
	return file_careerhub_provider_queue_message_v1_message_proto_rawDescGZIP(), []int{1}
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

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CompanyUrl    *string  `protobuf:"bytes,2,opt,name=companyUrl,proto3,oneof" json:"companyUrl,omitempty"`
	CompanyImages []string `protobuf:"bytes,3,rep,name=companyImages,proto3" json:"companyImages,omitempty"`
	Description   string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CompanyLogo   string   `protobuf:"bytes,5,opt,name=companyLogo,proto3" json:"companyLogo,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_careerhub_provider_queue_message_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_queue_message_v1_message_proto_msgTypes[2]
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
	return file_careerhub_provider_queue_message_v1_message_proto_rawDescGZIP(), []int{2}
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
		mi := &file_careerhub_provider_queue_message_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Career) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Career) ProtoMessage() {}

func (x *Career) ProtoReflect() protoreflect.Message {
	mi := &file_careerhub_provider_queue_message_v1_message_proto_msgTypes[3]
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
	return file_careerhub_provider_queue_message_v1_message_proto_rawDescGZIP(), []int{3}
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

var File_careerhub_provider_queue_message_v1_message_proto protoreflect.FileDescriptor

var file_careerhub_provider_queue_message_v1_message_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x22,
	0xc3, 0x03, 0x0a, 0x0e, 0x4a, 0x6f, 0x62, 0x50, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x43, 0x61, 0x72, 0x65, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0b, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a,
	0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x01, 0x52, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x02, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x55, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x61, 0x69, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x61, 0x69, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x0e, 0x72, 0x65, 0x63,
	0x72, 0x75, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x72, 0x65, 0x63, 0x72, 0x75,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x55, 0x72, 0x6c, 0x22, 0x46, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x65, 0x65,
	0x72, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88, 0x01, 0x01, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x61, 0x78, 0x42,
	0x3c, 0x5a, 0x3a, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x68, 0x75, 0x62, 0x2d, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72,
	0x68, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_careerhub_provider_queue_message_v1_message_proto_rawDescOnce sync.Once
	file_careerhub_provider_queue_message_v1_message_proto_rawDescData = file_careerhub_provider_queue_message_v1_message_proto_rawDesc
)

func file_careerhub_provider_queue_message_v1_message_proto_rawDescGZIP() []byte {
	file_careerhub_provider_queue_message_v1_message_proto_rawDescOnce.Do(func() {
		file_careerhub_provider_queue_message_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_careerhub_provider_queue_message_v1_message_proto_rawDescData)
	})
	return file_careerhub_provider_queue_message_v1_message_proto_rawDescData
}

var file_careerhub_provider_queue_message_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_careerhub_provider_queue_message_v1_message_proto_goTypes = []interface{}{
	(*JobPostingInfo)(nil), // 0: message.v1.JobPostingInfo
	(*MainContent)(nil),    // 1: message.v1.MainContent
	(*Company)(nil),        // 2: message.v1.Company
	(*Career)(nil),         // 3: message.v1.Career
}
var file_careerhub_provider_queue_message_v1_message_proto_depIdxs = []int32{
	2, // 0: message.v1.JobPostingInfo.company:type_name -> message.v1.Company
	1, // 1: message.v1.JobPostingInfo.mainContent:type_name -> message.v1.MainContent
	3, // 2: message.v1.JobPostingInfo.requiredCareer:type_name -> message.v1.Career
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_careerhub_provider_queue_message_v1_message_proto_init() }
func file_careerhub_provider_queue_message_v1_message_proto_init() {
	if File_careerhub_provider_queue_message_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_careerhub_provider_queue_message_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_careerhub_provider_queue_message_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_careerhub_provider_queue_message_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_careerhub_provider_queue_message_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_careerhub_provider_queue_message_v1_message_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_careerhub_provider_queue_message_v1_message_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_careerhub_provider_queue_message_v1_message_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_careerhub_provider_queue_message_v1_message_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_careerhub_provider_queue_message_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_careerhub_provider_queue_message_v1_message_proto_goTypes,
		DependencyIndexes: file_careerhub_provider_queue_message_v1_message_proto_depIdxs,
		MessageInfos:      file_careerhub_provider_queue_message_v1_message_proto_msgTypes,
	}.Build()
	File_careerhub_provider_queue_message_v1_message_proto = out.File
	file_careerhub_provider_queue_message_v1_message_proto_rawDesc = nil
	file_careerhub_provider_queue_message_v1_message_proto_goTypes = nil
	file_careerhub_provider_queue_message_v1_message_proto_depIdxs = nil
}

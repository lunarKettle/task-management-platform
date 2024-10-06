// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: project_management.proto

package proto

import (
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

type ProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ProjectRequest) Reset() {
	*x = ProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectRequest) ProtoMessage() {}

func (x *ProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectRequest.ProtoReflect.Descriptor instead.
func (*ProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_management_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectRequest) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type ProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId          uint32                 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ProjectName        string                 `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	ProjectDescription string                 `protobuf:"bytes,3,opt,name=project_description,json=projectDescription,proto3" json:"project_description,omitempty"`
	StartDate          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	PlannedEndDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=planned_end_date,json=plannedEndDate,proto3" json:"planned_end_date,omitempty"`
	ActualEndDate      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=actual_end_date,json=actualEndDate,proto3" json:"actual_end_date,omitempty"`
	Status             string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Priority           uint32                 `protobuf:"varint,8,opt,name=priority,proto3" json:"priority,omitempty"`
	ManagerId          uint32                 `protobuf:"varint,9,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	Budget             float64                `protobuf:"fixed64,10,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *ProjectResponse) Reset() {
	*x = ProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectResponse) ProtoMessage() {}

func (x *ProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectResponse.ProtoReflect.Descriptor instead.
func (*ProjectResponse) Descriptor() ([]byte, []int) {
	return file_project_management_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectResponse) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ProjectResponse) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ProjectResponse) GetProjectDescription() string {
	if x != nil {
		return x.ProjectDescription
	}
	return ""
}

func (x *ProjectResponse) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *ProjectResponse) GetPlannedEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PlannedEndDate
	}
	return nil
}

func (x *ProjectResponse) GetActualEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ActualEndDate
	}
	return nil
}

func (x *ProjectResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ProjectResponse) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ProjectResponse) GetManagerId() uint32 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

func (x *ProjectResponse) GetBudget() float64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

type CreateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectName        string                 `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	ProjectDescription string                 `protobuf:"bytes,2,opt,name=project_description,json=projectDescription,proto3" json:"project_description,omitempty"`
	StartDate          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	PlannedEndDate     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=planned_end_date,json=plannedEndDate,proto3" json:"planned_end_date,omitempty"`
	ActualEndDate      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=actual_end_date,json=actualEndDate,proto3" json:"actual_end_date,omitempty"`
	Status             string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Priority           uint32                 `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	ManagerId          uint32                 `protobuf:"varint,8,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	Budget             float64                `protobuf:"fixed64,9,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *CreateProjectRequest) Reset() {
	*x = CreateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRequest) ProtoMessage() {}

func (x *CreateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRequest.ProtoReflect.Descriptor instead.
func (*CreateProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_management_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProjectRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *CreateProjectRequest) GetProjectDescription() string {
	if x != nil {
		return x.ProjectDescription
	}
	return ""
}

func (x *CreateProjectRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *CreateProjectRequest) GetPlannedEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PlannedEndDate
	}
	return nil
}

func (x *CreateProjectRequest) GetActualEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ActualEndDate
	}
	return nil
}

func (x *CreateProjectRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateProjectRequest) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *CreateProjectRequest) GetManagerId() uint32 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

func (x *CreateProjectRequest) GetBudget() float64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

type CreateProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *CreateProjectResponse) Reset() {
	*x = CreateProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectResponse) ProtoMessage() {}

func (x *CreateProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectResponse.ProtoReflect.Descriptor instead.
func (*CreateProjectResponse) Descriptor() ([]byte, []int) {
	return file_project_management_proto_rawDescGZIP(), []int{3}
}

func (x *CreateProjectResponse) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

type UpdateProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId          uint32                 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ProjectName        string                 `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	ProjectDescription string                 `protobuf:"bytes,3,opt,name=project_description,json=projectDescription,proto3" json:"project_description,omitempty"`
	StartDate          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	PlannedEndDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=planned_end_date,json=plannedEndDate,proto3" json:"planned_end_date,omitempty"`
	ActualEndDate      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=actual_end_date,json=actualEndDate,proto3" json:"actual_end_date,omitempty"`
	Status             string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Priority           uint32                 `protobuf:"varint,8,opt,name=priority,proto3" json:"priority,omitempty"`
	ManagerId          uint32                 `protobuf:"varint,9,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	Budget             float64                `protobuf:"fixed64,10,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *UpdateProjectRequest) Reset() {
	*x = UpdateProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectRequest) ProtoMessage() {}

func (x *UpdateProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectRequest.ProtoReflect.Descriptor instead.
func (*UpdateProjectRequest) Descriptor() ([]byte, []int) {
	return file_project_management_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProjectRequest) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *UpdateProjectRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *UpdateProjectRequest) GetProjectDescription() string {
	if x != nil {
		return x.ProjectDescription
	}
	return ""
}

func (x *UpdateProjectRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *UpdateProjectRequest) GetPlannedEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PlannedEndDate
	}
	return nil
}

func (x *UpdateProjectRequest) GetActualEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ActualEndDate
	}
	return nil
}

func (x *UpdateProjectRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateProjectRequest) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *UpdateProjectRequest) GetManagerId() uint32 {
	if x != nil {
		return x.ManagerId
	}
	return 0
}

func (x *UpdateProjectRequest) GetBudget() float64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

type UpdateProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint32 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *UpdateProjectResponse) Reset() {
	*x = UpdateProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProjectResponse) ProtoMessage() {}

func (x *UpdateProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_project_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProjectResponse.ProtoReflect.Descriptor instead.
func (*UpdateProjectResponse) Descriptor() ([]byte, []int) {
	return file_project_management_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProjectResponse) GetProjectId() uint32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

var File_project_management_proto protoreflect.FileDescriptor

var file_project_management_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2f, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0xb4, 0x03, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65,
	0x64, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x61,
	0x63, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x9a, 0x03, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x44, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x45, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x75,
	0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x22, 0x36, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0xb9, 0x03, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x70, 0x6c, 0x61, 0x6e, 0x6e,
	0x65, 0x64, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x61, 0x63, 0x74,
	0x75, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x36, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x32, 0xbd, 0x02, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x28, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_management_proto_rawDescOnce sync.Once
	file_project_management_proto_rawDescData = file_project_management_proto_rawDesc
)

func file_project_management_proto_rawDescGZIP() []byte {
	file_project_management_proto_rawDescOnce.Do(func() {
		file_project_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_management_proto_rawDescData)
	})
	return file_project_management_proto_rawDescData
}

var file_project_management_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_project_management_proto_goTypes = []any{
	(*ProjectRequest)(nil),        // 0: project_management.ProjectRequest
	(*ProjectResponse)(nil),       // 1: project_management.ProjectResponse
	(*CreateProjectRequest)(nil),  // 2: project_management.CreateProjectRequest
	(*CreateProjectResponse)(nil), // 3: project_management.CreateProjectResponse
	(*UpdateProjectRequest)(nil),  // 4: project_management.UpdateProjectRequest
	(*UpdateProjectResponse)(nil), // 5: project_management.UpdateProjectResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_project_management_proto_depIdxs = []int32{
	6,  // 0: project_management.ProjectResponse.start_date:type_name -> google.protobuf.Timestamp
	6,  // 1: project_management.ProjectResponse.planned_end_date:type_name -> google.protobuf.Timestamp
	6,  // 2: project_management.ProjectResponse.actual_end_date:type_name -> google.protobuf.Timestamp
	6,  // 3: project_management.CreateProjectRequest.start_date:type_name -> google.protobuf.Timestamp
	6,  // 4: project_management.CreateProjectRequest.planned_end_date:type_name -> google.protobuf.Timestamp
	6,  // 5: project_management.CreateProjectRequest.actual_end_date:type_name -> google.protobuf.Timestamp
	6,  // 6: project_management.UpdateProjectRequest.start_date:type_name -> google.protobuf.Timestamp
	6,  // 7: project_management.UpdateProjectRequest.planned_end_date:type_name -> google.protobuf.Timestamp
	6,  // 8: project_management.UpdateProjectRequest.actual_end_date:type_name -> google.protobuf.Timestamp
	0,  // 9: project_management.ProjectManagementService.GetProject:input_type -> project_management.ProjectRequest
	2,  // 10: project_management.ProjectManagementService.CreateProject:input_type -> project_management.CreateProjectRequest
	4,  // 11: project_management.ProjectManagementService.UpdateProject:input_type -> project_management.UpdateProjectRequest
	1,  // 12: project_management.ProjectManagementService.GetProject:output_type -> project_management.ProjectResponse
	3,  // 13: project_management.ProjectManagementService.CreateProject:output_type -> project_management.CreateProjectResponse
	5,  // 14: project_management.ProjectManagementService.UpdateProject:output_type -> project_management.UpdateProjectResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_project_management_proto_init() }
func file_project_management_proto_init() {
	if File_project_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_management_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProjectRequest); i {
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
		file_project_management_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProjectResponse); i {
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
		file_project_management_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProjectRequest); i {
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
		file_project_management_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProjectResponse); i {
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
		file_project_management_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProjectRequest); i {
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
		file_project_management_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProjectResponse); i {
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
			RawDescriptor: file_project_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_management_proto_goTypes,
		DependencyIndexes: file_project_management_proto_depIdxs,
		MessageInfos:      file_project_management_proto_msgTypes,
	}.Build()
	File_project_management_proto = out.File
	file_project_management_proto_rawDesc = nil
	file_project_management_proto_goTypes = nil
	file_project_management_proto_depIdxs = nil
}

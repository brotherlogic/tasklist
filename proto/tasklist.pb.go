// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: tasklist.proto

package proto

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

type Task_State int32

const (
	Task_UNKNOWN          Task_State = 0
	Task_TASK_WAITING     Task_State = 1
	Task_TASK_IN_PROGRESS Task_State = 2
	Task_TASK_COMPLETE    Task_State = 3
)

// Enum value maps for Task_State.
var (
	Task_State_name = map[int32]string{
		0: "UNKNOWN",
		1: "TASK_WAITING",
		2: "TASK_IN_PROGRESS",
		3: "TASK_COMPLETE",
	}
	Task_State_value = map[string]int32{
		"UNKNOWN":          0,
		"TASK_WAITING":     1,
		"TASK_IN_PROGRESS": 2,
		"TASK_COMPLETE":    3,
	}
)

func (x Task_State) Enum() *Task_State {
	p := new(Task_State)
	*p = x
	return p
}

func (x Task_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_State) Descriptor() protoreflect.EnumDescriptor {
	return file_tasklist_proto_enumTypes[0].Descriptor()
}

func (Task_State) Type() protoreflect.EnumType {
	return &file_tasklist_proto_enumTypes[0]
}

func (x Task_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_State.Descriptor instead.
func (Task_State) EnumDescriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{1, 0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lists         []*TaskList `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
	TrackingIssue int32       `protobuf:"varint,2,opt,name=tracking_issue,json=trackingIssue,proto3" json:"tracking_issue,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetLists() []*TaskList {
	if x != nil {
		return x.Lists
	}
	return nil
}

func (x *Config) GetTrackingIssue() int32 {
	if x != nil {
		return x.TrackingIssue
	}
	return 0
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Job         string     `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	IssueNumber int32      `protobuf:"varint,3,opt,name=issue_number,json=issueNumber,proto3" json:"issue_number,omitempty"`
	State       Task_State `protobuf:"varint,4,opt,name=state,proto3,enum=tasklist.Task_State" json:"state,omitempty"`
	Index       int32      `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *Task) GetIssueNumber() int32 {
	if x != nil {
		return x.IssueNumber
	}
	return 0
}

func (x *Task) GetState() Task_State {
	if x != nil {
		return x.State
	}
	return Task_UNKNOWN
}

func (x *Task) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type TaskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks       []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Job         string  `protobuf:"bytes,3,opt,name=job,proto3" json:"job,omitempty"`
	IssueNumber int32   `protobuf:"varint,4,opt,name=issue_number,json=issueNumber,proto3" json:"issue_number,omitempty"`
	Tag         string  `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *TaskList) Reset() {
	*x = TaskList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskList) ProtoMessage() {}

func (x *TaskList) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskList.ProtoReflect.Descriptor instead.
func (*TaskList) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{2}
}

func (x *TaskList) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *TaskList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskList) GetJob() string {
	if x != nil {
		return x.Job
	}
	return ""
}

func (x *TaskList) GetIssueNumber() int32 {
	if x != nil {
		return x.IssueNumber
	}
	return 0
}

func (x *TaskList) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type AddTaskListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add *TaskList `protobuf:"bytes,1,opt,name=add,proto3" json:"add,omitempty"`
}

func (x *AddTaskListRequest) Reset() {
	*x = AddTaskListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskListRequest) ProtoMessage() {}

func (x *AddTaskListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskListRequest.ProtoReflect.Descriptor instead.
func (*AddTaskListRequest) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{3}
}

func (x *AddTaskListRequest) GetAdd() *TaskList {
	if x != nil {
		return x.Add
	}
	return nil
}

type AddTaskListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddTaskListResponse) Reset() {
	*x = AddTaskListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskListResponse) ProtoMessage() {}

func (x *AddTaskListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskListResponse.ProtoReflect.Descriptor instead.
func (*AddTaskListResponse) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{4}
}

type GetTaskListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTaskListsRequest) Reset() {
	*x = GetTaskListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskListsRequest) ProtoMessage() {}

func (x *GetTaskListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskListsRequest.ProtoReflect.Descriptor instead.
func (*GetTaskListsRequest) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{5}
}

type GetTaskListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lists []*TaskList `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
}

func (x *GetTaskListsResponse) Reset() {
	*x = GetTaskListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskListsResponse) ProtoMessage() {}

func (x *GetTaskListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskListsResponse.ProtoReflect.Descriptor instead.
func (*GetTaskListsResponse) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{6}
}

func (x *GetTaskListsResponse) GetLists() []*TaskList {
	if x != nil {
		return x.Lists
	}
	return nil
}

type ValidateTaskListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValidateTaskListsRequest) Reset() {
	*x = ValidateTaskListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTaskListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTaskListsRequest) ProtoMessage() {}

func (x *ValidateTaskListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTaskListsRequest.ProtoReflect.Descriptor instead.
func (*ValidateTaskListsRequest) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{7}
}

type ValidateTaskListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ValidateTaskListsResponse) Reset() {
	*x = ValidateTaskListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTaskListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTaskListsResponse) ProtoMessage() {}

func (x *ValidateTaskListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTaskListsResponse.ProtoReflect.Descriptor instead.
func (*ValidateTaskListsResponse) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{8}
}

type GetTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *GetTasksRequest) Reset() {
	*x = GetTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksRequest) ProtoMessage() {}

func (x *GetTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksRequest.ProtoReflect.Descriptor instead.
func (*GetTasksRequest) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{9}
}

func (x *GetTasksRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type GetTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetTasksResponse) Reset() {
	*x = GetTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksResponse) ProtoMessage() {}

func (x *GetTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksResponse.ProtoReflect.Descriptor instead.
func (*GetTasksResponse) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{10}
}

func (x *GetTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type RenameJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldJob string `protobuf:"bytes,1,opt,name=old_job,json=oldJob,proto3" json:"old_job,omitempty"`
	NewJob string `protobuf:"bytes,2,opt,name=new_job,json=newJob,proto3" json:"new_job,omitempty"`
}

func (x *RenameJobRequest) Reset() {
	*x = RenameJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameJobRequest) ProtoMessage() {}

func (x *RenameJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameJobRequest.ProtoReflect.Descriptor instead.
func (*RenameJobRequest) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{11}
}

func (x *RenameJobRequest) GetOldJob() string {
	if x != nil {
		return x.OldJob
	}
	return ""
}

func (x *RenameJobRequest) GetNewJob() string {
	if x != nil {
		return x.NewJob
	}
	return ""
}

type RenameJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RenameJobResponse) Reset() {
	*x = RenameJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tasklist_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameJobResponse) ProtoMessage() {}

func (x *RenameJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tasklist_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameJobResponse.ProtoReflect.Descriptor instead.
func (*RenameJobResponse) Descriptor() ([]byte, []int) {
	return file_tasklist_proto_rawDescGZIP(), []int{12}
}

var File_tasklist_proto protoreflect.FileDescriptor

var file_tasklist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x22, 0xe4, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x4f, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x49, 0x4e, 0x5f, 0x50,
	0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x22, 0x8b, 0x01, 0x0a,
	0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x3a, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x03, 0x61, 0x64, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x44, 0x0a, 0x10, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x6c, 0x64, 0x5f, 0x6a, 0x6f, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6c, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x65, 0x77, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x65, 0x77, 0x4a, 0x6f, 0x62, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x93, 0x03, 0x0a, 0x0f,
	0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x11, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12,
	0x22, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x12, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x52,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tasklist_proto_rawDescOnce sync.Once
	file_tasklist_proto_rawDescData = file_tasklist_proto_rawDesc
)

func file_tasklist_proto_rawDescGZIP() []byte {
	file_tasklist_proto_rawDescOnce.Do(func() {
		file_tasklist_proto_rawDescData = protoimpl.X.CompressGZIP(file_tasklist_proto_rawDescData)
	})
	return file_tasklist_proto_rawDescData
}

var file_tasklist_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tasklist_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tasklist_proto_goTypes = []interface{}{
	(Task_State)(0),                   // 0: tasklist.Task.State
	(*Config)(nil),                    // 1: tasklist.Config
	(*Task)(nil),                      // 2: tasklist.Task
	(*TaskList)(nil),                  // 3: tasklist.TaskList
	(*AddTaskListRequest)(nil),        // 4: tasklist.AddTaskListRequest
	(*AddTaskListResponse)(nil),       // 5: tasklist.AddTaskListResponse
	(*GetTaskListsRequest)(nil),       // 6: tasklist.GetTaskListsRequest
	(*GetTaskListsResponse)(nil),      // 7: tasklist.GetTaskListsResponse
	(*ValidateTaskListsRequest)(nil),  // 8: tasklist.ValidateTaskListsRequest
	(*ValidateTaskListsResponse)(nil), // 9: tasklist.ValidateTaskListsResponse
	(*GetTasksRequest)(nil),           // 10: tasklist.GetTasksRequest
	(*GetTasksResponse)(nil),          // 11: tasklist.GetTasksResponse
	(*RenameJobRequest)(nil),          // 12: tasklist.RenameJobRequest
	(*RenameJobResponse)(nil),         // 13: tasklist.RenameJobResponse
}
var file_tasklist_proto_depIdxs = []int32{
	3,  // 0: tasklist.Config.lists:type_name -> tasklist.TaskList
	0,  // 1: tasklist.Task.state:type_name -> tasklist.Task.State
	2,  // 2: tasklist.TaskList.tasks:type_name -> tasklist.Task
	3,  // 3: tasklist.AddTaskListRequest.add:type_name -> tasklist.TaskList
	3,  // 4: tasklist.GetTaskListsResponse.lists:type_name -> tasklist.TaskList
	2,  // 5: tasklist.GetTasksResponse.tasks:type_name -> tasklist.Task
	4,  // 6: tasklist.TaskListService.AddTaskList:input_type -> tasklist.AddTaskListRequest
	6,  // 7: tasklist.TaskListService.GetTaskLists:input_type -> tasklist.GetTaskListsRequest
	8,  // 8: tasklist.TaskListService.ValidateTaskLists:input_type -> tasklist.ValidateTaskListsRequest
	10, // 9: tasklist.TaskListService.GetTasks:input_type -> tasklist.GetTasksRequest
	12, // 10: tasklist.TaskListService.RenameJob:input_type -> tasklist.RenameJobRequest
	5,  // 11: tasklist.TaskListService.AddTaskList:output_type -> tasklist.AddTaskListResponse
	7,  // 12: tasklist.TaskListService.GetTaskLists:output_type -> tasklist.GetTaskListsResponse
	9,  // 13: tasklist.TaskListService.ValidateTaskLists:output_type -> tasklist.ValidateTaskListsResponse
	11, // 14: tasklist.TaskListService.GetTasks:output_type -> tasklist.GetTasksResponse
	13, // 15: tasklist.TaskListService.RenameJob:output_type -> tasklist.RenameJobResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_tasklist_proto_init() }
func file_tasklist_proto_init() {
	if File_tasklist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tasklist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_tasklist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_tasklist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskList); i {
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
		file_tasklist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskListRequest); i {
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
		file_tasklist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskListResponse); i {
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
		file_tasklist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskListsRequest); i {
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
		file_tasklist_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskListsResponse); i {
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
		file_tasklist_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTaskListsRequest); i {
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
		file_tasklist_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTaskListsResponse); i {
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
		file_tasklist_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksRequest); i {
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
		file_tasklist_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTasksResponse); i {
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
		file_tasklist_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameJobRequest); i {
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
		file_tasklist_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameJobResponse); i {
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
			RawDescriptor: file_tasklist_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tasklist_proto_goTypes,
		DependencyIndexes: file_tasklist_proto_depIdxs,
		EnumInfos:         file_tasklist_proto_enumTypes,
		MessageInfos:      file_tasklist_proto_msgTypes,
	}.Build()
	File_tasklist_proto = out.File
	file_tasklist_proto_rawDesc = nil
	file_tasklist_proto_goTypes = nil
	file_tasklist_proto_depIdxs = nil
}

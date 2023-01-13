// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
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

	Lists []*TaskList `protobuf:"bytes,1,rep,name=lists,proto3" json:"lists,omitempty"`
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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Job         string     `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	IssueNumber int32      `protobuf:"varint,3,opt,name=issue_number,json=issueNumber,proto3" json:"issue_number,omitempty"`
	State       Task_State `protobuf:"varint,4,opt,name=state,proto3,enum=tasklist.Task_State" json:"state,omitempty"`
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

type TaskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Name  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
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

var File_tasklist_proto protoreflect.FileDescriptor

var file_tasklist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x22, 0xce,
	0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6a, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x4f,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x57, 0x41, 0x49,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x49,
	0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d,
	0x54, 0x41, 0x53, 0x4b, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x22,
	0x44, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x74,
	0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x61,
	0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x03, 0x61, 0x64,
	0x64, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5d, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x41,
	0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_tasklist_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tasklist_proto_goTypes = []interface{}{
	(Task_State)(0),             // 0: tasklist.Task.State
	(*Config)(nil),              // 1: tasklist.Config
	(*Task)(nil),                // 2: tasklist.Task
	(*TaskList)(nil),            // 3: tasklist.TaskList
	(*AddTaskListRequest)(nil),  // 4: tasklist.AddTaskListRequest
	(*AddTaskListResponse)(nil), // 5: tasklist.AddTaskListResponse
}
var file_tasklist_proto_depIdxs = []int32{
	3, // 0: tasklist.Config.lists:type_name -> tasklist.TaskList
	0, // 1: tasklist.Task.state:type_name -> tasklist.Task.State
	2, // 2: tasklist.TaskList.tasks:type_name -> tasklist.Task
	3, // 3: tasklist.AddTaskListRequest.add:type_name -> tasklist.TaskList
	4, // 4: tasklist.TaskListService.AddTaskList:input_type -> tasklist.AddTaskListRequest
	5, // 5: tasklist.TaskListService.AddTaskList:output_type -> tasklist.AddTaskListResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tasklist_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: tasklist.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TaskListServiceClient is the client API for TaskListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskListServiceClient interface {
	AddTaskList(ctx context.Context, in *AddTaskListRequest, opts ...grpc.CallOption) (*AddTaskListResponse, error)
	GetTaskLists(ctx context.Context, in *GetTaskListsRequest, opts ...grpc.CallOption) (*GetTaskListsResponse, error)
}

type taskListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskListServiceClient(cc grpc.ClientConnInterface) TaskListServiceClient {
	return &taskListServiceClient{cc}
}

func (c *taskListServiceClient) AddTaskList(ctx context.Context, in *AddTaskListRequest, opts ...grpc.CallOption) (*AddTaskListResponse, error) {
	out := new(AddTaskListResponse)
	err := c.cc.Invoke(ctx, "/tasklist.TaskListService/AddTaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskListServiceClient) GetTaskLists(ctx context.Context, in *GetTaskListsRequest, opts ...grpc.CallOption) (*GetTaskListsResponse, error) {
	out := new(GetTaskListsResponse)
	err := c.cc.Invoke(ctx, "/tasklist.TaskListService/GetTaskLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskListServiceServer is the server API for TaskListService service.
// All implementations should embed UnimplementedTaskListServiceServer
// for forward compatibility
type TaskListServiceServer interface {
	AddTaskList(context.Context, *AddTaskListRequest) (*AddTaskListResponse, error)
	GetTaskLists(context.Context, *GetTaskListsRequest) (*GetTaskListsResponse, error)
}

// UnimplementedTaskListServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTaskListServiceServer struct {
}

func (UnimplementedTaskListServiceServer) AddTaskList(context.Context, *AddTaskListRequest) (*AddTaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTaskList not implemented")
}
func (UnimplementedTaskListServiceServer) GetTaskLists(context.Context, *GetTaskListsRequest) (*GetTaskListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskLists not implemented")
}

// UnsafeTaskListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskListServiceServer will
// result in compilation errors.
type UnsafeTaskListServiceServer interface {
	mustEmbedUnimplementedTaskListServiceServer()
}

func RegisterTaskListServiceServer(s grpc.ServiceRegistrar, srv TaskListServiceServer) {
	s.RegisterService(&TaskListService_ServiceDesc, srv)
}

func _TaskListService_AddTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskListServiceServer).AddTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasklist.TaskListService/AddTaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskListServiceServer).AddTaskList(ctx, req.(*AddTaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskListService_GetTaskLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskListServiceServer).GetTaskLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasklist.TaskListService/GetTaskLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskListServiceServer).GetTaskLists(ctx, req.(*GetTaskListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskListService_ServiceDesc is the grpc.ServiceDesc for TaskListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tasklist.TaskListService",
	HandlerType: (*TaskListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTaskList",
			Handler:    _TaskListService_AddTaskList_Handler,
		},
		{
			MethodName: "GetTaskLists",
			Handler:    _TaskListService_GetTaskLists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tasklist.proto",
}

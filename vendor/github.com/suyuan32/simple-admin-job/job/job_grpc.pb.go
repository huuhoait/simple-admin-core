// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: job.proto

package job

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

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	// group: base
	InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
	// Task management
	// group: task
	CreateTask(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
	// group: task
	UpdateTask(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (*BaseResp, error)
	// group: task
	GetTaskList(ctx context.Context, in *TaskListReq, opts ...grpc.CallOption) (*TaskListResp, error)
	// group: task
	GetTaskById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*TaskInfo, error)
	// group: task
	DeleteTask(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	// TaskLog management
	// group: tasklog
	CreateTaskLog(ctx context.Context, in *TaskLogInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
	// group: tasklog
	UpdateTaskLog(ctx context.Context, in *TaskLogInfo, opts ...grpc.CallOption) (*BaseResp, error)
	// group: tasklog
	GetTaskLogList(ctx context.Context, in *TaskLogListReq, opts ...grpc.CallOption) (*TaskLogListResp, error)
	// group: tasklog
	GetTaskLogById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*TaskLogInfo, error)
	// group: tasklog
	DeleteTaskLog(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/job.Job/initDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) CreateTask(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	out := new(BaseIDResp)
	err := c.cc.Invoke(ctx, "/job.Job/createTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) UpdateTask(ctx context.Context, in *TaskInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/job.Job/updateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetTaskList(ctx context.Context, in *TaskListReq, opts ...grpc.CallOption) (*TaskListResp, error) {
	out := new(TaskListResp)
	err := c.cc.Invoke(ctx, "/job.Job/getTaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetTaskById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*TaskInfo, error) {
	out := new(TaskInfo)
	err := c.cc.Invoke(ctx, "/job.Job/getTaskById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) DeleteTask(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/job.Job/deleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) CreateTaskLog(ctx context.Context, in *TaskLogInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	out := new(BaseIDResp)
	err := c.cc.Invoke(ctx, "/job.Job/createTaskLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) UpdateTaskLog(ctx context.Context, in *TaskLogInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/job.Job/updateTaskLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetTaskLogList(ctx context.Context, in *TaskLogListReq, opts ...grpc.CallOption) (*TaskLogListResp, error) {
	out := new(TaskLogListResp)
	err := c.cc.Invoke(ctx, "/job.Job/getTaskLogList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetTaskLogById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*TaskLogInfo, error) {
	out := new(TaskLogInfo)
	err := c.cc.Invoke(ctx, "/job.Job/getTaskLogById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) DeleteTaskLog(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	out := new(BaseResp)
	err := c.cc.Invoke(ctx, "/job.Job/deleteTaskLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility
type JobServer interface {
	// group: base
	InitDatabase(context.Context, *Empty) (*BaseResp, error)
	// Task management
	// group: task
	CreateTask(context.Context, *TaskInfo) (*BaseIDResp, error)
	// group: task
	UpdateTask(context.Context, *TaskInfo) (*BaseResp, error)
	// group: task
	GetTaskList(context.Context, *TaskListReq) (*TaskListResp, error)
	// group: task
	GetTaskById(context.Context, *IDReq) (*TaskInfo, error)
	// group: task
	DeleteTask(context.Context, *IDsReq) (*BaseResp, error)
	// TaskLog management
	// group: tasklog
	CreateTaskLog(context.Context, *TaskLogInfo) (*BaseIDResp, error)
	// group: tasklog
	UpdateTaskLog(context.Context, *TaskLogInfo) (*BaseResp, error)
	// group: tasklog
	GetTaskLogList(context.Context, *TaskLogListReq) (*TaskLogListResp, error)
	// group: tasklog
	GetTaskLogById(context.Context, *IDReq) (*TaskLogInfo, error)
	// group: tasklog
	DeleteTaskLog(context.Context, *IDsReq) (*BaseResp, error)
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (UnimplementedJobServer) InitDatabase(context.Context, *Empty) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitDatabase not implemented")
}
func (UnimplementedJobServer) CreateTask(context.Context, *TaskInfo) (*BaseIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedJobServer) UpdateTask(context.Context, *TaskInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedJobServer) GetTaskList(context.Context, *TaskListReq) (*TaskListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskList not implemented")
}
func (UnimplementedJobServer) GetTaskById(context.Context, *IDReq) (*TaskInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskById not implemented")
}
func (UnimplementedJobServer) DeleteTask(context.Context, *IDsReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedJobServer) CreateTaskLog(context.Context, *TaskLogInfo) (*BaseIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskLog not implemented")
}
func (UnimplementedJobServer) UpdateTaskLog(context.Context, *TaskLogInfo) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskLog not implemented")
}
func (UnimplementedJobServer) GetTaskLogList(context.Context, *TaskLogListReq) (*TaskLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskLogList not implemented")
}
func (UnimplementedJobServer) GetTaskLogById(context.Context, *IDReq) (*TaskLogInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskLogById not implemented")
}
func (UnimplementedJobServer) DeleteTaskLog(context.Context, *IDsReq) (*BaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTaskLog not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	s.RegisterService(&Job_ServiceDesc, srv)
}

func _Job_InitDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).InitDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/initDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).InitDatabase(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/createTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).CreateTask(ctx, req.(*TaskInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/updateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).UpdateTask(ctx, req.(*TaskInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/getTaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetTaskList(ctx, req.(*TaskListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetTaskById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetTaskById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/getTaskById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetTaskById(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/deleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).DeleteTask(ctx, req.(*IDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_CreateTaskLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskLogInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).CreateTaskLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/createTaskLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).CreateTaskLog(ctx, req.(*TaskLogInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_UpdateTaskLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskLogInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).UpdateTaskLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/updateTaskLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).UpdateTaskLog(ctx, req.(*TaskLogInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetTaskLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetTaskLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/getTaskLogList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetTaskLogList(ctx, req.(*TaskLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetTaskLogById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetTaskLogById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/getTaskLogById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetTaskLogById(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_DeleteTaskLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).DeleteTaskLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.Job/deleteTaskLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).DeleteTaskLog(ctx, req.(*IDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Job_ServiceDesc is the grpc.ServiceDesc for Job service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Job_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "initDatabase",
			Handler:    _Job_InitDatabase_Handler,
		},
		{
			MethodName: "createTask",
			Handler:    _Job_CreateTask_Handler,
		},
		{
			MethodName: "updateTask",
			Handler:    _Job_UpdateTask_Handler,
		},
		{
			MethodName: "getTaskList",
			Handler:    _Job_GetTaskList_Handler,
		},
		{
			MethodName: "getTaskById",
			Handler:    _Job_GetTaskById_Handler,
		},
		{
			MethodName: "deleteTask",
			Handler:    _Job_DeleteTask_Handler,
		},
		{
			MethodName: "createTaskLog",
			Handler:    _Job_CreateTaskLog_Handler,
		},
		{
			MethodName: "updateTaskLog",
			Handler:    _Job_UpdateTaskLog_Handler,
		},
		{
			MethodName: "getTaskLogList",
			Handler:    _Job_GetTaskLogList_Handler,
		},
		{
			MethodName: "getTaskLogById",
			Handler:    _Job_GetTaskLogById_Handler,
		},
		{
			MethodName: "deleteTaskLog",
			Handler:    _Job_DeleteTaskLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}

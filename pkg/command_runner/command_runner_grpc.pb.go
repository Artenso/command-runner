// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: command_runner.proto

package command_runner

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CommandRunner_AddCommand_FullMethodName  = "/github.com.Artenso.command_runner.api.command_runner.CommandRunner/AddCommand"
	CommandRunner_GetCommand_FullMethodName  = "/github.com.Artenso.command_runner.api.command_runner.CommandRunner/GetCommand"
	CommandRunner_ListCommand_FullMethodName = "/github.com.Artenso.command_runner.api.command_runner.CommandRunner/ListCommand"
	CommandRunner_StopCommand_FullMethodName = "/github.com.Artenso.command_runner.api.command_runner.CommandRunner/StopCommand"
)

// CommandRunnerClient is the client API for CommandRunner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandRunnerClient interface {
	AddCommand(ctx context.Context, in *AddCommandRequest, opts ...grpc.CallOption) (*AddCommandResponse, error)
	GetCommand(ctx context.Context, in *GetCommandRequest, opts ...grpc.CallOption) (*GetCommandResponse, error)
	ListCommand(ctx context.Context, in *ListCommandRequest, opts ...grpc.CallOption) (*ListCommandResponse, error)
	StopCommand(ctx context.Context, in *StopCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type commandRunnerClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandRunnerClient(cc grpc.ClientConnInterface) CommandRunnerClient {
	return &commandRunnerClient{cc}
}

func (c *commandRunnerClient) AddCommand(ctx context.Context, in *AddCommandRequest, opts ...grpc.CallOption) (*AddCommandResponse, error) {
	out := new(AddCommandResponse)
	err := c.cc.Invoke(ctx, CommandRunner_AddCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandRunnerClient) GetCommand(ctx context.Context, in *GetCommandRequest, opts ...grpc.CallOption) (*GetCommandResponse, error) {
	out := new(GetCommandResponse)
	err := c.cc.Invoke(ctx, CommandRunner_GetCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandRunnerClient) ListCommand(ctx context.Context, in *ListCommandRequest, opts ...grpc.CallOption) (*ListCommandResponse, error) {
	out := new(ListCommandResponse)
	err := c.cc.Invoke(ctx, CommandRunner_ListCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandRunnerClient) StopCommand(ctx context.Context, in *StopCommandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CommandRunner_StopCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandRunnerServer is the server API for CommandRunner service.
// All implementations must embed UnimplementedCommandRunnerServer
// for forward compatibility
type CommandRunnerServer interface {
	AddCommand(context.Context, *AddCommandRequest) (*AddCommandResponse, error)
	GetCommand(context.Context, *GetCommandRequest) (*GetCommandResponse, error)
	ListCommand(context.Context, *ListCommandRequest) (*ListCommandResponse, error)
	StopCommand(context.Context, *StopCommandRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCommandRunnerServer()
}

// UnimplementedCommandRunnerServer must be embedded to have forward compatible implementations.
type UnimplementedCommandRunnerServer struct {
}

func (UnimplementedCommandRunnerServer) AddCommand(context.Context, *AddCommandRequest) (*AddCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCommand not implemented")
}
func (UnimplementedCommandRunnerServer) GetCommand(context.Context, *GetCommandRequest) (*GetCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommand not implemented")
}
func (UnimplementedCommandRunnerServer) ListCommand(context.Context, *ListCommandRequest) (*ListCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommand not implemented")
}
func (UnimplementedCommandRunnerServer) StopCommand(context.Context, *StopCommandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopCommand not implemented")
}
func (UnimplementedCommandRunnerServer) mustEmbedUnimplementedCommandRunnerServer() {}

// UnsafeCommandRunnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandRunnerServer will
// result in compilation errors.
type UnsafeCommandRunnerServer interface {
	mustEmbedUnimplementedCommandRunnerServer()
}

func RegisterCommandRunnerServer(s grpc.ServiceRegistrar, srv CommandRunnerServer) {
	s.RegisterService(&CommandRunner_ServiceDesc, srv)
}

func _CommandRunner_AddCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandRunnerServer).AddCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandRunner_AddCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandRunnerServer).AddCommand(ctx, req.(*AddCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandRunner_GetCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandRunnerServer).GetCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandRunner_GetCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandRunnerServer).GetCommand(ctx, req.(*GetCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandRunner_ListCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandRunnerServer).ListCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandRunner_ListCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandRunnerServer).ListCommand(ctx, req.(*ListCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandRunner_StopCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandRunnerServer).StopCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandRunner_StopCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandRunnerServer).StopCommand(ctx, req.(*StopCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandRunner_ServiceDesc is the grpc.ServiceDesc for CommandRunner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandRunner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.Artenso.command_runner.api.command_runner.CommandRunner",
	HandlerType: (*CommandRunnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCommand",
			Handler:    _CommandRunner_AddCommand_Handler,
		},
		{
			MethodName: "GetCommand",
			Handler:    _CommandRunner_GetCommand_Handler,
		},
		{
			MethodName: "ListCommand",
			Handler:    _CommandRunner_ListCommand_Handler,
		},
		{
			MethodName: "StopCommand",
			Handler:    _CommandRunner_StopCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "command_runner.proto",
}
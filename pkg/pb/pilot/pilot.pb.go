// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openpitrix/pilot/pilot.proto

package pbpilot // import "openpitrix.io/openpitrix/pkg/pb/pilot"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import types "openpitrix.io/openpitrix/pkg/pb/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PilotService service

type PilotServiceClient interface {
	GetPilotConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.PilotConfig, error)
	GetFrontgateList(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.FrontgateIdList, error)
	GetFrontgateConfig(ctx context.Context, in *types.FrontgateId, opts ...grpc.CallOption) (*types.FrontgateConfig, error)
	SetFrontgateConfig(ctx context.Context, in *types.FrontgateConfig, opts ...grpc.CallOption) (*types.Empty, error)
	GetDroneConfig(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.DroneConfig, error)
	SetDroneConfig(ctx context.Context, in *types.SetDroneConfigRequest, opts ...grpc.CallOption) (*types.Empty, error)
	IsConfdRunning(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Bool, error)
	StartConfd(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	StopConfd(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	RegisterMetadata(ctx context.Context, in *types.SubTask_RegisterMetadata, opts ...grpc.CallOption) (*types.Empty, error)
	DeregisterMetadata(ctx context.Context, in *types.SubTask_DeregisterMetadata, opts ...grpc.CallOption) (*types.Empty, error)
	RegisterCmd(ctx context.Context, in *types.SubTask_RegisterCmd, opts ...grpc.CallOption) (*types.Empty, error)
	DeregisterCmd(ctx context.Context, in *types.SubTask_DeregisterCmd, opts ...grpc.CallOption) (*types.Empty, error)
	ReportSubTaskStatus(ctx context.Context, in *types.SubTaskStatus, opts ...grpc.CallOption) (*types.Empty, error)
	GetSubtaskStatus(ctx context.Context, in *types.SubTaskId, opts ...grpc.CallOption) (*types.SubTaskStatus, error)
	HandleSubtask(ctx context.Context, in *types.SubTaskMessage, opts ...grpc.CallOption) (*types.Empty, error)
	PingPilot(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error)
	PingFrontgate(ctx context.Context, in *types.FrontgateId, opts ...grpc.CallOption) (*types.Empty, error)
	PingDrone(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error)
	FrontgateChannel(ctx context.Context, opts ...grpc.CallOption) (PilotService_FrontgateChannelClient, error)
}

type pilotServiceClient struct {
	cc *grpc.ClientConn
}

func NewPilotServiceClient(cc *grpc.ClientConn) PilotServiceClient {
	return &pilotServiceClient{cc}
}

func (c *pilotServiceClient) GetPilotConfig(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.PilotConfig, error) {
	out := new(types.PilotConfig)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetPilotConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetFrontgateList(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.FrontgateIdList, error) {
	out := new(types.FrontgateIdList)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetFrontgateList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetFrontgateConfig(ctx context.Context, in *types.FrontgateId, opts ...grpc.CallOption) (*types.FrontgateConfig, error) {
	out := new(types.FrontgateConfig)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetFrontgateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) SetFrontgateConfig(ctx context.Context, in *types.FrontgateConfig, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/SetFrontgateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetDroneConfig(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.DroneConfig, error) {
	out := new(types.DroneConfig)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetDroneConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) SetDroneConfig(ctx context.Context, in *types.SetDroneConfigRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/SetDroneConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) IsConfdRunning(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Bool, error) {
	out := new(types.Bool)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/IsConfdRunning", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) StartConfd(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/StartConfd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) StopConfd(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/StopConfd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) RegisterMetadata(ctx context.Context, in *types.SubTask_RegisterMetadata, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/RegisterMetadata", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) DeregisterMetadata(ctx context.Context, in *types.SubTask_DeregisterMetadata, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/DeregisterMetadata", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) RegisterCmd(ctx context.Context, in *types.SubTask_RegisterCmd, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/RegisterCmd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) DeregisterCmd(ctx context.Context, in *types.SubTask_DeregisterCmd, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/DeregisterCmd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) ReportSubTaskStatus(ctx context.Context, in *types.SubTaskStatus, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/ReportSubTaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) GetSubtaskStatus(ctx context.Context, in *types.SubTaskId, opts ...grpc.CallOption) (*types.SubTaskStatus, error) {
	out := new(types.SubTaskStatus)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/GetSubtaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) HandleSubtask(ctx context.Context, in *types.SubTaskMessage, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/HandleSubtask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) PingPilot(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/PingPilot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) PingFrontgate(ctx context.Context, in *types.FrontgateId, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/PingFrontgate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) PingDrone(ctx context.Context, in *types.DroneEndpoint, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := grpc.Invoke(ctx, "/openpitrix.pilot.PilotService/PingDrone", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pilotServiceClient) FrontgateChannel(ctx context.Context, opts ...grpc.CallOption) (PilotService_FrontgateChannelClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PilotService_serviceDesc.Streams[0], c.cc, "/openpitrix.pilot.PilotService/FrontgateChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &pilotServiceFrontgateChannelClient{stream}
	return x, nil
}

type PilotService_FrontgateChannelClient interface {
	Send(*types.Bytes) error
	Recv() (*types.Bytes, error)
	grpc.ClientStream
}

type pilotServiceFrontgateChannelClient struct {
	grpc.ClientStream
}

func (x *pilotServiceFrontgateChannelClient) Send(m *types.Bytes) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pilotServiceFrontgateChannelClient) Recv() (*types.Bytes, error) {
	m := new(types.Bytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PilotService service

type PilotServiceServer interface {
	GetPilotConfig(context.Context, *types.Empty) (*types.PilotConfig, error)
	GetFrontgateList(context.Context, *types.Empty) (*types.FrontgateIdList, error)
	GetFrontgateConfig(context.Context, *types.FrontgateId) (*types.FrontgateConfig, error)
	SetFrontgateConfig(context.Context, *types.FrontgateConfig) (*types.Empty, error)
	GetDroneConfig(context.Context, *types.DroneEndpoint) (*types.DroneConfig, error)
	SetDroneConfig(context.Context, *types.SetDroneConfigRequest) (*types.Empty, error)
	IsConfdRunning(context.Context, *types.DroneEndpoint) (*types.Bool, error)
	StartConfd(context.Context, *types.DroneEndpoint) (*types.Empty, error)
	StopConfd(context.Context, *types.DroneEndpoint) (*types.Empty, error)
	RegisterMetadata(context.Context, *types.SubTask_RegisterMetadata) (*types.Empty, error)
	DeregisterMetadata(context.Context, *types.SubTask_DeregisterMetadata) (*types.Empty, error)
	RegisterCmd(context.Context, *types.SubTask_RegisterCmd) (*types.Empty, error)
	DeregisterCmd(context.Context, *types.SubTask_DeregisterCmd) (*types.Empty, error)
	ReportSubTaskStatus(context.Context, *types.SubTaskStatus) (*types.Empty, error)
	GetSubtaskStatus(context.Context, *types.SubTaskId) (*types.SubTaskStatus, error)
	HandleSubtask(context.Context, *types.SubTaskMessage) (*types.Empty, error)
	PingPilot(context.Context, *types.Empty) (*types.Empty, error)
	PingFrontgate(context.Context, *types.FrontgateId) (*types.Empty, error)
	PingDrone(context.Context, *types.DroneEndpoint) (*types.Empty, error)
	FrontgateChannel(PilotService_FrontgateChannelServer) error
}

func RegisterPilotServiceServer(s *grpc.Server, srv PilotServiceServer) {
	s.RegisterService(&_PilotService_serviceDesc, srv)
}

func _PilotService_GetPilotConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetPilotConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetPilotConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetPilotConfig(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetFrontgateList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetFrontgateList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetFrontgateList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetFrontgateList(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetFrontgateConfig(ctx, req.(*types.FrontgateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_SetFrontgateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).SetFrontgateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/SetFrontgateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).SetFrontgateConfig(ctx, req.(*types.FrontgateConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetDroneConfig(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_SetDroneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SetDroneConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).SetDroneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/SetDroneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).SetDroneConfig(ctx, req.(*types.SetDroneConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_IsConfdRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).IsConfdRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/IsConfdRunning",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).IsConfdRunning(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_StartConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).StartConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/StartConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).StartConfd(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_StopConfd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).StopConfd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/StopConfd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).StopConfd(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_RegisterMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTask_RegisterMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).RegisterMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/RegisterMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).RegisterMetadata(ctx, req.(*types.SubTask_RegisterMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_DeregisterMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTask_DeregisterMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).DeregisterMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/DeregisterMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).DeregisterMetadata(ctx, req.(*types.SubTask_DeregisterMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_RegisterCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTask_RegisterCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).RegisterCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/RegisterCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).RegisterCmd(ctx, req.(*types.SubTask_RegisterCmd))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_DeregisterCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTask_DeregisterCmd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).DeregisterCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/DeregisterCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).DeregisterCmd(ctx, req.(*types.SubTask_DeregisterCmd))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_ReportSubTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTaskStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).ReportSubTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/ReportSubTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).ReportSubTaskStatus(ctx, req.(*types.SubTaskStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_GetSubtaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTaskId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).GetSubtaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/GetSubtaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).GetSubtaskStatus(ctx, req.(*types.SubTaskId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_HandleSubtask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.SubTaskMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).HandleSubtask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/HandleSubtask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).HandleSubtask(ctx, req.(*types.SubTaskMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_PingPilot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).PingPilot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/PingPilot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).PingPilot(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_PingFrontgate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.FrontgateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).PingFrontgate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/PingFrontgate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).PingFrontgate(ctx, req.(*types.FrontgateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_PingDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.DroneEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PilotServiceServer).PingDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.pilot.PilotService/PingDrone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PilotServiceServer).PingDrone(ctx, req.(*types.DroneEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _PilotService_FrontgateChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PilotServiceServer).FrontgateChannel(&pilotServiceFrontgateChannelServer{stream})
}

type PilotService_FrontgateChannelServer interface {
	Send(*types.Bytes) error
	Recv() (*types.Bytes, error)
	grpc.ServerStream
}

type pilotServiceFrontgateChannelServer struct {
	grpc.ServerStream
}

func (x *pilotServiceFrontgateChannelServer) Send(m *types.Bytes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pilotServiceFrontgateChannelServer) Recv() (*types.Bytes, error) {
	m := new(types.Bytes)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PilotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.pilot.PilotService",
	HandlerType: (*PilotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPilotConfig",
			Handler:    _PilotService_GetPilotConfig_Handler,
		},
		{
			MethodName: "GetFrontgateList",
			Handler:    _PilotService_GetFrontgateList_Handler,
		},
		{
			MethodName: "GetFrontgateConfig",
			Handler:    _PilotService_GetFrontgateConfig_Handler,
		},
		{
			MethodName: "SetFrontgateConfig",
			Handler:    _PilotService_SetFrontgateConfig_Handler,
		},
		{
			MethodName: "GetDroneConfig",
			Handler:    _PilotService_GetDroneConfig_Handler,
		},
		{
			MethodName: "SetDroneConfig",
			Handler:    _PilotService_SetDroneConfig_Handler,
		},
		{
			MethodName: "IsConfdRunning",
			Handler:    _PilotService_IsConfdRunning_Handler,
		},
		{
			MethodName: "StartConfd",
			Handler:    _PilotService_StartConfd_Handler,
		},
		{
			MethodName: "StopConfd",
			Handler:    _PilotService_StopConfd_Handler,
		},
		{
			MethodName: "RegisterMetadata",
			Handler:    _PilotService_RegisterMetadata_Handler,
		},
		{
			MethodName: "DeregisterMetadata",
			Handler:    _PilotService_DeregisterMetadata_Handler,
		},
		{
			MethodName: "RegisterCmd",
			Handler:    _PilotService_RegisterCmd_Handler,
		},
		{
			MethodName: "DeregisterCmd",
			Handler:    _PilotService_DeregisterCmd_Handler,
		},
		{
			MethodName: "ReportSubTaskStatus",
			Handler:    _PilotService_ReportSubTaskStatus_Handler,
		},
		{
			MethodName: "GetSubtaskStatus",
			Handler:    _PilotService_GetSubtaskStatus_Handler,
		},
		{
			MethodName: "HandleSubtask",
			Handler:    _PilotService_HandleSubtask_Handler,
		},
		{
			MethodName: "PingPilot",
			Handler:    _PilotService_PingPilot_Handler,
		},
		{
			MethodName: "PingFrontgate",
			Handler:    _PilotService_PingFrontgate_Handler,
		},
		{
			MethodName: "PingDrone",
			Handler:    _PilotService_PingDrone_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FrontgateChannel",
			Handler:       _PilotService_FrontgateChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "openpitrix/pilot/pilot.proto",
}

func init() { proto.RegisterFile("openpitrix/pilot/pilot.proto", fileDescriptor_pilot_75c5e83052e5fc3e) }

var fileDescriptor_pilot_75c5e83052e5fc3e = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x51, 0x6f, 0xd3, 0x3e,
	0x14, 0xc5, 0xb5, 0x97, 0xbf, 0xb4, 0xfb, 0x5f, 0xab, 0xca, 0x48, 0x20, 0x65, 0x9b, 0x18, 0x0f,
	0x68, 0x08, 0xad, 0x0d, 0x82, 0x47, 0xc4, 0x4b, 0xbb, 0xae, 0x2d, 0xa2, 0x50, 0x1a, 0x04, 0x12,
	0x3c, 0x20, 0x77, 0xb9, 0x33, 0x56, 0x5b, 0xdb, 0xc4, 0xb7, 0x40, 0xbf, 0x16, 0x9f, 0x10, 0xc5,
	0xc9, 0xd6, 0x74, 0x89, 0x9b, 0x49, 0x7b, 0x49, 0xa4, 0x9c, 0x73, 0x7e, 0xbe, 0xb6, 0xe3, 0x6b,
	0x38, 0xd2, 0x06, 0x95, 0x91, 0x94, 0xc8, 0x3f, 0xa1, 0x91, 0x0b, 0x4d, 0xd9, 0xb3, 0x63, 0x12,
	0x4d, 0x9a, 0xb5, 0x36, 0x6a, 0xc7, 0x7d, 0x0f, 0x8e, 0x84, 0xd6, 0x62, 0x81, 0x21, 0x37, 0x32,
	0xe4, 0x4a, 0x69, 0xe2, 0x24, 0xb5, 0xb2, 0x99, 0x3f, 0x38, 0x73, 0xaf, 0xcb, 0xb6, 0x40, 0xd5,
	0xb6, 0xbf, 0xb9, 0x10, 0x98, 0x84, 0xda, 0x38, 0x47, 0x85, 0xbb, 0x38, 0x36, 0xad, 0x0d, 0xda,
	0xec, 0xe9, 0x55, 0xe3, 0x44, 0x2b, 0xcc, 0xd5, 0x93, 0x92, 0x7a, 0x95, 0x68, 0x45, 0x82, 0x13,
	0x7a, 0xf3, 0x85, 0x99, 0x05, 0x87, 0xe5, 0xb1, 0xb9, 0x9d, 0x67, 0xe2, 0xcb, 0xbf, 0x07, 0x70,
	0x30, 0x49, 0xcd, 0x11, 0x26, 0xbf, 0xe4, 0x25, 0xb2, 0x21, 0x34, 0x07, 0x48, 0xee, 0x53, 0x4f,
	0xab, 0x2b, 0x29, 0xd8, 0xa3, 0x4e, 0x61, 0x69, 0xb2, 0xb2, 0xfb, 0x4b, 0x43, 0xeb, 0xe0, 0xb8,
	0x2c, 0x14, 0x73, 0xef, 0xa1, 0x35, 0x40, 0xba, 0xb8, 0xae, 0xf5, 0x9d, 0xb4, 0xe4, 0x67, 0x3d,
	0x29, 0x0b, 0x37, 0xc9, 0x51, 0xec, 0xb2, 0x9f, 0x81, 0x15, 0x79, 0xf9, 0x28, 0xc7, 0x3b, 0x83,
	0x3b, 0xb9, 0x39, 0x61, 0x02, 0x2c, 0x2a, 0x73, 0xeb, 0x83, 0x81, 0x6f, 0x32, 0x6c, 0xe2, 0xd6,
	0xf0, 0x3c, 0xdd, 0xc3, 0x9c, 0xf6, 0xb8, 0x6c, 0x75, 0x72, 0x5f, 0xc5, 0x46, 0x4b, 0x45, 0x55,
	0x6b, 0x59, 0xcc, 0x4f, 0xa1, 0x19, 0x6d, 0x13, 0x4f, 0xcb, 0x81, 0x6d, 0xc7, 0x14, 0x7f, 0xae,
	0xd0, 0x92, 0xbf, 0xca, 0x11, 0x34, 0x47, 0x36, 0xf5, 0xc6, 0xd3, 0x95, 0x52, 0x52, 0xdd, 0xa1,
	0xca, 0x87, 0x65, 0x43, 0x57, 0xeb, 0x05, 0xbb, 0x00, 0x88, 0x88, 0x27, 0x6e, 0xe7, 0xe3, 0x7a,
	0x8c, 0xb7, 0xa4, 0x3e, 0xec, 0x47, 0xa4, 0xcd, 0x7d, 0x31, 0x5f, 0xa0, 0x35, 0x45, 0x21, 0x2d,
	0x61, 0x32, 0x46, 0xe2, 0x31, 0x27, 0xce, 0x9e, 0x57, 0xac, 0xd7, 0x6a, 0xf6, 0x89, 0xdb, 0xf9,
	0xf7, 0xdb, 0x5e, 0x3f, 0xf8, 0x1b, 0xb0, 0x73, 0x4c, 0x6e, 0xa3, 0xcf, 0xfc, 0xe8, 0xb2, 0xdb,
	0x0f, 0x1f, 0xc3, 0xff, 0xd7, 0x95, 0xf4, 0x96, 0x31, 0x7b, 0x5a, 0x5f, 0x70, 0x6f, 0x19, 0xfb,
	0x71, 0x1f, 0xa1, 0xb1, 0x19, 0x3d, 0x05, 0x9e, 0xde, 0xa5, 0xcc, 0x9d, 0xc8, 0x0f, 0xf0, 0x60,
	0x8a, 0x46, 0x27, 0x94, 0xe7, 0x22, 0xe2, 0xb4, 0xb2, 0x55, 0x1b, 0xb5, 0x65, 0xd8, 0x75, 0x50,
	0xd2, 0x16, 0x11, 0xad, 0x66, 0xb4, 0xa1, 0x1d, 0x7a, 0x69, 0xa3, 0x38, 0xa8, 0x1b, 0x8a, 0xbd,
	0x85, 0xc6, 0x90, 0xab, 0x78, 0x81, 0x39, 0x94, 0x9d, 0x78, 0x13, 0x63, 0xb4, 0x96, 0x0b, 0xf4,
	0x57, 0xf7, 0x06, 0xf6, 0x27, 0x52, 0x09, 0xd7, 0xd3, 0xfc, 0x9d, 0xcb, 0x1b, 0x1f, 0x40, 0x23,
	0x8d, 0xdf, 0x74, 0x8d, 0xba, 0x56, 0xb5, 0xeb, 0x54, 0xa4, 0x20, 0xf7, 0xf3, 0xdf, 0xe3, 0x54,
	0x0c, 0xa1, 0xb5, 0xe9, 0x60, 0x3f, 0xb8, 0x52, 0xb8, 0xa8, 0x9a, 0x55, 0x77, 0x4d, 0x58, 0xb9,
	0x65, 0x4e, 0x78, 0xb6, 0xf7, 0x62, 0xaf, 0x1b, 0x7e, 0x6d, 0x17, 0x54, 0xa9, 0xc3, 0xe2, 0xcd,
	0x3a, 0x17, 0xa1, 0x99, 0x65, 0x17, 0xd0, 0x6b, 0x33, 0x73, 0xef, 0xd9, 0x7f, 0xee, 0xb2, 0x79,
	0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x09, 0x30, 0x1d, 0x76, 0x83, 0x07, 0x00, 0x00,
}

// Copyright © 2025 Prabhjot Singh Sethi, All Rights reserved
// Author: Prabhjot Singh Sethi <prabhjot.sethi@gmail.com>

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: org-unit.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrgUnit_ListOrgUnits_FullMethodName  = "/api.OrgUnit/ListOrgUnits"
	OrgUnit_CreateOrgUnit_FullMethodName = "/api.OrgUnit/CreateOrgUnit"
	OrgUnit_UpdateOrgUnit_FullMethodName = "/api.OrgUnit/UpdateOrgUnit"
	OrgUnit_GetOrgUnit_FullMethodName    = "/api.OrgUnit/GetOrgUnit"
	OrgUnit_DeleteOrgUnit_FullMethodName = "/api.OrgUnit/DeleteOrgUnit"
)

// OrgUnitClient is the client API for OrgUnit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service provided to manage orgunits for user/tenant
type OrgUnitClient interface {
	// Get List of existing org units in my tenant
	ListOrgUnits(ctx context.Context, in *OrgUnitsListReq, opts ...grpc.CallOption) (*OrgUnitsListResp, error)
	// Create new org unit for my tenant
	CreateOrgUnit(ctx context.Context, in *OrgUnitCreateReq, opts ...grpc.CallOption) (*OrgUnitCreateResp, error)
	// Update an existing org unit for my tenant
	UpdateOrgUnit(ctx context.Context, in *OrgUnitUpdateReq, opts ...grpc.CallOption) (*OrgUnitUpdateResp, error)
	// get an existing org unit for my tenant
	GetOrgUnit(ctx context.Context, in *OrgUnitGetReq, opts ...grpc.CallOption) (*OrgUnitGetResp, error)
	// delete an existing org unit for my tenant
	DeleteOrgUnit(ctx context.Context, in *OrgUnitDeleteReq, opts ...grpc.CallOption) (*OrgUnitDeleteResp, error)
}

type orgUnitClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgUnitClient(cc grpc.ClientConnInterface) OrgUnitClient {
	return &orgUnitClient{cc}
}

func (c *orgUnitClient) ListOrgUnits(ctx context.Context, in *OrgUnitsListReq, opts ...grpc.CallOption) (*OrgUnitsListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrgUnitsListResp)
	err := c.cc.Invoke(ctx, OrgUnit_ListOrgUnits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgUnitClient) CreateOrgUnit(ctx context.Context, in *OrgUnitCreateReq, opts ...grpc.CallOption) (*OrgUnitCreateResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrgUnitCreateResp)
	err := c.cc.Invoke(ctx, OrgUnit_CreateOrgUnit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgUnitClient) UpdateOrgUnit(ctx context.Context, in *OrgUnitUpdateReq, opts ...grpc.CallOption) (*OrgUnitUpdateResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrgUnitUpdateResp)
	err := c.cc.Invoke(ctx, OrgUnit_UpdateOrgUnit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgUnitClient) GetOrgUnit(ctx context.Context, in *OrgUnitGetReq, opts ...grpc.CallOption) (*OrgUnitGetResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrgUnitGetResp)
	err := c.cc.Invoke(ctx, OrgUnit_GetOrgUnit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgUnitClient) DeleteOrgUnit(ctx context.Context, in *OrgUnitDeleteReq, opts ...grpc.CallOption) (*OrgUnitDeleteResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrgUnitDeleteResp)
	err := c.cc.Invoke(ctx, OrgUnit_DeleteOrgUnit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgUnitServer is the server API for OrgUnit service.
// All implementations must embed UnimplementedOrgUnitServer
// for forward compatibility.
//
// Service provided to manage orgunits for user/tenant
type OrgUnitServer interface {
	// Get List of existing org units in my tenant
	ListOrgUnits(context.Context, *OrgUnitsListReq) (*OrgUnitsListResp, error)
	// Create new org unit for my tenant
	CreateOrgUnit(context.Context, *OrgUnitCreateReq) (*OrgUnitCreateResp, error)
	// Update an existing org unit for my tenant
	UpdateOrgUnit(context.Context, *OrgUnitUpdateReq) (*OrgUnitUpdateResp, error)
	// get an existing org unit for my tenant
	GetOrgUnit(context.Context, *OrgUnitGetReq) (*OrgUnitGetResp, error)
	// delete an existing org unit for my tenant
	DeleteOrgUnit(context.Context, *OrgUnitDeleteReq) (*OrgUnitDeleteResp, error)
	mustEmbedUnimplementedOrgUnitServer()
}

// UnimplementedOrgUnitServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrgUnitServer struct{}

func (UnimplementedOrgUnitServer) ListOrgUnits(context.Context, *OrgUnitsListReq) (*OrgUnitsListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrgUnits not implemented")
}
func (UnimplementedOrgUnitServer) CreateOrgUnit(context.Context, *OrgUnitCreateReq) (*OrgUnitCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrgUnit not implemented")
}
func (UnimplementedOrgUnitServer) UpdateOrgUnit(context.Context, *OrgUnitUpdateReq) (*OrgUnitUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrgUnit not implemented")
}
func (UnimplementedOrgUnitServer) GetOrgUnit(context.Context, *OrgUnitGetReq) (*OrgUnitGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrgUnit not implemented")
}
func (UnimplementedOrgUnitServer) DeleteOrgUnit(context.Context, *OrgUnitDeleteReq) (*OrgUnitDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrgUnit not implemented")
}
func (UnimplementedOrgUnitServer) mustEmbedUnimplementedOrgUnitServer() {}
func (UnimplementedOrgUnitServer) testEmbeddedByValue()                 {}

// UnsafeOrgUnitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgUnitServer will
// result in compilation errors.
type UnsafeOrgUnitServer interface {
	mustEmbedUnimplementedOrgUnitServer()
}

func RegisterOrgUnitServer(s grpc.ServiceRegistrar, srv OrgUnitServer) {
	// If the following call pancis, it indicates UnimplementedOrgUnitServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrgUnit_ServiceDesc, srv)
}

func _OrgUnit_ListOrgUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgUnitsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgUnitServer).ListOrgUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrgUnit_ListOrgUnits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgUnitServer).ListOrgUnits(ctx, req.(*OrgUnitsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgUnit_CreateOrgUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgUnitCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgUnitServer).CreateOrgUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrgUnit_CreateOrgUnit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgUnitServer).CreateOrgUnit(ctx, req.(*OrgUnitCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgUnit_UpdateOrgUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgUnitUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgUnitServer).UpdateOrgUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrgUnit_UpdateOrgUnit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgUnitServer).UpdateOrgUnit(ctx, req.(*OrgUnitUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgUnit_GetOrgUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgUnitGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgUnitServer).GetOrgUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrgUnit_GetOrgUnit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgUnitServer).GetOrgUnit(ctx, req.(*OrgUnitGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgUnit_DeleteOrgUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgUnitDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgUnitServer).DeleteOrgUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrgUnit_DeleteOrgUnit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgUnitServer).DeleteOrgUnit(ctx, req.(*OrgUnitDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrgUnit_ServiceDesc is the grpc.ServiceDesc for OrgUnit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrgUnit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.OrgUnit",
	HandlerType: (*OrgUnitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOrgUnits",
			Handler:    _OrgUnit_ListOrgUnits_Handler,
		},
		{
			MethodName: "CreateOrgUnit",
			Handler:    _OrgUnit_CreateOrgUnit_Handler,
		},
		{
			MethodName: "UpdateOrgUnit",
			Handler:    _OrgUnit_UpdateOrgUnit_Handler,
		},
		{
			MethodName: "GetOrgUnit",
			Handler:    _OrgUnit_GetOrgUnit_Handler,
		},
		{
			MethodName: "DeleteOrgUnit",
			Handler:    _OrgUnit_DeleteOrgUnit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "org-unit.proto",
}

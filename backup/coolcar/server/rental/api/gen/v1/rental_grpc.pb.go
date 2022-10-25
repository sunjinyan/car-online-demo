// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: rental.proto

package rentalpb

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

// TripServiceClient is the client API for TripService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TripServiceClient interface {
	CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*TripEntity, error)
	GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*Trip, error)
	GetTrips(ctx context.Context, in *GetTripsRequest, opts ...grpc.CallOption) (*GetTripsResponse, error)
	UpdateTrip(ctx context.Context, in *UpdateTripsRequest, opts ...grpc.CallOption) (*Trip, error)
}

type tripServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTripServiceClient(cc grpc.ClientConnInterface) TripServiceClient {
	return &tripServiceClient{cc}
}

func (c *tripServiceClient) CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*TripEntity, error) {
	out := new(TripEntity)
	err := c.cc.Invoke(ctx, "/rental.v1.TripService/CreateTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*Trip, error) {
	out := new(Trip)
	err := c.cc.Invoke(ctx, "/rental.v1.TripService/GetTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) GetTrips(ctx context.Context, in *GetTripsRequest, opts ...grpc.CallOption) (*GetTripsResponse, error) {
	out := new(GetTripsResponse)
	err := c.cc.Invoke(ctx, "/rental.v1.TripService/GetTrips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) UpdateTrip(ctx context.Context, in *UpdateTripsRequest, opts ...grpc.CallOption) (*Trip, error) {
	out := new(Trip)
	err := c.cc.Invoke(ctx, "/rental.v1.TripService/UpdateTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TripServiceServer is the server API for TripService service.
// All implementations must embed UnimplementedTripServiceServer
// for forward compatibility
type TripServiceServer interface {
	CreateTrip(context.Context, *CreateTripRequest) (*TripEntity, error)
	GetTrip(context.Context, *GetTripRequest) (*Trip, error)
	GetTrips(context.Context, *GetTripsRequest) (*GetTripsResponse, error)
	UpdateTrip(context.Context, *UpdateTripsRequest) (*Trip, error)
	mustEmbedUnimplementedTripServiceServer()
}

// UnimplementedTripServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTripServiceServer struct {
}

func (UnimplementedTripServiceServer) CreateTrip(context.Context, *CreateTripRequest) (*TripEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}
func (UnimplementedTripServiceServer) GetTrip(context.Context, *GetTripRequest) (*Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (UnimplementedTripServiceServer) GetTrips(context.Context, *GetTripsRequest) (*GetTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrips not implemented")
}
func (UnimplementedTripServiceServer) UpdateTrip(context.Context, *UpdateTripsRequest) (*Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}
func (UnimplementedTripServiceServer) mustEmbedUnimplementedTripServiceServer() {}

// UnsafeTripServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TripServiceServer will
// result in compilation errors.
type UnsafeTripServiceServer interface {
	mustEmbedUnimplementedTripServiceServer()
}

func RegisterTripServiceServer(s grpc.ServiceRegistrar, srv TripServiceServer) {
	s.RegisterService(&TripService_ServiceDesc, srv)
}

func _TripService_CreateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).CreateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.TripService/CreateTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).CreateTrip(ctx, req.(*CreateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.TripService/GetTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTrip(ctx, req.(*GetTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_GetTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.TripService/GetTrips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTrips(ctx, req.(*GetTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_UpdateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).UpdateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.TripService/UpdateTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).UpdateTrip(ctx, req.(*UpdateTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TripService_ServiceDesc is the grpc.ServiceDesc for TripService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TripService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rental.v1.TripService",
	HandlerType: (*TripServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrip",
			Handler:    _TripService_CreateTrip_Handler,
		},
		{
			MethodName: "GetTrip",
			Handler:    _TripService_GetTrip_Handler,
		},
		{
			MethodName: "GetTrips",
			Handler:    _TripService_GetTrips_Handler,
		},
		{
			MethodName: "UpdateTrip",
			Handler:    _TripService_UpdateTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rental.proto",
}

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*Profile, error)
	SubmitProfile(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Profile, error)
	ClearProfile(ctx context.Context, in *ClearProfileRequest, opts ...grpc.CallOption) (*Profile, error)
	GetProfilePhoto(ctx context.Context, in *GetProfilePhotoRequest, opts ...grpc.CallOption) (*GetProfilePhotoResponse, error)
	CreateProfilePhoto(ctx context.Context, in *CreateProfilePhotoRequest, opts ...grpc.CallOption) (*CreateProfilePhotoResponse, error)
	CompleteProfilePhoto(ctx context.Context, in *CompleteProfilePhotoRequest, opts ...grpc.CallOption) (*Identity, error)
	ClearProfilePhoto(ctx context.Context, in *ClearProfilePhotoRequest, opts ...grpc.CallOption) (*ClearProfilePhotoResponse, error)
	UploadFilePhoto(ctx context.Context, in *UploadFilePhotoRequest, opts ...grpc.CallOption) (*UploadFilePhotoResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/rental.v1.ProfileService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) SubmitProfile(ctx context.Context, in *Identity, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/rental.v1.ProfileService/SubmitProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) ClearProfile(ctx context.Context, in *ClearProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/rental.v1.ProfileService/ClearProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) GetProfilePhoto(ctx context.Context, in *GetProfilePhotoRequest, opts ...grpc.CallOption) (*GetProfilePhotoResponse, error) {
	out := new(GetProfilePhotoResponse)
	err := c.cc.Invoke(ctx, "/rental.v1.ProfileService/GetProfilePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) CreateProfilePhoto(ctx context.Context, in *CreateProfilePhotoRequest, opts ...grpc.CallOption) (*CreateProfilePhotoResponse, error) {
	out := new(CreateProfilePhotoResponse)
	err := c.cc.Invoke(ctx, "/rental.v1.ProfileService/CreateProfilePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) CompleteProfilePhoto(ctx context.Context, in *CompleteProfilePhotoRequest, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, "/rental.v1.ProfileService/CompleteProfilePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) ClearProfilePhoto(ctx context.Context, in *ClearProfilePhotoRequest, opts ...grpc.CallOption) (*ClearProfilePhotoResponse, error) {
	out := new(ClearProfilePhotoResponse)
	err := c.cc.Invoke(ctx, "/rental.v1.ProfileService/ClearProfilePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) UploadFilePhoto(ctx context.Context, in *UploadFilePhotoRequest, opts ...grpc.CallOption) (*UploadFilePhotoResponse, error) {
	out := new(UploadFilePhotoResponse)
	err := c.cc.Invoke(ctx, "/rental.v1.ProfileService/UploadFilePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	GetProfile(context.Context, *GetProfileRequest) (*Profile, error)
	SubmitProfile(context.Context, *Identity) (*Profile, error)
	ClearProfile(context.Context, *ClearProfileRequest) (*Profile, error)
	GetProfilePhoto(context.Context, *GetProfilePhotoRequest) (*GetProfilePhotoResponse, error)
	CreateProfilePhoto(context.Context, *CreateProfilePhotoRequest) (*CreateProfilePhotoResponse, error)
	CompleteProfilePhoto(context.Context, *CompleteProfilePhotoRequest) (*Identity, error)
	ClearProfilePhoto(context.Context, *ClearProfilePhotoRequest) (*ClearProfilePhotoResponse, error)
	UploadFilePhoto(context.Context, *UploadFilePhotoRequest) (*UploadFilePhotoResponse, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) GetProfile(context.Context, *GetProfileRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedProfileServiceServer) SubmitProfile(context.Context, *Identity) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProfile not implemented")
}
func (UnimplementedProfileServiceServer) ClearProfile(context.Context, *ClearProfileRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearProfile not implemented")
}
func (UnimplementedProfileServiceServer) GetProfilePhoto(context.Context, *GetProfilePhotoRequest) (*GetProfilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfilePhoto not implemented")
}
func (UnimplementedProfileServiceServer) CreateProfilePhoto(context.Context, *CreateProfilePhotoRequest) (*CreateProfilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfilePhoto not implemented")
}
func (UnimplementedProfileServiceServer) CompleteProfilePhoto(context.Context, *CompleteProfilePhotoRequest) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteProfilePhoto not implemented")
}
func (UnimplementedProfileServiceServer) ClearProfilePhoto(context.Context, *ClearProfilePhotoRequest) (*ClearProfilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearProfilePhoto not implemented")
}
func (UnimplementedProfileServiceServer) UploadFilePhoto(context.Context, *UploadFilePhotoRequest) (*UploadFilePhotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFilePhoto not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.ProfileService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_SubmitProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Identity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).SubmitProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.ProfileService/SubmitProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).SubmitProfile(ctx, req.(*Identity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_ClearProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).ClearProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.ProfileService/ClearProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).ClearProfile(ctx, req.(*ClearProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_GetProfilePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfilePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).GetProfilePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.ProfileService/GetProfilePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).GetProfilePhoto(ctx, req.(*GetProfilePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_CreateProfilePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfilePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).CreateProfilePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.ProfileService/CreateProfilePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).CreateProfilePhoto(ctx, req.(*CreateProfilePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_CompleteProfilePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteProfilePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).CompleteProfilePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.ProfileService/CompleteProfilePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).CompleteProfilePhoto(ctx, req.(*CompleteProfilePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_ClearProfilePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearProfilePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).ClearProfilePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.ProfileService/ClearProfilePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).ClearProfilePhoto(ctx, req.(*ClearProfilePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_UploadFilePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFilePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).UploadFilePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.v1.ProfileService/UploadFilePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).UploadFilePhoto(ctx, req.(*UploadFilePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rental.v1.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _ProfileService_GetProfile_Handler,
		},
		{
			MethodName: "SubmitProfile",
			Handler:    _ProfileService_SubmitProfile_Handler,
		},
		{
			MethodName: "ClearProfile",
			Handler:    _ProfileService_ClearProfile_Handler,
		},
		{
			MethodName: "GetProfilePhoto",
			Handler:    _ProfileService_GetProfilePhoto_Handler,
		},
		{
			MethodName: "CreateProfilePhoto",
			Handler:    _ProfileService_CreateProfilePhoto_Handler,
		},
		{
			MethodName: "CompleteProfilePhoto",
			Handler:    _ProfileService_CompleteProfilePhoto_Handler,
		},
		{
			MethodName: "ClearProfilePhoto",
			Handler:    _ProfileService_ClearProfilePhoto_Handler,
		},
		{
			MethodName: "UploadFilePhoto",
			Handler:    _ProfileService_UploadFilePhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rental.proto",
}
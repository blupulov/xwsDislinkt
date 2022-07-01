// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: services/user-service/user-service.proto

package user_service

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	AddSkill(ctx context.Context, in *AddSkillRequest, opts ...grpc.CallOption) (*AddSkillResponse, error)
	AddExpirience(ctx context.Context, in *AddExpirienceRequest, opts ...grpc.CallOption) (*AddExpirienceResponse, error)
	AddEducation(ctx context.Context, in *AddEducationRequest, opts ...grpc.CallOption) (*AddEducationResponse, error)
	AddInterest(ctx context.Context, in *AddInterestRequest, opts ...grpc.CallOption) (*AddInterestResponse, error)
	//VIDETI KASNIJE DA LI MOZE NETO DA SE ISKOMBINUJE POSTO GET METODA NEMA BODY
	GetManyUsersById(ctx context.Context, in *GetManyUsersByIdRequest, opts ...grpc.CallOption) (*GetManyUsersByIdResponse, error)
	GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*GetUserByUsernameResponse, error)
	ChangeUser(ctx context.Context, in *ChangeUserRequest, opts ...grpc.CallOption) (*ChangeUserResponse, error)
	PromoteUserToCompanyOwner(ctx context.Context, in *PromoteUserRequest, opts ...grpc.CallOption) (*PromoteUserResponse, error)
	CreateApiToken(ctx context.Context, in *CreateApiTokenRequest, opts ...grpc.CallOption) (*CreateApiTokenResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddSkill(ctx context.Context, in *AddSkillRequest, opts ...grpc.CallOption) (*AddSkillResponse, error) {
	out := new(AddSkillResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/AddSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddExpirience(ctx context.Context, in *AddExpirienceRequest, opts ...grpc.CallOption) (*AddExpirienceResponse, error) {
	out := new(AddExpirienceResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/AddExpirience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddEducation(ctx context.Context, in *AddEducationRequest, opts ...grpc.CallOption) (*AddEducationResponse, error) {
	out := new(AddEducationResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/AddEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddInterest(ctx context.Context, in *AddInterestRequest, opts ...grpc.CallOption) (*AddInterestResponse, error) {
	out := new(AddInterestResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/AddInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetManyUsersById(ctx context.Context, in *GetManyUsersByIdRequest, opts ...grpc.CallOption) (*GetManyUsersByIdResponse, error) {
	out := new(GetManyUsersByIdResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/GetManyUsersById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*GetUserByUsernameResponse, error) {
	out := new(GetUserByUsernameResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/GetUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeUser(ctx context.Context, in *ChangeUserRequest, opts ...grpc.CallOption) (*ChangeUserResponse, error) {
	out := new(ChangeUserResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/ChangeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PromoteUserToCompanyOwner(ctx context.Context, in *PromoteUserRequest, opts ...grpc.CallOption) (*PromoteUserResponse, error) {
	out := new(PromoteUserResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/PromoteUserToCompanyOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateApiToken(ctx context.Context, in *CreateApiTokenRequest, opts ...grpc.CallOption) (*CreateApiTokenResponse, error) {
	out := new(CreateApiTokenResponse)
	err := c.cc.Invoke(ctx, "/userservice.UserService/CreateApiToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	AddSkill(context.Context, *AddSkillRequest) (*AddSkillResponse, error)
	AddExpirience(context.Context, *AddExpirienceRequest) (*AddExpirienceResponse, error)
	AddEducation(context.Context, *AddEducationRequest) (*AddEducationResponse, error)
	AddInterest(context.Context, *AddInterestRequest) (*AddInterestResponse, error)
	//VIDETI KASNIJE DA LI MOZE NETO DA SE ISKOMBINUJE POSTO GET METODA NEMA BODY
	GetManyUsersById(context.Context, *GetManyUsersByIdRequest) (*GetManyUsersByIdResponse, error)
	GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*GetUserByUsernameResponse, error)
	ChangeUser(context.Context, *ChangeUserRequest) (*ChangeUserResponse, error)
	PromoteUserToCompanyOwner(context.Context, *PromoteUserRequest) (*PromoteUserResponse, error)
	CreateApiToken(context.Context, *CreateApiTokenRequest) (*CreateApiTokenResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) AddSkill(context.Context, *AddSkillRequest) (*AddSkillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSkill not implemented")
}
func (UnimplementedUserServiceServer) AddExpirience(context.Context, *AddExpirienceRequest) (*AddExpirienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExpirience not implemented")
}
func (UnimplementedUserServiceServer) AddEducation(context.Context, *AddEducationRequest) (*AddEducationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEducation not implemented")
}
func (UnimplementedUserServiceServer) AddInterest(context.Context, *AddInterestRequest) (*AddInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInterest not implemented")
}
func (UnimplementedUserServiceServer) GetManyUsersById(context.Context, *GetManyUsersByIdRequest) (*GetManyUsersByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyUsersById not implemented")
}
func (UnimplementedUserServiceServer) GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*GetUserByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUserServiceServer) ChangeUser(context.Context, *ChangeUserRequest) (*ChangeUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUser not implemented")
}
func (UnimplementedUserServiceServer) PromoteUserToCompanyOwner(context.Context, *PromoteUserRequest) (*PromoteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteUserToCompanyOwner not implemented")
}
func (UnimplementedUserServiceServer) CreateApiToken(context.Context, *CreateApiTokenRequest) (*CreateApiTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApiToken not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/AddSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddSkill(ctx, req.(*AddSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddExpirience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExpirienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddExpirience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/AddExpirience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddExpirience(ctx, req.(*AddExpirienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEducationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/AddEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddEducation(ctx, req.(*AddEducationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/AddInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddInterest(ctx, req.(*AddInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetManyUsersById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyUsersByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetManyUsersById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/GetManyUsersById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetManyUsersById(ctx, req.(*GetManyUsersByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/GetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByUsername(ctx, req.(*GetUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/ChangeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeUser(ctx, req.(*ChangeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PromoteUserToCompanyOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PromoteUserToCompanyOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/PromoteUserToCompanyOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PromoteUserToCompanyOwner(ctx, req.(*PromoteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateApiToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateApiToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/CreateApiToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateApiToken(ctx, req.(*CreateApiTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userservice.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _UserService_GetById_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "AddSkill",
			Handler:    _UserService_AddSkill_Handler,
		},
		{
			MethodName: "AddExpirience",
			Handler:    _UserService_AddExpirience_Handler,
		},
		{
			MethodName: "AddEducation",
			Handler:    _UserService_AddEducation_Handler,
		},
		{
			MethodName: "AddInterest",
			Handler:    _UserService_AddInterest_Handler,
		},
		{
			MethodName: "GetManyUsersById",
			Handler:    _UserService_GetManyUsersById_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _UserService_GetUserByUsername_Handler,
		},
		{
			MethodName: "ChangeUser",
			Handler:    _UserService_ChangeUser_Handler,
		},
		{
			MethodName: "PromoteUserToCompanyOwner",
			Handler:    _UserService_PromoteUserToCompanyOwner_Handler,
		},
		{
			MethodName: "CreateApiToken",
			Handler:    _UserService_CreateApiToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/user-service/user-service.proto",
}

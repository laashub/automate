// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: api/interservice/local_user/users.proto

package local_user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the UsersMgmtServer interface (at compile time)
var _ UsersMgmtServer = &UsersMgmtServerMock{}

// NewUsersMgmtServerMock gives you a fresh instance of UsersMgmtServerMock.
func NewUsersMgmtServerMock() *UsersMgmtServerMock {
	return &UsersMgmtServerMock{validateRequests: true}
}

// NewUsersMgmtServerMockWithoutValidation gives you a fresh instance of
// UsersMgmtServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewUsersMgmtServerMockWithoutValidation() *UsersMgmtServerMock {
	return &UsersMgmtServerMock{}
}

// UsersMgmtServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type UsersMgmtServerMock struct {
	validateRequests bool
	GetUsersFunc     func(context.Context, *GetUsersReq) (*Users, error)
	GetUserFunc      func(context.Context, *Email) (*User, error)
	CreateUserFunc   func(context.Context, *CreateUserReq) (*User, error)
	DeleteUserFunc   func(context.Context, *Email) (*DeleteUserResp, error)
	UpdateUserFunc   func(context.Context, *UpdateUserReq) (*User, error)
	UpdateSelfFunc   func(context.Context, *UpdateSelfReq) (*User, error)
}

func (m *UsersMgmtServerMock) GetUsers(ctx context.Context, req *GetUsersReq) (*Users, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetUsersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetUsers' not implemented")
}

func (m *UsersMgmtServerMock) GetUser(ctx context.Context, req *Email) (*User, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetUserFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetUser' not implemented")
}

func (m *UsersMgmtServerMock) CreateUser(ctx context.Context, req *CreateUserReq) (*User, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateUserFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateUser' not implemented")
}

func (m *UsersMgmtServerMock) DeleteUser(ctx context.Context, req *Email) (*DeleteUserResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteUserFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteUser' not implemented")
}

func (m *UsersMgmtServerMock) UpdateUser(ctx context.Context, req *UpdateUserReq) (*User, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateUserFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateUser' not implemented")
}

func (m *UsersMgmtServerMock) UpdateSelf(ctx context.Context, req *UpdateSelfReq) (*User, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateSelfFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateSelf' not implemented")
}

// Reset resets all overridden functions
func (m *UsersMgmtServerMock) Reset() {
	m.GetUsersFunc = nil
	m.GetUserFunc = nil
	m.CreateUserFunc = nil
	m.DeleteUserFunc = nil
	m.UpdateUserFunc = nil
	m.UpdateSelfFunc = nil
}

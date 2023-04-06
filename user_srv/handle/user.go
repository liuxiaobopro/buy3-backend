package handle

import (
	"context"
	"fmt"

	"user_srv/proto"
)

type UserHandle struct {
	*proto.UnimplementedUserServer
}

var _ proto.UserServer = (*UserHandle)(nil)

// CreateUser implements proto.UserServer
func (*UserHandle) CreateUser(c context.Context, in *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	fmt.Printf("in: %v", in)
	return &proto.CreateUserResponse{}, nil
}

// GetUser implements proto.UserServer
func (*UserHandle) GetUser(context.Context, *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	panic("unimplemented")
}

// UpdateUser implements proto.UserServer
func (*UserHandle) UpdateUser(context.Context, *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	panic("unimplemented")
}

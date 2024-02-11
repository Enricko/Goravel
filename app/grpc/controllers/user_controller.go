package controllers

import (
	"context"
	"net/http"

	proto "github.com/goravel/example-proto"
)

type UserController struct {
}

// GetUser implements user.UserServiceServer.
func (*UserController) GetUser(context.Context, *proto.UserRequest) (*proto.UserResponse, error) {
	panic("unimplemented")
}

func NewUserController() *UserController {
	return &UserController{}
}

func (r *UserController) Show(ctx context.Context, req *proto.UserRequest) (protoBook *proto.UserResponse, err error) {
	return &proto.UserResponse{
		Code: http.StatusOK,
	}, nil
}

package user

import (
	"context"
	pb "github.com/faith2333/xuanwu/api/user/v1"
	biz "github.com/faith2333/xuanwu/internal/biz/user"
)

type ServiceUser struct {
	pb.UnimplementedUserServerServer
	biz biz.Interface
}

func NewServiceUser(bizUser biz.Interface) *ServiceUser {
	return &ServiceUser{
		biz: bizUser,
	}
}

func (svc *ServiceUser) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.EmptyResponse, error) {
	bizUser := &biz.User{
		Username:    request.Username,
		Password:    request.Password,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		ExtraInfo:   request.ExtraInfo.AsMap(),
	}
	return nil, svc.biz.SignUp(ctx, bizUser)
}

func (svc *ServiceUser) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := svc.biz.Login(ctx, request.GetUsername(), request.GetPassword())
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		JwtToken: token,
	}, nil
}

package user

import (
	"context"
	pb "github.com/faith2333/xuanwu/api/user/v1"
	biz "github.com/faith2333/xuanwu/internal/biz/user"
	"github.com/faith2333/xuanwu/internal/service/base"
)

type ServiceUser struct {
	pb.UnimplementedUserServerServer
	base.Base
	biz biz.Interface
}

func NewServiceUser(bizUser biz.Interface) *ServiceUser {
	return &ServiceUser{
		biz: bizUser,
	}
}

func (svc *ServiceUser) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.EmptyReply, error) {
	bizUser := &biz.User{
		Username:    request.Username,
		Password:    request.Password,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		ExtraInfo:   request.ExtraInfo.AsMap(),
	}
	return nil, svc.biz.SignUp(ctx, bizUser)
}

func (svc *ServiceUser) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := svc.biz.Login(ctx, request.GetUsername(), request.GetPassword())
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		JwtToken: token,
	}, nil
}

func (svc *ServiceUser) GetCurrentUser(ctx context.Context, request *pb.EmptyRequest) (*pb.GetCurrentUserReply, error) {
	user, err := svc.biz.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetCurrentUserReply{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (svc *ServiceUser) Logout(ctx context.Context, request *pb.EmptyRequest) (*pb.EmptyReply, error) {
	// todo jwt logout
	return nil, nil
}

func (svc *ServiceUser) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.EmptyReply, error) {
	bizReq := &biz.ChangePasswordReq{}

	err := svc.TransformJSON(request, &bizReq)
	if err != nil {
		return nil, err
	}

	return nil, svc.biz.ChangePassword(ctx, bizReq)
}

func (svc *ServiceUser) GetUserByUsername(ctx context.Context, request *pb.GetUserByUsernameRequest) (*pb.User, error) {
	bizResp, err := svc.biz.GetUserByUsername(ctx, request.GetUsername())
	if err != nil {
		return nil, err
	}

	pbUser := &pb.User{}
	err = svc.TransformJSON(bizResp, &pbUser)
	if err != nil {
		return nil, err
	}

	return pbUser, nil
}

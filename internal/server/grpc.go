package server

import (
	"github.com/faith2333/xuanwu/internal/conf"
	selfJwt "github.com/faith2333/xuanwu/pkg/middleware/jwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	pbUser "github.com/faith2333/xuanwu/api/user/v1"
	svcUser "github.com/faith2333/xuanwu/internal/service/user"

	pbApp "github.com/faith2333/xuanwu/api/application/v1"
	svcApp "github.com/faith2333/xuanwu/internal/service/application"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, userSvc *svcUser.ServiceUser, appSvc *svcApp.AppSvc, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(
				jwt.Server(selfJwt.NewJWTFunc([]byte(c.Auth.JwtSecretKey)), jwt.WithSigningMethod(selfJwt.SigningMethod)),
			).Match(selfJwt.NewWhiteListMatcher(AuthenticationWhiteList)).Build(),
			selfJwt.AuthPlugin,
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)

	pbUser.RegisterUserServerServer(srv, userSvc)
	pbApp.RegisterApplicationSvcServer(srv, appSvc)
	return srv
}

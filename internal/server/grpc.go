package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	cicdPB "github/faith2333/xuanwu/api/cicd/v1"
	"github/faith2333/xuanwu/internal/conf"
	cicdSvc "github/faith2333/xuanwu/internal/service/cicd"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, cicd *cicdSvc.CICDService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
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
	cicdPB.RegisterCICDServer(srv, cicd)
	return srv
}

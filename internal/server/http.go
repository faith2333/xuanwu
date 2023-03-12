package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	cicdPB "github/faith2333/xuanwu/api/cicd/v1"
	"github/faith2333/xuanwu/internal/conf"
	cicdSvc "github/faith2333/xuanwu/internal/service/cicd"
	"github/faith2333/xuanwu/pkg/httpencoder"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, cicd *cicdSvc.CICDService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ErrorEncoder(httpencoder.ErrorEncoder),
		http.ResponseEncoder(httpencoder.ResponseEncoder),
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	cicdPB.RegisterCICDHTTPServer(srv, cicd)
	return srv
}

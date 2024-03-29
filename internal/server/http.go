package server

import (
	"github.com/faith2333/xuanwu/internal/conf"
	"github.com/faith2333/xuanwu/pkg/httpencoder"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"

	pbUser "github.com/faith2333/xuanwu/api/user/v1"
	svcUser "github.com/faith2333/xuanwu/internal/service/user"

	pbApp "github.com/faith2333/xuanwu/api/application/v1"
	svcApp "github.com/faith2333/xuanwu/internal/service/application"

	pbORG "github.com/faith2333/xuanwu/api/organization/v1"
	svcORG "github.com/faith2333/xuanwu/internal/service/organization"

	selfJwt "github.com/faith2333/xuanwu/pkg/middleware/jwt"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userSvc *svcUser.ServiceUser, appSvc *svcApp.AppSvc, orgSVC *svcORG.OrgSvc, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ErrorEncoder(httpencoder.ErrorEncoder),
		http.ResponseEncoder(httpencoder.ResponseEncoder),
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(
				jwt.Server(selfJwt.NewJWTFunc([]byte(c.Auth.JwtSecretKey)), jwt.WithSigningMethod(selfJwt.SigningMethod)),
			).Match(selfJwt.NewWhiteListMatcher(AuthenticationWhiteList)).Build(),
			selfJwt.AuthPlugin,
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
	pbUser.RegisterUserServerHTTPServer(srv, userSvc)
	pbApp.RegisterApplicationSvcHTTPServer(srv, appSvc)
	pbORG.RegisterOrganizationServiceHTTPServer(srv, orgSVC)

	return srv
}

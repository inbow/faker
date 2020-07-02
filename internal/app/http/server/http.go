package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	errch "github.com/proxeter/errors-channel"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"go.uber.org/zap"

	"github.com/oxyd-io/faker/internal/app/config"
	"github.com/oxyd-io/faker/internal/app/env"
	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	Server struct {
		ctx context.Context

		logger *zap.Logger
		config *config.AppConfig

		generator generator.IGenerator
	}

	HandlerResponse struct {
		StatusCode int
		Body       []byte
	}
)

func New(
	ctx context.Context,
	logger *zap.Logger,
	config *config.AppConfig,
	generator generator.IGenerator,
) <-chan error {
	return errch.Register(func() error {
		return (&Server{
			ctx: ctx,

			logger: logger,
			config: config,

			generator: generator,
		}).Start(ctx)
	})
}

func (s *Server) Start(ctx context.Context) error {
	server := atreugo.New(atreugo.Config{
		Name: s.ctx.Value(env.Name).(string) + " server",

		Prefork:   true,
		Reuseport: true,

		GracefulShutdown: true,
		Compress:         false,
	})

	ln, err := reuseport.Listen("tcp4", ":"+strconv.Itoa(s.config.HTTP.Port))
	if err != nil {
		return err
	}

	rootPath := server.NewGroupPath("")
	apiV1Path := rootPath.NewGroupPath("/api/v1")

	apiV1Path.UseBefore(s.beforeMiddleware)
	apiV1Path.UseAfter(s.afterMiddleware)

	apiV1Path.GET("/zeropark", s.ZeroParkPopunder)
	apiV1Path.GET("/zeropark/push", s.ZeroParkPush)

	apiV1Path.GET("/propellerads/push", s.PropellerAdsPush)
	apiV1Path.GET("/propellerads/custom", s.PropellerAdsCustom)

	apiV1Path.GET("/chevrolet", s.ChevroletPush)
	apiV1Path.POST("/chevrolet/impression", s.ChevroletImpression)

	apiV1Path.GET("/meetads", s.MeetAdsXML)
	apiV1Path.GET("/intango", s.IntangoXML)
	apiV1Path.GET("/evadav", s.Evadav)
	apiV1Path.GET("/datsun", s.Datsun)
	apiV1Path.GET("/volvo", s.Volvo)
	apiV1Path.GET("/mazda", s.Mazda)

	apiV1Path.POST("/openrtb", s.OpenRTB)
	apiV1Path.GET("/openrtb/burl", s.OpenRTBBiddingURL)
	apiV1Path.GET("/openrtb/nurl", s.OpenRTBNotificationURL)
	apiV1Path.GET("/openrtb/lurl", s.OpenRTBLossURL)

	apiV1Path.POST("/openrtb/native", s.OpenRTBNative)
	apiV1Path.POST("/openrtb/native/multibid", s.OpenRTBNativeMultiBid)

	rootPath.GET("/check", s.check)
	rootPath.NetHTTPPath(http.MethodGet, "/metrics", promhttp.Handler())

	s.logger.Info("Server running", zap.Int("port", s.config.HTTP.Port))
	select {
	case <-errch.Register(func() error { return server.ServeGracefully(ln) }):
		s.logger.Info("Shutdown server", zap.String("reason", "error"))
		return ln.Close()
	case <-ctx.Done():
		s.logger.Info("Shutdown server", zap.String("reason", "ctx.Done()"))
		return ln.Close()
	}
}

func (s *Server) price(args *fasthttp.Args) float64 {
	return args.GetUfloatOrZero("price")
}

func (s *Server) NewResponse() *HandlerResponse {
	return &HandlerResponse{
		StatusCode: http.StatusNoContent,
	}
}

func (s *Server) beforeMiddleware(ctx *atreugo.RequestCtx) error {
	skip := ctx.QueryArgs().GetBool("skip")

	if skip {
		ctx.SetStatusCode(http.StatusNoContent)
		return nil
	}

	return ctx.Next()
}

func (s *Server) afterMiddleware(ctx *atreugo.RequestCtx) error {
	delay := ctx.QueryArgs().GetUintOrZero("delay")

	time.Sleep(time.Duration(delay) * time.Millisecond)

	return ctx.Next()
}

func (s *Server) check(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("OK"))

	return nil
}

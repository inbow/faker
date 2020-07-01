package server

import (
	"context"
	"net/http"
	"strconv"

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

	server.GET("/api/v1/zeropark", s.ZeroPark)
	server.GET("/api/v1/zeropark/push", s.ZeroparkPush)

	server.GET("/api/v1/propellerads/push", s.PropellerAdsPush)
	server.GET("/api/v1/propellerads/custom", s.PropellerAdsCustom)

	server.GET("/api/v1/chevrolet", s.ChevroletPush)
	server.POST("/api/v1/chevrolet/impression", s.ChevroletImpression)

	server.GET("/api/v1/meetads", s.MeetAdsXML)
	server.GET("/api/v1/intango", s.IntangoXML)
	server.GET("/api/v1/evadav", s.Evadav)
	server.GET("/api/v1/datsun", s.Datsun)
	server.GET("/api/v1/volvo", s.Volvo)
	server.GET("/api/v1/mazda", s.Mazda)

	server.POST("/api/v1/openrtb", s.OpenRTB)
	server.POST("/api/v1/openrtb/native", s.OpenRTBNative)
	server.POST("/api/v1/openrtb/native/multibid", s.OpenRTBNativeMultiBid)
	server.GET("/api/v1/openrtb/burl", s.Burl)
	server.GET("/api/v1/openrtb/nurl", s.Nurl)
	server.GET("/api/v1/openrtb/lurl", s.Lurl)

	server.GET("/check", s.check)
	server.NetHTTPPath(http.MethodGet, "/metrics", promhttp.Handler())

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

func (s *Server) RequestValues(args *fasthttp.Args) (float64, int, bool) {
	price := args.GetUfloatOrZero("price")
	delay := args.GetUintOrZero("delay")
	skip := args.GetBool("skip")

	return price, delay, skip
}

func (s *Server) NewResponse() *HandlerResponse {
	return &HandlerResponse{
		StatusCode: http.StatusNoContent,
	}
}

func (s *Server) check(ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(http.StatusOK)
	ctx.SetBody([]byte("OK"))

	return nil
}

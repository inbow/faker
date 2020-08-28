package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/savsgio/atreugo/v11"
	"go.uber.org/zap"

	"github.com/oxyd-io/faker/internal/config"
	"github.com/oxyd-io/faker/pkg/env"
)

type (
	Server struct {
		ctx context.Context

		logger *zap.Logger
		config *config.AppConfig
	}
)

type UserValue string

const (
	URL   UserValue = "url"
	Price UserValue = "price"
)

func NewServer(
	ctx context.Context,
	logger *zap.Logger,
	config *config.AppConfig,
) *Server {
	return &Server{
		ctx: ctx,

		logger: logger,
		config: config,
	}
}

func (s *Server) Start() error {
	server := atreugo.New(atreugo.Config{
		Name: s.ctx.Value(env.Name).(string) + " server",
		Addr: s.config.HTTP.Host + ":" + strconv.Itoa(s.config.HTTP.Port),

		GracefulShutdown: true,
	})

	rootPath := server.NewGroupPath("")
	apiV1Path := rootPath.NewGroupPath("/api/v1")

	apiV1Path.UseBefore(s.beforeMiddleware)
	apiV1Path.UseAfter(s.afterMiddleware)

	apiV1Path.GET("/aesearch/push", s.AesearchPush)

	apiV1Path.GET("/audi/popunder", s.AudiPopunder)

	apiV1Path.GET("/buick/push", s.BuickPush)

	apiV1Path.GET("/chevrolet/push", s.ChevroletPush)
	apiV1Path.POST("/chevrolet/impression", s.ChevroletImpression)

	apiV1Path.GET("/datsun/native", s.DatsunNative)

	apiV1Path.GET("/evadav/popunder", s.EvadavPopunder)
	apiV1Path.GET("/evadav/push", s.EvadavPush)

	apiV1Path.GET("/ford/push", s.FordPush)

	apiV1Path.GET("/intango/push", s.Intango)

	apiV1Path.GET("/jaguar/push", s.JaguarPush)
	apiV1Path.GET("/jaguar/native", s.JaguarNative)

	apiV1Path.GET("/mack/push", s.MackPush)

	apiV1Path.GET("/mazda/popunder", s.MazdaPopunder)

	apiV1Path.GET("/meetads/popunder", s.MeetAds)

	apiV1Path.GET("/mgid/push", s.MgidPush)

	apiV1Path.POST("/openrtb/banner", s.OpenRTBBanner)
	apiV1Path.POST("/openrtb/popunder", s.OpenRTBPopunder)
	apiV1Path.POST("/openrtb/native", s.OpenRTBNative)
	apiV1Path.POST("/openrtb/native/multibid", s.OpenRTBNativeMultiBid)

	apiV1Path.GET("/openrtb/loss_url", s.OpenRTBLossURL)
	apiV1Path.GET("/openrtb/notice_url", s.OpenRTBNoticeURL)
	apiV1Path.GET("/openrtb/bidding_url", s.OpenRTBBiddingURL)

	apiV1Path.GET("/peugeot/push", s.PeugeotPush)
	apiV1Path.GET("/peugeot/popunder", s.PeugeotPopunder)

	apiV1Path.GET("/pontiac/popunder", s.PontiacPopunder)

	apiV1Path.GET("/propellerads/push", s.PropellerAdsPush)
	apiV1Path.GET("/propellerads/popunder", s.PropellerAdsPopunder)

	apiV1Path.GET("/volvo/push", s.Volvo)

	apiV1Path.GET("/zeropark/push", s.ZeroParkPush)
	apiV1Path.GET("/zeropark/popunder", s.ZeroParkPopunder)

	rootPath.GET("/check", s.check)
	rootPath.NetHTTPPath(http.MethodGet, "/metrics", promhttp.Handler())

	s.logger.Info(
		"Server is running",
		zap.String("host", s.config.HTTP.Host),
		zap.Int("port", s.config.HTTP.Port),
	)

	return server.ListenAndServe()
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) beforeMiddleware(ctx *atreugo.RequestCtx) error {
	skip := ctx.QueryArgs().GetBool("skip")

	if skip {
		ctx.SetStatusCode(http.StatusNoContent)
		return nil
	}

	price := ctx.QueryArgs().GetUfloatOrZero("price")
	ctx.SetUserValue(string(Price), price)

	url := string(ctx.QueryArgs().Peek("url"))
	ctx.SetUserValue(string(URL), url)

	return ctx.Next()
}

func (s *Server) afterMiddleware(ctx *atreugo.RequestCtx) error {
	delay := ctx.QueryArgs().GetUintOrZero("delay")

	time.Sleep(time.Duration(delay) * time.Millisecond)

	return ctx.Next()
}

func (s *Server) check(requestCtx *atreugo.RequestCtx) error {
	return requestCtx.TextResponse("OK", http.StatusOK)
}

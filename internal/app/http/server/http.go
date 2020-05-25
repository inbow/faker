package server

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	errch "github.com/proxeter/errors-channel"
	"go.uber.org/zap"

	"github.com/oxyd-io/faker/internal/app/config"
	"github.com/oxyd-io/faker/internal/app/generator"
)

type (
	Server struct {
		ctx context.Context

		logger *zap.Logger
		config *config.Config

		generator generator.IGenerator
	}

	panicHandler struct {
		logger *zap.Logger

		handler http.Handler
	}

	HandlerResponse struct {
		StatusCode int
		Body       []byte
	}
)

func (h panicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			body, _ := ioutil.ReadAll(r.Body)

			h.logger.Error(
				"recovered",
				zap.ByteString("body", body),
				zap.Any("error", err),
			)

			w.WriteHeader(http.StatusNoContent)
			_, _ = w.Write([]byte(""))

			return
		}
	}()

	body, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(body))

	h.logger.Info(
		"http request",
		zap.String("url", r.URL.String()),
		zap.Any("headers", r.Header),
		zap.ByteString("body", body),
	)

	h.handler.ServeHTTP(w, r)
}

func New(
	ctx context.Context,
	logger *zap.Logger,
	config *config.Config,
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
	server := http.Server{
		Addr:    ":" + strconv.Itoa(s.config.HTTP.Port),
		Handler: s.handler(),
	}

	s.logger.Info("Server running", zap.Int("port", s.config.HTTP.Port))
	select {
	case <-errch.Register(server.ListenAndServe):
		s.logger.Info("Shutdown server", zap.String("reason", "error"))
		return server.Shutdown(ctx)
	case <-ctx.Done():
		s.logger.Info("Shutdown server", zap.String("reason", "ctx.Done()"))
		return server.Shutdown(ctx)
	}
}

func (s *Server) ph(handler http.HandlerFunc) http.Handler {
	return &panicHandler{
		logger:  s.logger,
		handler: handler,
	}
}

func (s *Server) handler() *httprouter.Router {
	return func() *httprouter.Router {
		router := httprouter.New()

		router.Handler(http.MethodGet, "/propellerads/custom", s.ph(s.PropellerAdsCustom))
		router.Handler(http.MethodGet, "/propellerads/push", s.ph(s.PropellerAdsPush))

		router.Handler(http.MethodGet, "/zeropark", s.ph(s.ZeroPark))
		router.Handler(http.MethodGet, "/meetads", s.ph(s.MeetAdsXML))
		router.Handler(http.MethodGet, "/intango", s.ph(s.IntangoXML))
		router.Handler(http.MethodGet, "/evadav", s.ph(s.Evadav))
		router.Handler(http.MethodGet, "/datsun", s.ph(s.Datsun))
		router.Handler(http.MethodGet, "/volvo", s.ph(s.Volvo))
		router.Handler(http.MethodGet, "/mazda", s.ph(s.Mazda))

		router.Handler(http.MethodPost, "/openrtb", s.ph(s.OpenRTB))
		router.Handler(http.MethodPost, "/openrtb/vast", s.ph(s.OpenRTBVast))
		router.Handler(http.MethodPost, "/openrtb/native", s.ph(s.OpenRTBNative))
		router.Handler(http.MethodPost, "/openrtb/native/multibid", s.ph(s.OpenRTBNativeMultiBid))

		router.Handler(http.MethodGet, "/burl", s.ph(s.Burl))
		router.Handler(http.MethodGet, "/nurl", s.ph(s.Nurl))

		router.Handler(http.MethodGet, "/metrics", promhttp.Handler())
		router.HandlerFunc(http.MethodGet, "/check", s.check)

		return router
	}()
}

func (s *Server) RequestValues(query url.Values) (float64, int, bool) {
	var price float64
	if len(query.Get("price")) > 0 {
		price, _ = strconv.ParseFloat(query.Get("price"), 64)
	}

	delay := randomdata.Number(s.config.Bid.DelayMin, s.config.Bid.DelayMax)
	if len(query.Get("delay")) > 0 {
		delay, _ = strconv.Atoi(query.Get("delay"))
	}

	var skip bool
	if len(query.Get("skip")) > 0 {
		skip = true
	}

	return price, delay, skip
}

func (s *Server) NewResponse() *HandlerResponse {
	return &HandlerResponse{
		StatusCode: http.StatusNoContent,
	}
}

func (s *Server) check(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

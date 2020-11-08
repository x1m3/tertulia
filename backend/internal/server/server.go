package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const V1 = "/v1"

type Endpoint struct {
	Version string
	Method  string
	Path    string
	Handler HandlerFunc
}

type TokenAuthValidator func(ctx context.Context, jwtTokenStr string) error

type HTTPd struct {
	router *gin.Engine
	server *http.Server
}

func NewHTTPd(p int) *HTTPd {
	router := gin.New()

	router.Use(gin.Recovery())
	s := &HTTPd{
		router: router,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", p),
			Handler: router,
		},
	}
	return s
}

func (d *HTTPd) RegisterEndpoints(endpoints ...Endpoint) {
	for _, e := range endpoints {
		d.router.Handle(e.Method, e.Version+e.Path, d.encodeTo(e.Handler))
	}
}

func (d *HTTPd) CustomizeNotFoundError(handlers ...gin.HandlerFunc) {
	d.router.NoRoute(handlers...)
}

func (d *HTTPd) ListenAndServe(ctx context.Context) {
	if err := d.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithFields(logrus.Fields{"err": err}).Error("Error starting server")
	}
	logrus.Info(ctx, "Server down")
}

func (d *HTTPd) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.server.Handler.ServeHTTP(w, r)
}

func (d *HTTPd) Shutdown(ctx context.Context) {
	canCtx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFn()
	if err := d.server.Shutdown(canCtx); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("Error shooting down server")
	}
}

type HandlerFunc func(ctx *gin.Context) (resp interface{}, err error)
type encoder func(code int, obj interface{})

type errJSON struct {
	Error   string `form:"error" json:"error" xml:"error" yaml:"error"`
	Message string `form:"message" json:"message" xml:"message" yaml:"message"`
}

func (d *HTTPd) toFormat(ctx *gin.Context, fn HandlerFunc, encode encoder) {
	resp, err := fn(ctx)
	if err != nil {
		errMsg, HTTPCode := getHTTPCode(err)
		errFormatted := &errJSON{
			Error:   errMsg,
			Message: err.Error(),
		}
		encode(HTTPCode, errFormatted)
		return
	}
	encode(http.StatusOK, resp)
}

func (d *HTTPd) toJSON(ctx *gin.Context, fn HandlerFunc) {
	d.toFormat(ctx, fn, ctx.JSON)
}

func (d *HTTPd) toXML(ctx *gin.Context, fn HandlerFunc) {
	d.toFormat(ctx, fn, ctx.XML)
}

func (d *HTTPd) toYAML(ctx *gin.Context, fn HandlerFunc) {
	d.toFormat(ctx, fn, ctx.YAML)
}

func (d *HTTPd) encodeTo(fn HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contentType := ctx.GetHeader("Content-Type")
		if contentType == "" {
			ctx.Request.Header.Set("Content-Type", "application/json")
		}
		accept := ctx.GetHeader("Accept")
		switch accept {
		case gin.MIMEJSON:
			d.toJSON(ctx, fn)
		case gin.MIMEXML, gin.MIMEXML2:
			d.toXML(ctx, fn)
		case gin.MIMEYAML:
			d.toYAML(ctx, fn)
		default:
			d.toJSON(ctx, fn)
		}
	}
}

func getHTTPCode(err error) (string, int) {
	if err != nil {
		return "Internal_server_error", http.StatusInternalServerError
	}
	return "OK", http.StatusOK
}

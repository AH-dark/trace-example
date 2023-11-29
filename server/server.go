package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("server") // usually we use package name as the tracer name

func NewServer(ctx context.Context) (*server.Hertz, error) {
	ctx, span := tracer.Start(ctx, "server.NewServer")
	defer span.End()

	traceOption, cfg := hertztracing.NewServerTracer()
	svr := server.Default(
		traceOption,
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithTransport(netpoll.NewTransporter),
	)
	svr.Use(hertztracing.ServerMiddleware(cfg))

	return svr, nil
}

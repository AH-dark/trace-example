package server

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func HandlePing(ctx context.Context, c *app.RequestContext) {
	// In this step, a value called "span" is created, and it will be used to record the trace.
	// Also, it will fill the trace id into the context, so that the trace id can be passed to other spans.
	ctx, span := tracer.Start(ctx, "server.HandlePing")
	defer span.End()

	c.String(200, "pong")
}

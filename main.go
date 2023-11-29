package main

import (
	"context"
	"trace-example/common/observability"
	"trace-example/server"
)

var ctx = context.Background()

func main() {
	if err := observability.InitTracer(ctx); err != nil {
		panic(err)
	}

	svr, err := server.NewServer(ctx)
	if err != nil {
		panic(err)
	}

	svr.GET("/ping", server.HandlePing)

	svr.Spin()
}

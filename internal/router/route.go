package router

import (
	"context"

	"github.com/doffy007/coffee-shop/config"
	"github.com/labstack/echo/v4"
)

type router struct {
	ctx   context.Context
	conf  *config.Config
	route *echo.Echo
	// handler handler.Handler
}

// All implements Router.
func (r *router) All() {
	r.BaseRouter()
}

func Register(ctx context.Context, conf *config.Config, route *echo.Echo) Router {
	return &router{
		ctx:   ctx,
		conf:  conf,
		route: route,
		// handler: handler.NewHandler(ctx, conf),
	}
}

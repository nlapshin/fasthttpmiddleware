package fasthttpmiddelware

import "github.com/valyala/fasthttp"

type MiddlewareInstance func(reqHandler fasthttp.RequestHandler) fasthttp.RequestHandler

type Middleware struct {
	list []MiddlewareInstance
}

func New() *Middleware {
	var list []MiddlewareInstance

	return &Middleware{
		list,
	}
}

func (middleware *Middleware) Use(instance MiddlewareInstance) {
	middleware.list = append(middleware.list, instance)
}

func (middleware *Middleware) Handler(reqHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	for i := range middleware.list {
		reqHandler = middleware.list[i](reqHandler)
	}

	return reqHandler
}

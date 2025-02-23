package handler

import (
	"net/http"
)

type middleware func(http.Handler) http.Handler

type middlewareHandler struct {
	baseHandler http.Handler
	middlewares []middleware
}

func NewMiddlewareHandler(baseHandler http.Handler, middlewares ...middleware) *middlewareHandler {
	return &middlewareHandler{
		baseHandler: baseHandler,
		middlewares: middlewares,
	}
}

func (middlewareHandler *middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	baseHandler := middlewareHandler.baseHandler
	for i := 0; i < len(middlewareHandler.middlewares); i++ {
		baseHandler = middlewareHandler.middlewares[i](baseHandler)
	}
	baseHandler.ServeHTTP(w, r)
}

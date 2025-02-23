package hello

import (
	"net/http"

	handler "github.com/mainframematrix/todo/server/Handlers/Middleware"
	logging "github.com/mainframematrix/todo/server/Middleware/Logging"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func HelloMux() *http.ServeMux {
	helloMuxServer := http.NewServeMux()
	helloWorldHandler := http.HandlerFunc(helloWorld)
	middlewareHelloWorldHandler := handler.NewMiddlewareHandler(helloWorldHandler, logging.PathLogger)
	helloMuxServer.Handle("/", middlewareHelloWorldHandler)
	return helloMuxServer
}

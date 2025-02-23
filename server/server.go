package main

import (
	"log"
	"net/http"
	"os"

	logger "github.com/mainframematrix/go-log"
	"github.com/mainframematrix/todo/server/API/hello"
)

func main() {
	globalLogger, err := logger.CreateLogger(logger.INFO, "stdout", false)
	if err != nil {
		log.Fatalf("Error initializing logger.")
		os.Exit(1)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	helloMuxServer := hello.HelloMux()
	http.Handle("/hello", helloMuxServer)
	globalLogger.Log(logger.INFO, "Serving on port: %s", port)
	http.ListenAndServe(":"+port, nil)
}

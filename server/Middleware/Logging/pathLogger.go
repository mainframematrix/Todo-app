package logging

import (
	"log"
	"net/http"
	"os"

	logger "github.com/mainframematrix/go-log"
)

func PathLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggingPath := os.Getenv("LOGGING_PATH")
		if loggingPath == "" {
			loggingPath = "stdout"
		}
		l, err := logger.CreateLogger(logger.INFO, loggingPath, true)
		if err != nil {
			log.Fatalf("Error initializing logger: %v", err)
		}
		l.Log(logger.INFO, "%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

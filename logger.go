// Taken and adapted for standard http from https://github.com/gin-gonic/gin/blob/develop/logger.go
package logger

import (
	"log"
	"net/http"
	"os"
	"time"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	return w.ResponseWriter.Write(b)
}

// Logger returns a new logger to be wrapped around your main http.Handler
func Logger(next http.Handler) http.HandlerFunc {
	stdlogger := log.New(os.Stdout, "", 0)
	//errlogger := log.New(os.Stderr, "", 0)

	return func(w http.ResponseWriter, r *http.Request) {
		// Start timer
		start := time.Now()

		// Process request
		writer := statusWriter{w, 0}
		next.ServeHTTP(&writer, r)

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := r.RemoteAddr
		method := r.Method
		statusCode := writer.status
		statusColor := colorForStatus(statusCode)
		methodColor := colorForMethod(method)

		stdlogger.Printf("[HTTP] %v |%s %3d %s| %12v | %s |%s  %s %-7s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			statusColor, statusCode, reset,
			latency,
			clientIP,
			methodColor, reset, method,
			r.URL.Path,
		)
	}
}

func colorForStatus(code int) string {
	switch {
	case code >= 200 && code <= 299:
		return green
	case code >= 300 && code <= 399:
		return white
	case code >= 400 && code <= 499:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch {
	case method == "GET":
		return blue
	case method == "POST":
		return cyan
	case method == "PUT":
		return yellow
	case method == "DELETE":
		return red
	case method == "PATCH":
		return green
	case method == "HEAD":
		return magenta
	case method == "OPTIONS":
		return white
	default:
		return reset
	}
}

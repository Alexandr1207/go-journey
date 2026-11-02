package httpapi

import (
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

var idCounter atomic.Int64

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := idCounter.Add(1)
		idStr := strconv.FormatInt(id, 10)
		w.Header().Set("X-Request-ID", idStr)
		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &statusRecorder{ResponseWriter: w, status: 200}

		next.ServeHTTP(rec, r)

		id := w.Header().Get("X-Request-ID")

		log.Printf("[%s] %s %s -> %d (%v)", id, r.Method, r.URL.Path, rec.status, time.Since(start))
	})
}

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if v := recover(); v != nil {
				id := w.Header().Get("X-Request-ID")
				log.Printf("[%s] panic: %v", id, v)
				writeError(w, 500, "internal", "internal error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

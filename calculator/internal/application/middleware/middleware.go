package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func MethodRequest(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, `{"error":"Method is not allowed"}`, http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Panic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatalf("| Panic error: %s |", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("\n| Request |\n| Path: %s |\n| Method: %s |\n| Time: %d%s |", r.URL.Path, r.Method, time.Since(start).Milliseconds(), "ms")
		fmt.Println("")
	})
}

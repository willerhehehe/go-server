// Package middlewares: mux middlewares
package middlewares

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"time"
)

func RecoverWrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = fmt.Errorf(t)
				case error:
					err = t
				default:
					err = fmt.Errorf("Unknown error")
				}
				logrus.Errorf("stacktrace from panic: \n" + string(debug.Stack()))
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logrus.Info(fmt.Sprintf("%v %v\n", r.RequestURI, r.Method))
		// Call the next handlers, which can be another middleware in the chain, or the final handlers.
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		// Call the next handlers, which can be another middleware in the chain, or the final handlers.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func SlowRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()
		d := end.Sub(start)
		if d >= time.Second*1 {
			logrus.Warn(fmt.Sprintf("%v %v slow log duration: %v\n", r.RequestURI, r.Method, d))
		}
	})
}

package main

import (
	"net/http"
)

type Middleware func(*application, http.HandlerFunc) http.HandlerFunc

func httpReqLogger() Middleware {
	return func(app *application, next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			app.logger.Debug("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
			next(w, r)
		}
	}
}

// func loadAndSaveSession() Middleware {
// 	return func(app *application, next http.HandlerFunc) http.HandlerFunc {
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			app.sessionManager.LoadAndSave(next).ServeHTTP(w, r)
// 		}
// 	}
// }

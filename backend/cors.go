package main

import (
	"fmt"
	"net/http"
)

func CORSMiddleware(next http.Handler) http.Handler {
	allowedOrigins := map[string]bool{
		fmt.Sprintf("http://127.0.0.1:%v", frontendPort):           developmentMode, // allow this origin only in development mode
		fmt.Sprintf("https://127.0.0.1:%v", frontendPort):          developmentMode, // allow this origin only in development mode
		fmt.Sprintf("http://localhost:%v", frontendPort):           developmentMode, // allow this origin only in development mode
		fmt.Sprintf("https://localhost:%v", frontendPort):          developmentMode, // allow this origin only in development mode
		fmt.Sprintf("http://%s", frontendDomain):                   developmentMode, // allow this origin only in development mode
		fmt.Sprintf("https://%s", frontendDomain):                  true,
		fmt.Sprintf("http://%s:%v", frontendDomain, frontendPort):  developmentMode, // allow this origin only in development mode
		fmt.Sprintf("https://%s:%v", frontendDomain, frontendPort): true,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowed, found := allowedOrigins[origin]; !found || !allowed {
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next.ServeHTTP(w, r)
	})
}

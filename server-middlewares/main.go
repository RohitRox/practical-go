package main

import (
	"fmt"
	"net/http"
)

func MiddlewareAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Visiting Auth Middleware")
		next(w, r)
	}
}

func MiddlewareGlobal(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Visiting Global Middleware")
		next.ServeHTTP(w, r)
	})
}

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Visiting Index Handler")
	fmt.Fprintf(w, "Hello!")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/index", MiddlewareAuth(indexHandler))

	f := &Logger{router}

	finalRouter := MiddlewareGlobal(f)

	fmt.Println("Starting server at 0.0.0.0:8090")

	http.ListenAndServe(":8090", finalRouter)

}

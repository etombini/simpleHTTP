package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type server struct {
	router  *http.ServeMux
	http    *http.Server
	address string
	static  string
	listDir bool
}

func newServer(listen string, static string, listDir bool) *server {
	if s, err := os.Stat(static); err != nil {
		fmt.Fprintf(os.Stderr, "error with directory for static content: %s\n", err)
		return nil
	} else if !s.IsDir() {
		fmt.Fprintf(os.Stderr, "%s is not a directory\n", static)
		return nil
	}

	s := &server{
		router:  http.NewServeMux(),
		address: listen,
		static:  static,
		listDir: listDir,
	}
	s.setRoutes()

	s.http = &http.Server{
		Addr:    listen,
		Handler: s.router,
	}
	return s

}

func (s *server) setRoutes() {
	s.router.Handle("/", isListDir(s.listDir, http.FileServer(http.Dir(s.static))))
}

func isListDir(listDir bool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !listDir && strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

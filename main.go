package main

import (
	"fmt"
	"os"
)

func main() {
	cfg := getConfig()
	s := newServer(cfg.Listen, cfg.Path, cfg.ListDir)
	if s == nil {
		fmt.Fprintf(os.Stderr, "can not create server\n")
		os.Exit(1)
	}

	if err := s.http.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/etombini/flag"
)

type config struct {
	Path    string `names:"-p,--path"`
	Listen  string `names:"-l,--listen"`
	ListDir bool   `names:"-ld,--list-dir"`
}

func getConfig() *config {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	c := &config{
		Path:   wd,
		Listen: "127.0.0.1:8080",
	}

	fs := flag.NewFlagSet(c)
	fs.Parse()
	return c
}

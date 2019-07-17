package main

import (
	"fmt"
	"os"

	"github.com/izumin5210/clig/pkg/clib"

	"github.com/rerost/dagg/pkg/dagg/cmd"
)

const (
	appName = "dagg"
	version = "v0.0.1"
)

var (
	revision, buildDate string
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	cmd := cmd.NewDefaultDaggCommand(clib.Path(wd), clib.Build{
		AppName:   appName,
		Version:   version,
		Revision:  revision,
		BuildDate: buildDate,
	})

	return cmd.Execute()
}

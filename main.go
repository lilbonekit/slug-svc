package main

import (
	"os"

	"github.com/lilbonekit/slug-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}

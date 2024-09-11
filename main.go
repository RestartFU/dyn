package main

import (
	"github.com/restartfu/dyn/cmd/cli"
	"github.com/restartfu/dyn/internal/logger"
)

var version = logger.Color("<aqua>v0.1.1</aqua>")

func main() {
	cli.Execute(version)
}

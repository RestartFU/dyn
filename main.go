package main

import (
	"flag"

	"github.com/restartfu/dyn/cmd/cli"
	"github.com/restartfu/dyn/internal/logger"
)

var version = logger.Color("<aqua>v0.1.4</aqua>")

func main() {
	nosudo := flag.Bool("nosudo", false, "")
	pkgDir := flag.String("pkgdir", "/usr/local/dyn-pkg", "")
	flag.Parse()

	c := cli.CLI{
		Version:   version,
		ForceSudo: !*nosudo,
		PkgDir:    *pkgDir,
	}
	c.Execute()
}

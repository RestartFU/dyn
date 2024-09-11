package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/restartfu/dyn/internal/logger"
	"github.com/savioxavier/termlink"
)

type CLI struct {
	Version   string
	ForceSudo bool
	PkgDir    string
}

func (c CLI) Execute() {
	if arg, ok := arg(1); ok && arg == "version" {
		fmt.Println(c.Version)
		return
	}

	_, su := os.LookupEnv("SUDO_COMMAND")
	if !su && c.ForceSudo {
		logger.Fatalf("dyn must be run as a super user.\n")
	}

	act, ok := arg(1)
	if !ok || !isAnyString(act, "update", "install", "remove", "fetch") {
		logger.Fatalf("valid actions: update|install|remove|fetch.\n")
	}

	pkgArgN := 2
	if act == "fetch" {
		c.fetch()
		act, ok = arg(pkgArgN)
		if !ok || !isAnyString(act, "update", "install", "remove") {
			return
		}
		pkgArgN = 3
	}

	pkg, ok := arg(pkgArgN)
	if !ok {
		if act != "update" {
			logger.Fatalf("please specify the package you wish to %s.\n", act)
		}
		act = "update"
		pkg = "dyn"
	}

	c.executePackage(pkg, act)
	logger.Dynf("done %s package %s.\n", verb(act), pkg)
}

func (c CLI) executePackage(pkg string, act string) {
	targetPkgPath := filepath.Join(c.PkgDir, pkg)
	if _, err := os.Stat(targetPkgPath); os.IsNotExist(err) {
		logger.Fatalf("no package found with the name %s, maybe run 'dyn fetch', and try again?\n", pkg)
	}

	scriptPath := filepath.Join(targetPkgPath, "DYNPKG")
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		logger.Fatalf("no DYNPKG file found for package %s, maybe run 'dyn fetch', and try again?\n", pkg)
	}
	scriptBuf, err := os.ReadFile(scriptPath)
	if err != nil {
		logger.Fatalf("could not read DYNPKG file for package %s", pkg)
	}

	script := []string{
		string(scriptBuf),
		act,
		`if [ -n "$maintainers" ]; then
			credits=$(echo "special thanks to ( $maintainers ) for maintaining this package")
			echo $credits;
		 fi`,

		"echo \"if you wish to contribute, make sure to check out " + termlink.ColorLink("our github page",
			"https://github.com/restartfu/dyn", "yellow") + "\" >&1",
	}
	tmpScriptPath := filepath.Join(os.TempDir(), "dyn-pkg", pkg, "script.sh")
	tmpDir := filepath.Dir(tmpScriptPath)
	os.RemoveAll(tmpDir)
	os.MkdirAll(tmpDir, os.ModePerm)
	os.WriteFile(tmpScriptPath, []byte(strings.Join(script, "\n")), os.ModePerm)

	logger.Dynf("%s package %s.\n", verb(act), pkg)
	cmd := exec.Command("sh", "-c", tmpScriptPath)
	cmd.Stdout = logger.InfoOut
	cmd.Stderr = os.Stderr
	cmd.Run()

	os.RemoveAll(tmpDir)
}

func (c CLI) fetch() {
	logger.Dynf("fetching dyn-pkg git repository.\n")

	os.RemoveAll(c.PkgDir)
	_, err := git.PlainClone(c.PkgDir, false, &git.CloneOptions{
		Depth:    1,
		URL:      "https://github.com/RestartFU/dyn-pkg",
		Progress: logger.InfoOut,
	})

	if err != nil {
		logger.Fatalf("error fetching dyn package repository: %s.\n", err)
	}
}

package cli

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/restartfu/dyn/internal/logger"
)

var (
	pkgPath = "/usr/local/dyn-pkg"
)

func Execute() {
	exe, err := os.Executable()
	if err != nil {
		logger.Fatalf("somehow we cannot know where the executable is being run from: %s", err)
	}

	executable := filepath.Base(exe)
	_, su := os.LookupEnv("SUDO_COMMAND")
	if !su {
		logger.Fatalf("%s must be run as a super user.", executable)
	}

	act, ok := flag(1)
	if !ok || !isAnyString(act, "update", "install", "remove", "fetch") {
		logger.Fatalf("valid actions: update|install|remove|fetch")
	}

	if act == "fetch" {
		fetch()
		return
	}

	pkg, ok := flag(2)
	if !ok {
		logger.Fatalf("please specify the package you wish to %s", act)
	}

	targetPkgPath := filepath.Join(pkgPath, pkg)
	if _, err := os.Stat(targetPkgPath); os.IsNotExist(err) {
		logger.Fatalf("no package found with the name %s, maybe run '%s fetch', and try again?", pkg, executable)
	}

	scriptPath := filepath.Join(targetPkgPath, act+".sh")
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		logger.Fatalf("no %s action found for the package %s, maybe run '%s fetch', and try again?", act, pkg, executable)
	}

	cmd := exec.Command("sh", "-c", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func flag(n int) (string, bool) {
	args := os.Args
	if len(args) <= n {
		return "", false
	}
	return args[n], true
}

func fetch() {
	os.RemoveAll(pkgPath)
	_, err := git.PlainClone(pkgPath, false, &git.CloneOptions{
		Depth:    1,
		URL:      "https://github.com/RestartFU/dyn-pkg",
		Progress: os.Stdout,
	})

	if err != nil {
		logger.Fatalf("error fetching dyn package repository: %s", err)
	}
}

func isAnyString(s string, a ...string) bool {
	for _, str := range a {
		if str == s {
			return true
		}
	}
	return false
}

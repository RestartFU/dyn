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
		logger.Fatalf("somehow we cannot know where the executable is being run from: %s.\n", err)
	}

	executable := filepath.Base(exe)
	_, su := os.LookupEnv("SUDO_COMMAND")
	if !su {
		logger.Fatalf("%s must be run as a super user.\n", executable)
	}

	act, ok := arg(1)
	if !ok || !isAnyString(act, "update", "install", "remove", "fetch") {
		logger.Fatalf("valid actions: update|install|remove|fetch.\n")
	}

	pkgArgN := 2
	if act == "fetch" {
		fetch()
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

	targetPkgPath := filepath.Join(pkgPath, pkg)
	if _, err := os.Stat(targetPkgPath); os.IsNotExist(err) {
		logger.Fatalf("no package found with the name %s, maybe run '%s fetch', and try again?\n", pkg, executable)
	}

	scriptPath := filepath.Join(targetPkgPath, "DYNPKG")
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		logger.Fatalf("no DYNPKG file found for package %s, maybe run '%s fetch', and try again?\n", pkg, executable)
	}
	scriptBuf, err := os.ReadFile(scriptPath)
	if err != nil {
		logger.Fatalf("could not read DYNPKG file for package %s", pkg)
	}
	script := string(scriptBuf) + "\n" + act
	tmpScriptPath := filepath.Join(os.TempDir(), "dyn-pkg", pkg, "script.sh")
	tmpDir := filepath.Dir(tmpScriptPath)
	os.RemoveAll(tmpDir)
	os.MkdirAll(tmpDir, os.ModePerm)
	os.WriteFile(tmpScriptPath, []byte(script), os.ModePerm)

	logger.Dynf("%s package %s.\n", verb(act), pkg)
	cmd := exec.Command("sh", "-c", tmpScriptPath)
	cmd.Stdout = logger.InfoOut
	cmd.Stderr = os.Stderr
	cmd.Run()

	logger.Dynf("done %s package %s.\n", verb(act), pkg)
}

func verb(s string) string {
	switch s {
	case "install":
		return "installing"
	case "update":
		return "updating"
	case "remove":
		return "removing"
	case "fetch":
		return "fetching"
	}
	panic("should never happend")
}

func arg(n int) (string, bool) {
	args := os.Args
	if len(args) <= n {
		return "", false
	}
	return args[n], true
}

func fetch() {
	logger.Dynf("fetching dyn-pkg git repository.\n")

	os.RemoveAll(pkgPath)
	_, err := git.PlainClone(pkgPath, false, &git.CloneOptions{
		Depth:    1,
		URL:      "https://github.com/RestartFU/dyn-pkg",
		Progress: logger.InfoOut,
	})

	if err != nil {
		logger.Fatalf("error fetching dyn package repository: %s.\n", err)
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

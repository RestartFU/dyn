package cli

import (
	"os"
)

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

func isAnyString(s string, a ...string) bool {
	for _, str := range a {
		if str == s {
			return true
		}
	}
	return false
}

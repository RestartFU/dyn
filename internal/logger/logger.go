package logger

import (
	"fmt"
	"os"

	"github.com/sandertv/gophertunnel/minecraft/text"
)

func Debugf(str string, args ...any) {
	print("<blue>DEBU</blue>", str, args...)
}

func Fatalf(str string, args ...any) {
	print("<red>FATA</red>", str, args...)
	os.Exit(0)
}

func Errorf(str string, args ...any) {
	print("<redstone>ERRO</redstone>", str, args...)
}

func Infof(str string, args ...any) {
	print("<yellow>INFO</yellow>", str, args...)
}

func Dynf(str string, args ...any) {
	print("<aqua>DYN </aqua>", str, args...)
}

func print(prefix string, str string, args ...any) {
	fmt.Print(text.ANSI(text.Colourf(prefix+"<grey>|</grey> ")) + fmt.Sprintf(str, args...))
}

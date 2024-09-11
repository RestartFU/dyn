package logger

import (
	"fmt"
	"os"

	"github.com/sandertv/gophertunnel/minecraft/text"
)

func Debugf(str string, args ...any) {
	printf("<blue>DEBU</blue>", str, args...)
}

func Fatalf(str string, args ...any) {
	printf("<red>FATA</red>", str, args...)
	os.Exit(0)
}

func Fatal(str string) {
	print("<red>FATA</red>", str)
	os.Exit(0)
}

func Errorf(str string, args ...any) {
	printf("<redstone>ERRO</redstone>", str, args...)
}

func Error(str string) {
	print("<redstone>ERRO</redstone>", str)
}

func Infof(str string, args ...any) {
	printf("<yellow>INFO</yellow>", str, args...)
}

func Info(str string) {
	print("<yellow>INFO</yellow>", str)
}

func Dynf(str string, args ...any) {
	printf("<aqua>DYN </aqua>", str, args...)
}

func printf(prefix string, str string, args ...any) {
	fmt.Print(text.ANSI(text.Colourf(prefix+"<grey>|</grey> ")) + fmt.Sprintf(str, args...))
}

func print(prefix string, str string) {
	fmt.Print(text.ANSI(text.Colourf(prefix+"<grey>|</grey> ")) + str)
}

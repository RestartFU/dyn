package logger

import (
	"fmt"
	"os"

	"github.com/sandertv/gophertunnel/minecraft/text"
)

func Debugf(str string, args ...any) {
	print("<blue>DEBUG</blue>", str, args...)
}

func Fatalf(str string, args ...any) {
	print("<red>FATAL</red>", str, args...)
	os.Exit(0)
}

func Infof(str string, args ...any) {
	print("<yellow>INFO</yellow>", str, args...)
}

func print(prefix string, str string, args ...any) {
	fmt.Println(text.ANSI(text.Colourf(prefix+"<grey>|</grey> ")) + fmt.Sprintf(str, args...))
}

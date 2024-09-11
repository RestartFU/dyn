package logger

var (
	FatalOut = Writer{log: Fatalf}
	ErrorOut = Writer{log: Errorf}
	InfoOut  = Writer{log: Infof}
)

type Writer struct {
	log func(str string, args ...any)
}

func (w Writer) Write(p []byte) (n int, err error) {
	w.log(string(p))
	return len(p), err
}

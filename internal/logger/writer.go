package logger

var (
	FatalOut = Writer{log: Fatal}
	ErrorOut = Writer{log: Error}
	InfoOut  = Writer{log: Info}
)

type Writer struct {
	log func(str string)
}

func (w Writer) Write(p []byte) (n int, err error) {
	w.log(string(p))
	return len(p), err
}

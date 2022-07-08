package utils

type Logger struct {
	log []string
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.log = append(l.log, string(p))
	return len(p), nil
}

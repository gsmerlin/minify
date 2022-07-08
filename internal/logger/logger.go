package logger

import (
	"log"
	"time"
)

type internalLogger struct {
	log []string
}

func (l *internalLogger) Write(p []byte) (n int, err error) {
	l.log = append(l.log, string(p))
	return len(p), nil
}

func Start(isInternal bool) {
	// We're removing flags because we're adding timestamps manually so we can use the colors we want
	log.SetFlags(0)
	if isInternal {
		log.SetOutput(&internalLogger{})
	}
}

func Info(t string) {
	var message Text
	message.Blue().Bold().Text(time.Now().Format(time.UnixDate))
	message.Yellow().Bold().Text("Info: ")
	message.Text(t)
	log.Println(string(message))
}

func Success(t string) {
	var message Text
	message.Blue().Bold().Text(time.Now().Format(time.UnixDate))
	message.Green().Bold().Text("Success: ")
	message.Text(t)
	log.Println(string(message))
}

func Error(t string) {
	var message Text
	message.Blue().Bold().Text(time.Now().Format(time.UnixDate))
	message.Red().Bold().Text("Error: ")
	message.Text(t)
	log.Println(string(message))
}

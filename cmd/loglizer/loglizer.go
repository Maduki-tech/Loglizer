package loglizer

import (
	"log"
	"os"
)

type Loglizer struct {
	file *os.File
}

func NewLoglizer(file *os.File) *Loglizer {

	return &Loglizer{
		file: file,
	}
}

func (l *Loglizer) Info(message string) {
	infoLogger := log.New(l.file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger.Println(message)
}

func (l *Loglizer) Warn(message string) {
	warnLogger := log.New(l.file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger.Println(message)
}

func (l *Loglizer) Error(message string) {
	errorLogger := log.New(l.file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger.Println(message)
}

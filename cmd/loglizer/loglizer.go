package loglizer

import (
	"io"
	"log"
)

type Loglizer struct {
	logger *log.Logger
}

func NewLoglizer(writer io.Writer) *Loglizer {

	logger := log.New(writer, "", log.LstdFlags)
	return &Loglizer{
		logger: logger,
	}
}

func (l *Loglizer) Info(message string) {
	l.logger.Println(message)
}

func (l *Loglizer) Warn(message string) {
	l.logger.Println(message)
}

func (l *Loglizer) Error(message string) {
	l.logger.Println(message)
}

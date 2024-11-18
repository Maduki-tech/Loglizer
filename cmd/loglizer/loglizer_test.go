package loglizer

import (
	"os"
	"testing"
)

func TestLoglizer(t *testing.T) {
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	log := NewLoglizer(file)

	log.Info("info")
	log.Warn("warn")
	log.Error("error")
}

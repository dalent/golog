package golog

import (
	"os"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	writer := NewFileWriter()
	log := New(writer)
	log.Debug("%s", "me")

	fi, _ := os.Open(writer.writer.fd.Name())
	defer fi.Close()
	var bytes [50]byte
	fi.Read(bytes[:])
	if strings.HasSuffix(string(bytes[:]), "me") {
		t.Fatal("write byte not equal")
	}

	os.RemoveAll(writer.dir)
}
func TestCallDepth(t *testing.T) {
	writer := NewFileWriter()
	log := New(writer)
	log.SetCallDepth(2)
	log.Debug("%s", "me")

	fi, _ := os.Open(writer.writer.fd.Name())
	defer fi.Close()
	var bytes [50]byte
	fi.Read(bytes[:])
	if !strings.Contains(string(bytes[:]), "log_test.go:28") {
		t.Fatal("call depth not equal")
	}

	os.RemoveAll(writer.dir)
}
func TestConsole(t *testing.T) {
	writer := NewConsoleWriter()
	log := New(writer)
	log.Debug("%s", "me")
	log.Error("%s", "me")
	log.Info("%s", "me")
	log.Warn("%s", "me")
}

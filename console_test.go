package golog

import (
	"testing"
)

func TestCon(t *testing.T) {
	console := NewConsoleWriter()
	console.StartLogger()
	console.WriteMsg("test")
}

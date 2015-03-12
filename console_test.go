package golog

import (
	"testing"
)

func TestCon(t *testing.T) {
	console := NewConsoleWriter()
	console.startLogger()
	console.writeMsg("test")
}

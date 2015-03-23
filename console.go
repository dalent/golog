package golog

import (
	"github.com/mattn/go-colorable"
	"log"
)

var (
	color = []string{
		string([]byte{27, 91, 48, 109}),                 //reset
		string([]byte{27, 91, 57, 55, 59, 52, 50, 109}), //green
		string([]byte{27, 91, 57, 55, 59, 52, 51, 109}), //yellow
		string([]byte{27, 91, 57, 55, 59, 52, 52, 109}), //blue
		string([]byte{27, 91, 57, 55, 59, 52, 49, 109}), //red
		string([]byte{27, 91, 57, 55, 59, 52, 49, 109}), //red
	}
)

type ConsoleWriter struct {
	lg *log.Logger
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{log.New(colorable.NewColorableStdout(), "", 0)}
}

func (p *ConsoleWriter) startLogger() {
}

func (p *ConsoleWriter) writeMsg(msg string) {
	p.lg.Printf(msg)
}

func (p *ConsoleWriter) color(level int) string {
	return color[level]
}
func (p *ConsoleWriter) flush() {
}

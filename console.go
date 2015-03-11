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
	}
)

type ConsoleWriter struct {
	lg    *log.Logger
	level int
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{log.New(colorable.NewColorableStdout(), "", 0), LDEBUG}
}

func (p *ConsoleWriter) SetLevel(level int) {
	p.level = level
}
func (p *ConsoleWriter) StartLogger() {
}

func (p *ConsoleWriter) WriteMsg(msg string) {
	p.lg.Printf(msg)
}

func (p *ConsoleWriter) Color(level int) string {
	return color[level]
}
func (p *ConsoleWriter) Flush() {
	p.Flush()
}

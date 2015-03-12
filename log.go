package golog

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

const (
	RESET = iota // color
	LDEBUG
	LWARN
	LINFO
	LERROR
	//LFATAL   no fatal
	FROM_TIME = "2006-01-02 15:04:05"
)

var gLevelName = []string{"RESET", "DEBUG", "WARN", "INFO", "ERROR"}

type GoLog struct {
	level     int
	callDepth int
	writer    Writer
}

type Writer interface {
	color(int) string
	startLogger()
	writeMsg(msg string)
	flush()
}

func New(logger Writer) *GoLog {
	logger.startLogger()
	goLog := &GoLog{LDEBUG, 0, logger}
	return goLog
}

func (p *GoLog) SetLevel(level int) {
	if level > LERROR || level < LDEBUG {
		panic("level error")
	}
	p.level = level
}

func (p *GoLog) SetCallDepth(depth int) {
	p.callDepth = depth
}

func (p *GoLog) writeString(level int, msg string) {
	if p.level > level {
		return
	}

	var (
		file    string
		line    int
		strCall string
	)
	if p.callDepth != 0 {
		var ok bool
		_, file, line, ok = runtime.Caller(p.callDepth)
		if !ok {
			file = "???"
			line = 0
		} else {
			_, file = path.Split(file)
		}

		strCall = fmt.Sprintf("[%s:%d]", file, line)
	}

	p.writer.writeMsg(
		fmt.Sprintf("[%s]%s[%s%s%s] %s\n",
			time.Now().Format(FROM_TIME),
			strCall,
			p.writer.color(level),
			gLevelName[level],
			p.writer.color(RESET),
			msg,
		))
}

func (p *GoLog) Error(format string, msg ...interface{}) {
	p.writeString(LERROR, fmt.Sprintf(format, msg...))
}
func (p *GoLog) Info(format string, msg ...interface{}) {
	p.writeString(LINFO, fmt.Sprintf(format, msg...))
}
func (p *GoLog) Warn(format string, msg ...interface{}) {
	p.writeString(LWARN, fmt.Sprintf(format, msg...))
}
func (p *GoLog) Debug(format string, msg ...interface{}) {
	p.writeString(LDEBUG, fmt.Sprintf(format, msg...))
}

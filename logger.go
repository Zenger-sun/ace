package ace

import (
	"fmt"
	"os"

	"ace/utils"
)

type LogLevel string

const (
	LOG_FILE = "ace.log"

	DEBUG_LEVEL   LogLevel = "Debug"
	INFO_LEVEL    LogLevel = "Info"
	WARNING_LEVEL LogLevel = "Warning"
	ERR_LEVEL     LogLevel = "Err"
	CRIT_LEVEL    LogLevel = "Crit"
)

type Logger interface {
	Debug(m string, a ...interface{})
	Info(m string, a ...interface{})
	Warning(m string, a ...interface{})
	Err(m string, a ...interface{})
	Crit(m string, a ...interface{})
	Close()
}

type tagLog struct {
	tag     string
	path    string
	logFile *os.File
}

func GetLogItf(tag, path string) Logger {
	logFile, _ := os.OpenFile(utils.UnionStr(path, LOG_FILE), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	return &tagLog{
		tag:     tag,
		path:    path,
		logFile: logFile,
	}
}

func (l *tagLog) Close() {
	l.logFile.Close()
}

func (l *tagLog) print(level LogLevel, m string, a []interface{}) {
	t := utils.TimeNowStr()
	content := ""

	if len(a) == 0 {
		content = fmt.Sprintf("[%s][%s] %s %s\n", l.tag, level, t, m)
	} else {
		content = fmt.Sprintf("[%s][%s] %s %s\n", l.tag, level, t, fmt.Sprintf(m, a...))
	}

	fmt.Print(content)
	l.logFile.WriteString(content)
}

func (l *tagLog) Log(level LogLevel, m string, a ...interface{}) {
	l.print(level, m, a)
}

func (l *tagLog) Debug(m string, a ...interface{}) {
	l.print(DEBUG_LEVEL, m, a)
}

func (l *tagLog) Info(m string, a ...interface{}) {
	l.print(INFO_LEVEL, m, a)
}

func (l *tagLog) Warning(m string, a ...interface{}) {
	l.print(WARNING_LEVEL, m, a)
}

func (l *tagLog) Err(m string, a ...interface{}) {
	l.print(ERR_LEVEL, m, a)
}

func (l *tagLog) Crit(m string, a ...interface{}) {
	l.print(CRIT_LEVEL, m, a)
}

package olog

import (
	"fmt"
	debug2 "runtime/debug"
	"time"
)

type Olog int

const (
	ALL   Olog = iota //打印所有日志
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF    //关闭日志

	debug = "[ DEBUG ]"
	info  = "[ INFO  ]"
	warn  = "[ WARN  ]"
	err   = "[ ERROR ]"
	fatal = "[ FATAL ]"
)

var (
	logLevel   = ALL
	timeFormat = "| 2006-01-02 15:04:05 |"
	logOps     LogOptions
)

type LogOptions struct {
	LogLevel Olog
	IsPrint  bool
	FilePath string
}

func SetOlogOptions(ops LogOptions) {
	logOps = ops
}

func SetLogLevel(olog Olog) {
	logLevel = olog
}

func getTimeNow() (timeStr string) {
	return time.Now().Format(timeFormat)
}

func Simple() {
	Debug("hello debug")
	Info("hello information")
	Warn("hello warning")
	Error("hello error")
	Fatal("hello fatal")
}

func Debug(v ...interface{}) {
	if logLevel > DEBUG {
		return
	}
	v = append([]interface{}{getTimeNow(), debug}, v...)
	fmt.Println(v...)
}

func Info(v ...interface{}) {
	if logLevel > INFO {
		return
	}
	v = append([]interface{}{getTimeNow(), info}, v...)
	fmt.Println(v...)
}

func Warn(v ...interface{}) {
	if logLevel > WARN {
		return
	}
	v = append([]interface{}{getTimeNow(), warn}, v...)
	fmt.Println(v...)
}

func Error(v ...interface{}) {
	if logLevel > ERROR {
		return
	}
	stack := debug2.Stack()
	v = append([]interface{}{getTimeNow(), err}, v...)
	v = append(v, string(stack))
	fmt.Println(v...)
}

func Fatal(v ...interface{}) {
	if logLevel > FATAL {
		return
	}
	stack := debug2.Stack()
	v = append([]interface{}{getTimeNow(), fatal}, v...)
	v = append(v, string(stack))
	fmt.Println(v...)
}

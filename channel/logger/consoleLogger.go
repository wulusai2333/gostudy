package logger

import (
	"fmt"
	"time"
)

/*
	打印日志日志的几种形态
	Debug Trace Info Warning Error Fatal
	需求:不同位置输出日志 日志分级别 日志支持开关 日志记录包含时间行号等信息 日志文件要切割
*/
/*
	下面开始写日志类
*/

//日志类
type Logger struct {
	Level Loglevel
}

//日志类构造函数
func NewLogger(loglevel Loglevel) (l *Logger) {
	if loglevel > FATAL {
		panic("unknow level!")
	}
	return &Logger{loglevel}
}

/*
	日志打印格式
*/
func (l *Logger) log(LogLevelName string, format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	file, funcName, line := getInfo(3)
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s [\"%s\" %s:%d] %s\n", now, LogLevelName, file, funcName, line, str)
}

func (l *Logger) Debug(format string, a ...interface{}) {
	if l.Level == DEBUG {
		l.log("Debug", format, a...)
	}
}

func (l *Logger) Trace(format string, a ...interface{}) {
	if l.Level <= TRACE {
		l.log("Trace", format, a...)
	}

}
func (l *Logger) Info(format string, a ...interface{}) {
	if l.Level <= INFO {
		l.log("Info", format, a...)
	}

}
func (l *Logger) Warning(format string, a ...interface{}) {
	if l.Level <= WARNING {
		l.log("Warning", format, a...)
	}

}
func (l *Logger) Error(format string, a ...interface{}) {
	if l.Level <= ERROR {
		l.log("Error", format, a...)
	}

}
func (l *Logger) Fatal(format string, a ...interface{}) {
	if l.Level <= FATAL {
		l.log("Fatal", format, a...)
	}

}

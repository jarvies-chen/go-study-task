package logger

import (
	"log"
	"os"
)

// Logger 结构体
type Logger struct {
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	warningLogger *log.Logger
}

// 全局 Logger 实例
var defaultLogger *Logger

// 初始化默认 Logger
func init() {
	defaultLogger = NewLogger()
}

// NewLogger 创建新的 Logger 实例
func NewLogger() *Logger {
	return &Logger{
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info 记录信息日志
func (l *Logger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Infof 格式化记录信息日志
func (l *Logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

// Error 记录错误日志
func (l *Logger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

// Errorf 格式化记录错误日志
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

// Warning 记录警告日志
func (l *Logger) Warning(v ...interface{}) {
	l.warningLogger.Println(v...)
}

// Warningf 格式化记录警告日志
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warningLogger.Printf(format, v...)
}

// 全局便捷函数
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

func Warning(v ...interface{}) {
	defaultLogger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	defaultLogger.Warningf(format, v...)
}

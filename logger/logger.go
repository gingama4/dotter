package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Log Color
const (
	_log_c_debug int = 32
	_log_c_info  int = 34
	_log_c_warn  int = 33
	_log_c_error int = 31
	_log_c_field int = 35
)

type logger struct {
	Logging      bool
	Step         int
	StepSeparate string
	Level        LogLevel
	field        Fields
}

type logInfo struct {
	message    string
	levelColor int
	level      string
	step       int
	separate   string
	field      Fields
	output     io.Writer
}

var log *logger

func InitLog(isLogging bool) {
	l := logger{
		Logging:      isLogging,
		Step:         0,
		StepSeparate: "   ",
		Level:        INFO,
		field:        Fields{},
	}

	log = &l
}

func Log() *logger {
	return log
}

func SetLevel(l LogLevel) {
	log.Level = l
}

func (l *logger) SetStep(s int) *logger {
	l.Step = s
	return l
}

func (l *logger) StepUp() *logger {
	l.Step++
	return l
}

func (l *logger) StepDown() *logger {
	l.Step--
	return l
}

func (l *logger) StepReset() *logger {
	l.Step = 0
	return l
}

func (l *logger) AddField(k string, v interface{}) *logger {
	l.field[k] = v

	return l
}

func (l *logger) log(level LogLevel, ls string, lc int, w io.Writer, m string, args ...interface{}) {
	if !l.isLogging(level) {
		return
	}

	var mes string
	if len(args) > 0 {
		mes = fmt.Sprintf(m, args...)
	} else {
		mes = m
	}

	li := logInfo{
		message:    mes,
		levelColor: lc,
		level:      ls,
		step:       l.Step,
		separate:   l.StepSeparate,
		field:      l.field,
		output:     w,
	}
	li.log()

	l.field = Fields{}
}

func (l *logger) Debug(m string, args ...interface{}) {
	l.log(DEBUG, "DEBG", _log_c_debug, os.Stdout, m, args...)
}

func (l *logger) Info(m string, args ...interface{}) {
	l.log(INFO, "INFO", _log_c_info, os.Stdout, m, args...)
}

func (l *logger) Warn(m string, args ...interface{}) {
	l.log(WARN, "WARN", _log_c_warn, os.Stdout, m, args...)
}

func (l *logger) Error(m string, args ...interface{}) {
	l.log(ERROR, "ERRO", _log_c_error, os.Stderr, m, args...)
}

func (l *logger) isLogging(lev LogLevel) bool {
	if !l.Logging {
		return false
	}

	if l.Level > lev && l.Level != DEBUG {
		return false
	}

	return true
}

func (l *logInfo) log() {
	mes := l.message

	if len(l.field) > 0 {
		mes = fmt.Sprintf("%s\t", mes)
		for k, v := range l.field {
			mes = fmt.Sprintf("%s\x1b[%dm[%s]: %+v\x1b[0m  ", mes, _log_c_field, k, v)
		}
	}

	step := strings.Repeat(l.separate, l.step)

	fmt.Fprintf(l.output, "\x1b[%d;1m[%s]\x1b[0m %s%s\n", l.levelColor, l.level, step, mes)
}

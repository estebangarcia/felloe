package js

import (
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
	"strings"
)

type jsConsole struct{
	logger *log.Logger
}

func NewJSConsole(logger *log.Logger) *jsConsole {
	return &jsConsole{
		logger: logger,
	}
}

func (c jsConsole) log(level log.Level, msg goja.Value, args ...goja.Value) {
	m := msg.String()
	for _, v := range args {
		m = strings.Join([]string{m, v.String()}, " ")
	}

	switch level {
	case log.DebugLevel:
		c.logger.Debug(msg)
	case log.InfoLevel:
		c.logger.Info(msg)
	case log.WarnLevel:
		c.logger.Warn(msg)
	case log.ErrorLevel:
		c.logger.Error(msg)
	case log.FatalLevel:
		c.logger.Fatal(msg)
	case log.TraceLevel:
		c.logger.Trace(msg)
	}
}

func (c jsConsole) Log(msg goja.Value, args ...goja.Value) {
	c.Info(msg, args...)
}

func (c jsConsole) Info(msg goja.Value, args ...goja.Value) {
	c.log(log.InfoLevel, msg, args...)
}

func (c jsConsole) Debug(msg goja.Value, args ...goja.Value) {
	c.log(log.DebugLevel, msg, args...)
}

func (c jsConsole) Warn(msg goja.Value, args ...goja.Value) {
	c.log(log.WarnLevel, msg, args...)
}

func (c jsConsole) Error(msg goja.Value, args ...goja.Value) {
	c.log(log.ErrorLevel, msg, args...)
}


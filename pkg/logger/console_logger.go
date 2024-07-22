package logger

import (
	"github.com/sirupsen/logrus"
)

type ConsoleLogger struct{}

func New() *ConsoleLogger{
	return &ConsoleLogger{}
}


func (cl *ConsoleLogger) Debugf(format string, args ...interface{}){
	logrus.Debugf(format,args...)
}
func (cl *ConsoleLogger) Errorf(format string, args ...interface{}){
	logrus.Errorf(format,args...)
}
func (cl *ConsoleLogger) Infof(format string, args ...interface{}){
	logrus.Infof(format,args...)
}
func (cl *ConsoleLogger) Warnf(format string, args ...interface{}){
	logrus.Warnf(format,args...)
}
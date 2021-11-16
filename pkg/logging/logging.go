package logging

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func newLogger() *Logger{
	return &Logger{
    log : logrus.New(),
	}

}

func (l *Logger) setLevel(){
	l.log.Formatter = new(logrus.JSONFormatter)
}


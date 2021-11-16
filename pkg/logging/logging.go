package logging

import (
	"github.com/sirupsen/logrus"
	"github.com/sadetskayfu/rest-golang/internal/logger"
)

type Logger struct {
	log *logrus.logger
}

func newLogger() *Logger{
	return &Logger{
    log : logrus.new()
	}

}

func (l *Logger) setLevel(){
	l.log.Formatter = new(logrus.JSONFormatter)
}
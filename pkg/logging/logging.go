package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	Log *logrus.Logger
	Fields *logrus.Fields
}

func NewLogger() *Logger{
	return &Logger{
    Log : logrus.New(),
	Fields: &logrus.Fields{},
	}
}

func (l *Logger) SetLevel(){
	l.Log.SetOutput(os.Stdout)
	l.Log.Formatter = new(logrus.JSONFormatter)
	l.Log.Level = logrus.InfoLevel

	// logCtx := log.Log.WithField("a", "message here")
	// logCtx.Info("Message info")
}


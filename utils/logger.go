package utils

import (
	"os"

	formatter "github.com/cenk1cenk2/ci-cd-pipes/utils/formatter"
	"github.com/sirupsen/logrus"
)

// Log Returns a new logrus logger instance.
var Log = logrus.New()

// InitiateLogger the default logger
func InitiateLogger(level logrus.Level) {
	Log.Out = os.Stdout

	Log.SetFormatter(&formatter.Formatter{
		FieldsOrder:      []string{"context"},
		TimestampFormat:  "",
		HideKeys:         true,
		NoColors:         false,
		NoFieldsColors:   false,
		NoFieldsSpace:    false,
		ShowFullLevel:    false,
		NoUppercaseLevel: false,
		TrimMessages:     true,
		CallerFirst:      false,
	})

	Log.SetLevel(level)
}

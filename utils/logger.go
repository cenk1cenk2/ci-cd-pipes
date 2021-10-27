package utils

import (
	"fmt"
	"os"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

// Log Returns a new logrus logger instance.
var Log = logrus.New()

// InitiateLogger the default logger
func InitiateLogger(level logrus.Level) {
	Log.Out = os.Stdout

	Log.SetFormatter(&nested.Formatter{
		FieldsOrder: []string{"context"},
		// TimestampFormat:  "[20060102-15:04:05]",
		TimestampFormat:  "[15:04:05]",
		HideKeys:         true,
		NoColors:         false,
		NoFieldsColors:   false,
		NoFieldsSpace:    false,
		ShowFullLevel:    true,
		NoUppercaseLevel: false,
		TrimMessages:     true,
		CallerFirst:      true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			return fmt.Sprintf("%s", frame.Func.Name())
		},
	})

	Log.SetLevel(level)
}

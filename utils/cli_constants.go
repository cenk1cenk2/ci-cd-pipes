package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var CliDefaultFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "utils.ci",
		Usage:   "Indicates this is running inside a CI/CD environment to act accordingly.",
		EnvVars: []string{"CI"},
	},
	&cli.StringFlag{
		Name:    "utils.debug",
		Usage:   "Set the log level debug for the application.",
		EnvVars: []string{"DEBUG", "PLUGIN_DEBUG"},
	},
	&cli.StringFlag{
		Name:    "utils.log",
		Usage:   "Define the log level for the application.",
		EnvVars: []string{"LOG_LEVEL", "PLUGIN_LOG_LEVEL"},
		Value:   logrus.InfoLevel.String(),
	},
}

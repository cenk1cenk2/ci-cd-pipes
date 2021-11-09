package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var CliDefaultFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "utils.ci",
		Usage:  "Indicates this is running inside a CI/CD environment to act accordingly.",
		EnvVar: "CI",
	},
	cli.StringFlag{
		Name:   "utils.debug",
		Usage:  "Set the log level debug for the application.",
		EnvVar: "DEBUG,PLUGIN_DEBUG",
	},
	cli.StringFlag{
		Name:   "utils.log",
		Usage:  "Define the log level for the application.",
		EnvVar: "LOG_LEVEL,PLUGIN_LOG_LEVEL",
		Value:  logrus.InfoLevel.String(),
	},
}

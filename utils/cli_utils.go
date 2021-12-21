package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func CliLoadEnvironment() {
	if env := os.Getenv("ENV_FILE"); env != "" {
		godotenv.Load(env)
	}
}

func CliBeforeFunction(c *cli.Context) error {
	level, err := logrus.ParseLevel(c.String("utils.log"))

	if err != nil {
		fmt.Println(fmt.Sprintf("Log level is not valid with %s, using default.", level))

		level = logrus.InfoLevel
	}

	if c.String("utils.debug") != "" {
		level = logrus.DebugLevel
	}

	InitiateLogger(level)

	return nil
}

func CliGreet(c *cli.Context) {
	fmt.Println(fmt.Sprintf("%s - %s", c.App.Name, c.App.Version))
	fmt.Println(strings.Repeat("-", 20))
}

func CliRun(app *cli.App) {
	if err := app.Run(os.Args); err != nil {
		Log.Fatal(err)
	}
}

func CliCreate(options *cli.App) {
	CliLoadEnvironment()

	app := options

	CliRun(app)
}

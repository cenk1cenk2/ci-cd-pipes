package main

import (
	"github.com/urfave/cli/v2"

	pipe "github.com/cenk1cenk2/ci-cd-pipes/git-push/pipe"
	utils "github.com/cenk1cenk2/ci-cd-pipes/utils"
)

func main() {
	utils.CliCreate(
		utils.CliCreateArgs{
			Name:    pipe.CLI_NAME,
			Version: pipe.VERSION,
			Run:     run,
			Flags:   pipe.Flags,
		},
	)
}

func run(c *cli.Context) error {
	utils.CliGreet(c)

	return pipe.Pipe.Exec()
}

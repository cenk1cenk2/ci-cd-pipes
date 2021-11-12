package pipe

import (
	utils "github.com/cenk1cenk2/ci-cd-pipes/utils"
	"github.com/urfave/cli/v2"
)

type (
	Markdown struct {
		Patterns  cli.StringSlice
		Arguments string
	}

	Plugin struct {
		Markdown Markdown
	}
)

var Pipe Plugin = Plugin{}

func (p Plugin) Exec() error {
	utils.AddTasks(
		[]utils.Task{
			TaskFindMarkdownFiles(),
			TaskRunMarkdownToc(),
		},
	)

	utils.RunAllTasks(utils.DefaultRunAllTasksOptions)

	return nil
}

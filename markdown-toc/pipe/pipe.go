package pipe

import (
	utils "github.com/cenk1cenk2/ci-cd-pipes/utils"
)

type (
	Markdown struct {
		Patterns  []string
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

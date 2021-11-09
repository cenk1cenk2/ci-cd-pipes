package pipe

import (
	utils "github.com/cenk1cenk2/ci-cd-pipes/utils"
)

type (
	DockerHub struct {
		Username string
		Password string
		Address  string
	}

	Readme struct {
		Repository  string
		File        string
		Description string
	}

	Plugin struct {
		DockerHub DockerHub
		Readme    Readme
	}
)

var Pipe Plugin = Plugin{}

func (p Plugin) Exec() error {
	utils.AddTasks(
		[]utils.Task{
			TaskLoginToDockerHubRegistry(),
			TaskUpdateDockerReadme(),
		},
	)

	utils.RunAllTasks(utils.DefaultRunAllTasksOptions)

	return nil
}

package pipe

import (
	base "github.com/appleboy/drone-git-push/repo"
	utils "github.com/cenk1cenk2/ci-cd-pipes/utils"
)

type Ctx struct {
}

var Context Ctx

func TaskVerifyVariables() utils.Task {
	metadata := utils.TaskMetadata{Context: "Verify"}

	return utils.Task{Metadata: metadata, Task: func(t utils.Task) error {
		return nil
	}}
}

func TaskGitConfiguration() utils.Task {
	metadata := utils.TaskMetadata{Context: "Git Configuration"}

	return utils.Task{
		Metadata: metadata,
		Commands: []utils.Command{
			base.GlobalUser(Pipe.Commit.Author.Email),
			base.GlobalName(Pipe.Commit.Author.Name),
		},
	}
}

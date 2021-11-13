package pipe

import (
	utils "github.com/cenk1cenk2/ci-cd-pipes/utils"
)

type (
	Netrc struct {
		Machine  string
		Login    string
		Password string
	}

	Commit struct {
		Author Author
	}

	Author struct {
		Name  string
		Email string
	}

	Config struct {
		Key           string
		Remote        string
		RemoteName    string
		Branch        string
		LocalBranch   string
		Force         bool
		FollowTags    bool
		SkipVerify    bool
		Commit        bool
		CommitMessage string
		EmptyCommit   bool
		NoVerify      bool
	}

	Plugin struct {
		Netrc  Netrc
		Commit Commit
		Config Config
	}
)

var Pipe Plugin = Plugin{}

func (p Plugin) Exec() error {
	utils.AddTasks(
		[]utils.Task{},
	)

	utils.RunAllTasks(utils.DefaultRunAllTasksOptions)

	return nil
}

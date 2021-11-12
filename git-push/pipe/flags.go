package pipe

import (
	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "git.commit.author.name",
		Usage:       "git author name",
		EnvVars:     []string{"PLUGIN_AUTHOR_NAME", "DRONE_COMMIT_AUTHOR"},
		Destination: &Pipe.Commit.Author.Name,
	},
	&cli.StringFlag{
		Name:        "git.commit.author.email",
		Usage:       "git author email",
		EnvVars:     []string{"PLUGIN_AUTHOR_EMAIL", "DRONE_COMMIT_AUTHOR_EMAIL"},
		Destination: &Pipe.Commit.Author.Email,
	},
	&cli.StringFlag{
		Name:        "git.netrc.machine",
		Usage:       "netrc machine",
		EnvVars:     []string{"DRONE_NETRC_MACHINE"},
		Destination: &Pipe.Netrc.Machine,
	},
	&cli.StringFlag{
		Name:        "git.netrc.username",
		Usage:       "netrc username",
		EnvVars:     []string{"DRONE_NETRC_USERNAME"},
		Destination: &Pipe.Netrc.Login,
	},
	&cli.StringFlag{
		Name:        "git.netrc.password",
		Usage:       "netrc password",
		EnvVars:     []string{"DRONE_NETRC_PASSWORD"},
		Destination: &Pipe.Netrc.Password,
	},
	&cli.StringFlag{
		Name:        "git.ssh-key",
		Usage:       "private ssh key",
		EnvVars:     []string{"PLUGIN_SSH_KEY", "GIT_PUSH_SSH_KEY"},
		Destination: &Pipe.Config.Key,
	},
	&cli.StringFlag{
		Name:        "git.remote",
		Usage:       "url of the remote repo",
		EnvVars:     []string{"PLUGIN_REMOTE", "GIT_PUSH_REMOTE"},
		Destination: &Pipe.Config.Remote,
	},
	&cli.StringFlag{
		Name:        "git.remote.name",
		Usage:       "name of the remote repo",
		Value:       "deploy",
		EnvVars:     []string{"PLUGIN_REMOTE_NAME", "GIT_PUSH_REMOTE_NAME"},
		Destination: &Pipe.Config.RemoteName,
	},
	&cli.StringFlag{
		Name:        "git.branch",
		Usage:       "name of remote branch",
		EnvVars:     []string{"PLUGIN_BRANCH", "GIT_PUSH_BRANCH"},
		Value:       "master",
		Destination: &Pipe.Config.Branch,
	},
	&cli.StringFlag{
		Name:        "git.local-branch",
		Usage:       "name of local branch",
		Value:       "HEAD",
		EnvVars:     []string{"PLUGIN_LOCAL_BRANCH", "GIT_PUSH_LOCAL_BRANCH"},
		Destination: &Pipe.Config.LocalBranch,
	},
	&cli.StringFlag{
		Name:        "git.repo.path",
		Usage:       "path to git repo",
		EnvVars:     []string{"PLUGIN_PATH"},
		Destination: &Pipe.Config.Path,
	},
	&cli.BoolFlag{
		Name:        "git.force",
		Usage:       "force push to remote",
		EnvVars:     []string{"PLUGIN_FORCE", "GIT_PUSH_FORCE"},
		Destination: &Pipe.Config.Force,
	},
	&cli.BoolFlag{
		Name:        "git.followtags",
		Usage:       "push to remote with tags",
		EnvVars:     []string{"PLUGIN_FOLLOWTAGS", "GIT_PUSH_FOLLOWTAGS"},
		Destination: &Pipe.Config.FollowTags,
	},
	&cli.BoolFlag{
		Name:        "git.skip-verify",
		Usage:       "skip ssl verification",
		EnvVars:     []string{"PLUGIN_SKIP_VERIFY", "GIT_PUSH_SKIP_VERIFY"},
		Destination: &Pipe.Config.SkipVerify,
	},
	&cli.BoolFlag{
		Name:        "git.commit",
		Usage:       "commit dirty changes",
		EnvVars:     []string{"PLUGIN_COMMIT", "GIT_PUSH_COMMIT"},
		Destination: &Pipe.Config.Commit,
	},
	&cli.StringFlag{
		Name:        "git.commit-message",
		Usage:       "commit message",
		EnvVars:     []string{"PLUGIN_COMMIT_MESSAGE", "GIT_PUSH_COMMIT_MESSAGE"},
		Destination: &Pipe.Config.CommitMessage,
	},
	&cli.BoolFlag{
		Name:        "git.empty-commit",
		Usage:       "empty commit",
		EnvVars:     []string{"PLUGIN_EMPTY_COMMIT", "GIT_PUSH_EMPTY_COMMIT"},
		Destination: &Pipe.Config.EmptyCommit,
	},
	&cli.BoolFlag{
		Name:        "git.no-verify",
		Usage:       "bypasses the pre-commit and commit-msg hooks",
		EnvVars:     []string{"PLUGIN_NO_VERIFY", "GIT_PUSH_NO_VERIFY"},
		Destination: &Pipe.Config.NoVerify,
		Value:       true,
	},
}

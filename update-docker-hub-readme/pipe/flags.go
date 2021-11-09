package pipe

import (
	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:        "docker_hub.username",
		Usage:       "Docker Hub username for updating the readme.",
		EnvVar:      "DOCKER_USERNAME",
		Required:    true,
		Destination: &Pipe.DockerHub.Username,
	},
	cli.StringFlag{
		Name:        "docker_hub.password",
		Usage:       "Docker Hub password for updating the readme.",
		EnvVar:      "DOCKER_PASSWORD",
		Required:    true,
		Destination: &Pipe.DockerHub.Password,
	},
	cli.StringFlag{
		Name:        "docker_hub.address",
		Usage:       "HTTP address for the docker hub. There is only one!",
		EnvVar:      "DOCKER_HUB_ADDRESS",
		Value:       "https://hub.docker.com/v2/repositories",
		Destination: &Pipe.DockerHub.Address,
	},
	cli.StringFlag{
		Name:        "readme.repository",
		Usage:       "Repository for applying the readme on.",
		EnvVar:      "README_REPOSITORY",
		Required:    true,
		Destination: &Pipe.Readme.Repository,
	},
	cli.StringFlag{
		Name:        "readme.file",
		Usage:       "Readme file for the given repossitory.",
		EnvVar:      "README_FILE",
		Value:       "README.md",
		Destination: &Pipe.Readme.File,
		Required:    false,
	},
	cli.StringFlag{
		Name:        "readme.short_description",
		Usage:       "Pass in description to send it in the request.",
		EnvVar:      "README_DESCRIPTION",
		Destination: &Pipe.Readme.Description,
		Required:    false,
	},
}

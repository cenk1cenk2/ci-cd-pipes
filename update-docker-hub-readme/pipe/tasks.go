package pipe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	utils "github.com/cenk1cenk2/ci-cd-pipes/utils"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var DockerClient *client.Client

type Ctx struct {
	token string
}

var Context Ctx

var ctx = context.Background()

func TaskLoginToDockerHubRegistry() utils.Task {
	metadata := utils.TaskMetadata{Context: "docker hub - login"}

	return utils.Task{Metadata: metadata, Task: func(t utils.Task) error {
		log := utils.Log.WithField("context", t.Metadata.Context)

		login, err := json.Marshal(types.AuthConfig{
			Username: Pipe.DockerHub.Username,
			Password: Pipe.DockerHub.Password,
		})

		if err != nil {
			return err
		}

		res, err := http.Post(
			"https://hub.docker.com/v2/users/login/",
			JSON_REQUEST,
			bytes.NewBuffer(login),
		)

		if err != nil {
			return err
		}

		log.Debugln("Authentication token obtained.")

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return err
		}

		b := DockerHubLoginResponse{}
		err = json.Unmarshal(body, &b)

		Context.token = b.Token

		return nil

	}}
}

func TaskUpdateDockerReadme() utils.Task {
	metadata := utils.TaskMetadata{Context: "docker hub - update readme"}

	return utils.Task{Metadata: metadata, Task: func(t utils.Task) error {
		log := utils.Log.WithField("context", t.Metadata.Context)

		log.Debugln(fmt.Sprintf("Trying to readfile %s...", Pipe.Readme.File))

		content, err := ioutil.ReadFile(Pipe.Readme.File)

		if err != nil {
			return err
		}

		readme := string(content)

		log.Debugln(fmt.Sprintf("File read: %s", Pipe.Readme.File))
		log.Debugln(fmt.Sprintf("Running against repository: %s", Pipe.Readme.Repository))
		log.Debugln(fmt.Sprintf("Running against registry: %s", Pipe.DockerHub.Registry))

		body, err := json.Marshal(
			DockerHubUpdateReadmeRequest{
				Registry: Pipe.DockerHub.Registry,
				Readme:   readme,
			},
		)

		if err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPatch,
			fmt.Sprintf("%s/%s", Pipe.DockerHub.Address, Pipe.Readme.Repository),
			bytes.NewBuffer(body),
		)

		req = addAuthenticationHeadersToRequest(req)

		if err != nil {
			return err
		}

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			return err
		}

		log.Debugln(fmt.Sprintf("Status Code: %d", res.StatusCode))

		defer res.Body.Close()

		body, err = ioutil.ReadAll(res.Body)

		if err != nil {
			return err
		}

		b := DockerHubUpdateReadmeResponse{}
		err = json.Unmarshal(body, &b)

		if err != nil {
			return err
		}

		switch res.StatusCode {
		case 200:
			if b.FullDescription != readme {
				log.Fatalln("Uploaded README does not match with current repository README file.")
			}

			log.Infoln(
				fmt.Sprintf(
					"Successfully pushed readme file to: %s -> %s/%s",
					Pipe.Readme.File,
					Pipe.DockerHub.Address,
					Pipe.Readme.Repository,
				),
			)
		case 404:
			log.Fatalln(
				fmt.Sprintf(
					"Repository does not exists: %s/%s",
					Pipe.DockerHub.Address,
					Pipe.Readme.Repository,
				),
			)
		default:
			log.Fatalln(
				fmt.Sprintf(
					"Pushing readme failed with code: %d",
					res.StatusCode,
				),
			)
		}

		return nil
	}}
}

func addAuthenticationHeadersToRequest(req *http.Request) *http.Request {
	req.Header.Add("User-Agent", CLI_NAME)
	req.Header.Add("Authorization", fmt.Sprintf("JWT %s", Context.token))
	req.Header.Add("Content-Type", JSON_REQUEST)

	return req
}

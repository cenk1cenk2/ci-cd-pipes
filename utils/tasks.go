package utils

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
)

type (
	TaskMetadata struct {
		Context string
		Skip    bool
	}

	Task struct {
		Command  *exec.Cmd
		Task     func(Task) error
		Metadata TaskMetadata
	}

	RunAllTasksOptions struct {
		Sync bool
	}
)

var TaskList []Task = []Task{}

func AddTask(command Task) []Task {
	TaskList = append(TaskList, command)

	return TaskList
}

func AddTasks(commands []Task) []Task {
	TaskList = append(TaskList, commands...)

	return TaskList
}

var DefaultRunAllTasksOptions = RunAllTasksOptions{Sync: true}

func RunAllTasks(options RunAllTasksOptions) {
	if len(TaskList) == 0 {
		Log.WithField("context", "COMMAND").Fatalln("Task list is empty!")
	}

	if options.Sync == true {
		for _, task := range TaskList {
			if task.Task != nil {
				err := task.Task(task)

				if err != nil {
					Log.WithField("context", "FAILED").
						Fatalln(fmt.Sprintf("$ Task > %s", err))
				}
			}

			if task.Metadata.Skip != true {
				if task.Command != nil {
					cmd := strings.Join(task.Command.Args, " ")
					Log.WithField("context", "RUN").
						Infoln(fmt.Sprintf("$ %s", cmd))

					task.Command.Args = deleteEmptyStrings(task.Command.Args)

					err := ExecuteAndPipeToLogger(task.Command, task.Metadata)

					if err != nil {
						Log.WithField("context", "FAILED").
							Fatalln(fmt.Sprintf("$ %s > %s", cmd, err))
					} else {
						Log.WithField("context", "FINISH").Infoln(fmt.Sprintf("%s", cmd))
					}
				}
			} else {
				Log.Warnln(fmt.Sprintf("Task skipped: %s", task.Metadata.Context))
			}

		}
	} else {
		Log.Fatalln("Not implemented yet!")
	}

	TaskList = []Task{}
}

func ExecuteAndPipeToLogger(cmd *exec.Cmd, context TaskMetadata) error {
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		Log.Fatalln("Failed creating command stdout pipe: ", err)
	}

	defer stdout.Close()
	stdoutReader := bufio.NewReader(stdout)

	stderr, err := cmd.StderrPipe()

	if err != nil {
		Log.Fatalln("Failed creating command stderr pipe: ", err)
	}

	defer stderr.Close()
	stderrReader := bufio.NewReader(stderr)

	if err := cmd.Start(); err != nil {
		Log.Fatalln("Failed starting command: ", err)
	}

	go handleReader(stdoutReader, context)
	go handleReader(stderrReader, context)

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				Log.Debugln("Exit Status: ", status.ExitStatus())
			}
		}
		return err
	}

	return nil
}

func handleReader(reader *bufio.Reader, context TaskMetadata) {
	var log *logrus.Entry = Log.WithFields(logrus.Fields{})

	if context.Context != "" {
		log = log.WithField("context", context.Context)
	}

	for {
		str, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		log.Infoln(str)
	}
}

func deleteEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

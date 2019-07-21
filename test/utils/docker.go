package testutils

import (
	"os/exec"
	"strings"
	"testing"
)

const (
	dockerCliBin        = "docker"
	DockerKillCommand   = "kill"
	DockerRunCommand    = "run"
	dockerCliLabelFlag  = "-l"
	dockerCliDetachFlag = "-d"
	dockerTTYFlag       = "-t"
	spc                 = " "
)

type DockerCliArgs struct {
	args       strings.Builder
	isImageSet bool
}

func NewDockerCliBuilder(command string) *DockerCliArgs {
	cliBuilder := &DockerCliArgs{
		strings.Builder{},
		false,
	}

	cliBuilder.args.WriteString(command + spc)
	cliBuilder.args.WriteString(dockerCliDetachFlag + spc)
	cliBuilder.args.WriteString(dockerTTYFlag + spc)

	return cliBuilder
}

func (d *DockerCliArgs) Label(label string) *DockerCliArgs {
	if !d.isImageSet {
		d.args.WriteString(dockerCliLabelFlag +
			spc +
			label +
			spc)
	}

	return d
}

func (d *DockerCliArgs) Image(image string) *DockerCliArgs {
	if !d.isImageSet {
		d.isImageSet = true
		d.args.WriteString(image)
	}

	return d
}

func (d *DockerCliArgs) Command(cmd string) *DockerCliArgs {
	if d.isImageSet {
		d.args.WriteString(spc + cmd)
	}

	return d
}

func (d *DockerCliArgs) build() []string {
	buildString := d.args.String()
	return strings.Split(buildString, spc)
}

// Launch a container from CLI with the flag passed
func DockerCli(flags *DockerCliArgs, fn func(string), t *testing.T) {
	containerId, err := exec.Command(dockerCliBin, flags.build()...).Output()
	if err != nil {
		t.Fatalf(err.Error())
	}

	sanitizedContainerId := strings.Trim(string(containerId), "\n")

	fn(sanitizedContainerId)

	KillContainer(sanitizedContainerId, t)
}

func KillContainer(containerId string, t *testing.T) {
	cmd := exec.Command(dockerCliBin, DockerKillCommand, containerId)
	err := cmd.Run()

	if err != nil {
		t.Fatalf(err.Error())
	}
}

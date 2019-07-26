package testutils

import (
	"os/exec"
	"strings"
	"testing"
)

const (
	dockerCliBin         = "docker"
	DockerKillCommand    = "kill"
	DockerRunCommand     = "run"
	DockerServiceCommand = "service"
	DockerCreateCommand  = "create"
	DockerRmCommand      = "rm"
	dockerCliLabelFlag   = "-l"
	DockerCliDetachFlag  = "-d"
	DockerTTYFlag        = "-t"
	dockerNameFlag       = "--name"
	spc                  = " "
)

type DockerCliArgs struct {
	args       strings.Builder
	isImageSet bool
	isService  bool
}

func NewDockerCliBuilder(command string) *DockerCliArgs {
	cliBuilder := &DockerCliArgs{
		strings.Builder{},
		false,
		false,
	}

	cliBuilder.args.WriteString(command + spc)
	return cliBuilder
}

func (d *DockerCliArgs) Flag(flag string) *DockerCliArgs {
	if !d.isImageSet {
		d.args.WriteString(flag + spc)
	}

	return d
}

func (d *DockerCliArgs) Name(name string) *DockerCliArgs {
	if !d.isImageSet {
		d.args.WriteString(dockerNameFlag +
			spc +
			name +
			spc)
	}

	return d
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
	if !d.isImageSet {
		d.args.WriteString(cmd + spc)

		if cmd == DockerCreateCommand {
			d.isService = true
		}
	}

	return d
}

func (d *DockerCliArgs) build() []string {
	buildString := d.args.String()
	return strings.Split(buildString, spc)
}

// Launch a container from CLI with the flag passed
func DockerCli(flags *DockerCliArgs, fn func(string), t *testing.T) {
	deploymentId, err := exec.Command(dockerCliBin, flags.build()...).Output()
	if err != nil {
		t.Fatalf(err.Error())
	}
	sanitizedDeploymentId := strings.Trim(string(deploymentId), "\n")

	fn(sanitizedDeploymentId)

	if flags.isService {
		KillService(sanitizedDeploymentId, t)
	} else {
		KillContainer(sanitizedDeploymentId, t)
	}
}

func KillContainer(containerId string, t *testing.T) {
	cmd := exec.Command(dockerCliBin, DockerKillCommand, containerId)
	err := cmd.Run()

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func KillService(serviceId string, t *testing.T) {
	cmd := exec.Command(dockerCliBin, DockerServiceCommand, DockerRmCommand, serviceId)
	err := cmd.Run()

	if err != nil {
		t.Fatalf(err.Error())
	}
}

package gantry

import (
	"fmt"
	"strings"
	"os"
	"github.com/docker/docker/pkg/stringutils"
)

func StartCommand(args []string) {
	project := getProjectName()
	if (getProjectRunning(project)) {
		fmt.Printf("Restarting project: %s\n", project)
		runningCont := getNewestProjectMainContainers(project);
		var runningPort int;
		for _, port := range runningCont.Ports {
			runningPort = port.PublicPort
		}
		fmt.Printf("Current port: %d\n", runningPort)
		// Flip Flop between DOCKER_HTTP_PORT and DOCKER_HTTP_PORT+1
		mainPort := getHTTPPort()
		newPort := mainPort
		if (runningPort == mainPort) {
			newPort++;
		}
		env := os.Environ()
		env = append(env, fmt.Sprintf("DOCKER_HTTP_PORT=%d", newPort))

		fmt.Printf("Building: %s\n", project)

		args := getComposeFilesArgs()
		args = append(args, "build")
		args = append(args, "main")

		fmt.Printf("docker-compose %s\n", strings.Join(args, " "));

		process := exec.Command("docker-compose", args...)
		process.Env = env
		process.Stdout = os.Stdout
		process.Stderr = os.Stderr
		process.Run()
		process.Wait()

		fmt.Printf("Scaling: %s:main\n", project)
		args = getComposeFilesArgs()
		args = append(args, "scale")
		args = append(args, "main=2")
		fmt.Printf("docker-compose %s\n", strings.Join(args, " "));

		process = exec.Command("docker-compose", args...)
		process.Env = env
		process.Stdout = os.Stdout
		process.Stderr = os.Stderr
		process.Run()
		process.Wait()

		fmt.Printf("Stopping: %s\n", runningCont.Names)
		stopContainer(runningCont);

	} else {
		fmt.Printf("Starting project: %s\n", project)
		process := exec.Command("docker-compose", stringutils.ShellQuoteArguments(append(getComposeFilesArgs(), "ip", "-d")))
		process.Start()
		process.Wait()
	}
}
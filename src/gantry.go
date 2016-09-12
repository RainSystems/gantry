package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/docker/docker/client"
	"os"
	"fmt"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"time"
	"github.com/docker/docker/api/types/filters"
	"strings"
	"path"
	"regexp"
	"os/exec"
	"github.com/docker/docker/pkg/stringutils"
	"strconv"
)

func main() {

	argsWithoutProg := os.Args[1:]

	var opts struct {
		// Slice of bool will append 'true' each time the option
		// is encountered (can be set multiple times, like -vvv)
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
		Version bool `long:"version" description:"Show version information"`
	}

	argsWithoutProg , err := flags.ParseArgs(&opts, argsWithoutProg )

	if err != nil {
		panic(err)
		os.Exit(1)
	}

	if(opts.Version) {
		fmt.Println("Gantry 1.4")
		os.Exit(0)
	}

	if(len(argsWithoutProg) > 0) {

		switch argsWithoutProg[0] {
		case "start":
			start(argsWithoutProg[1:])
		}

	}


}

func getProjectName() string {
	project := os.Getenv("GANTRY_PROJECT");
	if len(project) > 0 {
		return project
	}
	project = os.Getenv("COMPOSE_PROJECT_NAME");
	if len(project) > 0 {
		return project
	}
	project, err := os.Getwd();
	if err == nil {
		re := regexp.MustCompile("[^a-z]")
		return re.ReplaceAllString(strings.ToLower(path.Base(project)),"");
	}
	return ""
}
func getApplicationEnv() string {
	return os.Getenv("APP_ENV");
}
func getHTTPPort() int {
	port, _ := strconv.ParseInt(os.Getenv("DOCKER_HTTP_PORT"), 10, 32);
	return (int)(port)
}
func getProjectBasePath() string {
	return os.Getenv("PWD");
}

func getComposeFilesArgs() []string {
	var composeArgs []string
	base := getProjectBasePath()
	compose := path.Join(base,"docker-compose.yml")
	if _, err := os.Stat(compose); err == nil {
		composeArgs = append(composeArgs, "-f", compose)
	}
	env := getApplicationEnv()
	compose = path.Join(base,fmt.Sprintf("docker-compose-%s.yml", env))
	if _, err := os.Stat(compose); err == nil {
		composeArgs = append(composeArgs, "-f", compose)

	}
	return composeArgs
}

func start(args []string) {
	project := getProjectName()
	if(getProjectRunning(project)) {
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
		if(runningPort == mainPort) {
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

func getProjectRunning(project string) bool {
	running := getProjectContainers(project);
	if len(running) > 0 {
		return true;
	}
	return false;
}

func getProjectContainers(project string) []types.Container {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	listOpts := types.ContainerListOptions{}
	listOpts.Filter = filters.NewArgs();
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.project=%s",project))

	containers, _ := dockerClient.ContainerList(ctx, listOpts)
	return containers
}

func stopContainer(container types.Container) {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	timeout := 10*time.Second;

	dockerClient.ContainerStop(ctx, container.ID, &timeout)
}

func getProjectMainContainers(project string) []types.Container {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	listOpts := types.ContainerListOptions{}
	listOpts.Filter = filters.NewArgs();
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.project=%s",project))
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.service=%s","main"))

	containers, _ := dockerClient.ContainerList(ctx, listOpts)
	return containers
}

func getNewestProjectMainContainers(project string) types.Container {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	listOpts := types.ContainerListOptions{}
	listOpts.Filter = filters.NewArgs();
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.project=%s",project))
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.service=%s","main"))

	containers, _ := dockerClient.ContainerList(ctx, listOpts)
	newestDate := (int64)(0);
	var newestContainer types.Container;
	for _, c := range containers {
		if c.Created > newestDate {
			newestDate = c.Created
			newestContainer = c
		}
	}
	return newestContainer
}
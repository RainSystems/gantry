package gantry

import (
	"github.com/docker/docker/api/types/filters"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stringutils"
	"golang.org/x/net/context"


)

func getProjectRunning(project string) bool {
	running := getProjectContainers(project);
	if len(running) > 0 {
		return true;
	}
	return false;
}

func getProjectContainers(project string) []types.Container {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	listOpts := types.ContainerListOptions{}
	listOpts.Filter = filters.NewArgs();
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.project=%s", project))

	containers, _ := dockerClient.ContainerList(ctx, listOpts)
	return containers
}

func stopContainer(container types.Container) {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	timeout := 10 * time.Second;

	dockerClient.ContainerStop(ctx, container.ID, &timeout)
}

func getProjectMainContainers(project string) []types.Container {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	listOpts := types.ContainerListOptions{}
	listOpts.Filter = filters.NewArgs();
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.project=%s", project))
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.service=%s", "main"))

	containers, _ := dockerClient.ContainerList(ctx, listOpts)
	return containers
}

func getNewestProjectMainContainers(project string) types.Container {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	listOpts := types.ContainerListOptions{}
	listOpts.Filter = filters.NewArgs();
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.project=%s", project))
	listOpts.Filter.Add("label", fmt.Sprintf("com.docker.compose.service=%s", "main"))

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
func dockerTagPush(imageID string, tag string) {
	dockerClient, _ := client.NewEnvClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	dockerClient.ImageTag(ctx, imageID, tag)

	pushOptions := types.ImagePushOptions{}
	dockerClient.ImagePush(ctx, tag, pushOptions)
}


func dockerLogin() {
	auth := types.AuthConfig{
	}

	dockerClient, _ := client.NewEnvClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	dockerClient.RegistryLogin(ctx, auth)

}
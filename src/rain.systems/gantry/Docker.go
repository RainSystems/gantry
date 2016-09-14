package gantry

import (
	"fmt"
	"time"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"log"
	"os"
)

func getProjectRunning(project string) bool {
	running := getProjectContainers(project);
	if len(running) > 0 {
		return true;
	}
	return false;
}

type ErrImageNotFound struct {
	image string
}
func (i ErrImageNotFound) Error() string {
	return fmt.Sprintf("Image %s not Fround", i)
}

func getImageByName(name string) (types.Image, error) {
	dockerClient, _ := client.NewEnvClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	listOpts := types.ImageListOptions{}
	listOpts.MatchName = name

	images, _ := dockerClient.ImageList(ctx, listOpts)
	if len(images) > 0 {
		return images[0], nil
	}
	return types.Image{}, &ErrImageNotFound{name}
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
func dockerTagPush(client *client.Client, token string, imageID string, tag string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err := client.ImageTag(ctx, imageID, tag)
	if err != nil {
		log.Fatalf("Tag Error: %s\n", err)
		os.Exit(1)
	}

	pushOptions := types.ImagePushOptions{
		RegistryAuth:token,
	}
	_, err = client.ImagePush(ctx, tag, pushOptions)
	if err != nil {
		log.Fatalf("Push Error: %s\n", err)
		os.Exit(1)
	}
}


func dockerLogin(token *string) (*client.Client, string) {
	auth := types.AuthConfig{
		RegistryToken:*token,
	}

	dockerClient, _ := client.NewEnvClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	authResp, err := dockerClient.RegistryLogin(ctx, auth)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Docker Login: %s\n", authResp.Status)
	return dockerClient, authResp.IdentityToken
}
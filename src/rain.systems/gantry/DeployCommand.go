package gantry

import (
	"log"
	"os"
)

type DeployCommand struct {
	Env string `short:"e" long:"env" default:"prod" description:"Deploy enviroment"`
}

func (deploy *DeployCommand) Execute(args []string) error {

	conf := LoadConfig()

	client,token := dockerLogin(EcrLogin(conf))
	image, err := getImageByName(conf.Project.Name)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	dockerTagPush(client, token, image.ID, conf.Docker.Tag)

	return nil
}
package gantry

import (
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
	"log"
)

// See configuration-sample.yml
type Config struct {
	Vcs VcsConfig
	Docker DockerConfig
	Aws AwsConfig
	Project ProjectConfig
}
type AwsConfig struct {
	Region string   `yaml:"region"`
	Config string   `yaml:"config"`
	Credentials string   `yaml:"credentials"`
	Profile string   `yaml:"profile"`
	AccessId string   `yaml:"accessId"`
	SecretKey string   `yaml:"secretKey"`
}
type VcsConfig struct {
	GitHub string   `yaml:"github"`
	Aws string   `yaml:"aws"`
}
type DockerConfig struct {
	Tag string   `yaml:"tag"`
}
type ProjectConfig struct {
	Label string   `yaml:"label"`
	Name string   `yaml:"name"`
}

func LoadConfig() Config {
	dat, err := ioutil.ReadFile(Options.ConfigFile)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	conf := Config{
		Aws: AwsConfig{
			Config: "~/.aws/config",
			Credentials: "~/.aws/credentials",
		},
	}

	err = yaml.Unmarshal(dat, &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return conf
}

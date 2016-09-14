package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"gantry"
)

// See configuration-sample.yml
type Config struct {
	Vcs struct {
		    GitHub string   `yaml:"github"`
	    }
	Aws AwsConfig
	Project ProjectConfig
}
type AwsConfig struct {
	Region string   `yaml:"region"`
	Config string   `yaml:"config"`
	Credentials string   `yaml:"credentials"`
	Profile string   `yaml:"profile"`
}
type ProjectConfig struct {
	Label string   `yaml:"label"`
	Name string   `yaml:"name"`
}

func main() {

	argsWithoutProg := os.Args[1:]

	var opts struct {
		// Slice of bool will append 'true' each time the option
		// is encountered (can be set multiple times, like -vvv)
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
		Version bool `long:"version" description:"Show version information"`
		ConfigFile string `short:"c" long:"config" description:"Show version information" default:"gantry.yml"`

		Start gantry.StartCommand `long:"start"`
		New gantry.StartCommand `long:"start"`
		Deploy gantry.StartCommand `long:"start"`
	}

	argsWithoutProg, err := flags.ParseArgs(&opts, argsWithoutProg)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(opts.ConfigFile)
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
	fmt.Printf("git: %s\nregion:%s\n\n", conf.Vcs.GitHub, conf.Aws.Region)
	os.Exit(0)

}

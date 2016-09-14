package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"rain.systems/gantry"
)



func main() {

	parser := flags.NewParser(&gantry.Options, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok {
			if e.Type == flags.ErrHelp {
				os.Exit(0)
			}
			if e.Type == flags.ErrCommandRequired {
				os.Exit(1)
			}
		}
		panic(err)
		os.Exit(1)
	}

	//dat, err := ioutil.ReadFile(opts.ConfigFile)
	//if err != nil {
	//	panic(err)
	//	os.Exit(1)
	//}
	//
	//conf := Config{
	//	Aws: AwsConfig{
	//		Config: "~/.aws/config",
	//		Credentials: "~/.aws/credentials",
	//	},
	//}
	//
	//err = yaml.Unmarshal(dat, &conf)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("git: %s\nregion:%s\n\n", conf.Vcs.GitHub, conf.Aws.Region)
	//os.Exit(0)

}

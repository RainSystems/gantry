package gantry

type opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Version bool `long:"version" description:"Show version information"`
	ConfigFile string `short:"c" long:"config" description:"Show version information" default:"gantry.yml"`

	Start StartCommand 	`command:"start" description:"Start App"`
	New NewCommand 	`command:"new" description:"Create New App"`
	Deploy DeployCommand 	`command:"deploy" description:"Deploy App"`
}

var Options opts
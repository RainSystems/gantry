package gantry

import (
	"fmt"
)

type NewCommand struct {
	Symfony NewSymfonyCommand `command:"symfony" description:"New Symfony Application"`
	WordPress NewWordPressCommand `command:"wordpress" description:"New Symfony Application"`
}


func (command *NewCommand) Execute(args []string) {

}

type NewSymfonyCommand struct {}
type NewWordPressCommand struct {}


func (command *NewSymfonyCommand) Execute(args []string) {
	fmt.Println("Setup up Symfony Enviroment")
}
func (command *NewWordPressCommand) Execute(args []string) {
	fmt.Println("Setup up Wordpress Enviroment")
}

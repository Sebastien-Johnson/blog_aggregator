package main

import (
	"log"
	"os"

	"github.com/Sebastien-Johnson/blog_aggregator/internal/config"
)

//holds state of one 'config'/user
type state struct{
	cfg *config.Config
}

func main() {
	//Creates user from reading cfg file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//sets up state with read config
	programState := state{
		cfg: &cfg,
	}

	//sets up commands struct & map
	cmds := commands{
		registeredComms: make(map[string]func(*state, command) error),
	}
	
	cmds.Register("login", handlerLogin)
	userArgs := os.Args

	if len(userArgs) < 2 {
		log.Fatalf("Not enough commands entered")
	}
	//grabs command from user input
	userCmd := userArgs[1]
	//grabs remaining args from user input
	userCmdArg := userArgs[2:]

	//generates command struct from user input
	userCmdStruct := command{
		name : userCmd,
		args : userCmdArg,
	}

	//runs command struct from user input
	err = cmds.Run(&programState, userCmdStruct)
	if err != nil {
		log.Fatalf("Err: %s", err)
	}
}

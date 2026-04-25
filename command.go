package main

import "errors"
//holds params for single command
type command struct{
	name string //ex: 'login'
	args []string //ex: 'username'
}

//holds map of commands
type commands struct{
	registeredComms map[string]func(*state, command) error
}

//Runs a given command with the provided state/config if it exists, Ex: c.comms[name](s, cmd)
func (c *commands) Run(s *state, cmd command) error {
	f, ok := c.registeredComms[cmd.name]
	//f = function to run from comms
	if !ok {
		return errors.New("command not found")
	}
	//returns function with submitted state and command
	return f(s, cmd)
}

//Registers a new handler function for a command name
func (c *commands) Register(name string, f func(*state, command) error) {
	c.registeredComms[name] = f
} 


package main

import (
	"fmt"
)

//Users set user to update state with new user config
func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return fmt.Errorf("No username submitted")
	}
	//gets username from user input
	name := cmd.args[0]

	//attempts to update gatorconfig with user data
	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldnt set current user: %w", err)
	}
	
	fmt.Printf("User %s switched successfully", name)
	return nil
}
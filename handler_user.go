package main

import (
	//context carries info between api/process boundries
	"context"
	"fmt"
	"time"
	"github.com/Sebastien-Johnson/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

//Users set user to update state with new user config
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("No username submitted")
	}
	//gets username from user input
	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("Couldn't find user: %w", err)
	}

	//attempts to update gatorconfig with user data
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldnt set current user: %w", err)
	}
	
	fmt.Printf("User %s switched successfully", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	//check for name in args
	if len(cmd.args) != 1{
		return fmt.Errorf("No username submitted")
	}
	//get name from args
	name := cmd.args[0]
	//create new user with context var and params
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: name,
	})

	if err != nil {
		return fmt.Errorf("Could not create user: %w", err)
	}

	//checks if username is in use or sets to gatorconfig file
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return  fmt.Errorf("Could not set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
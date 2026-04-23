package main

import (
	"github.com/Sebastien-Johnson/blog_aggregator/internal/config"
	"fmt"
	"log"
)


func main() {
	//Creates user from reading cfg file
	userCfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//sets username
	err = userCfg.SetUser("Seb")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	//reads/prints updated config with username
	userCfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Print(userCfg)
}

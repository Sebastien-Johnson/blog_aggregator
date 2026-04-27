package main
//import driver, underscore state that only effects and not direct use is needed
import _ "github.com/lib/pq"

import (
	"log"
	"os"
	"database/sql"
	"github.com/Sebastien-Johnson/blog_aggregator/internal/config"
	//import database
	"github.com/Sebastien-Johnson/blog_aggregator/internal/database"
)


//holds state of one 'config'/user
type state struct{
	//add connection to generated database in /internal/database
	db  *database.Queries
	cfg *config.Config
}



func main() {
	//Creates user from reading cfg file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//load in database with its name and url
	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	//sets up state with read config
	programState := state{
		db: dbQueries,
		cfg: &cfg,
	}

	//sets up commands struct & map
	cmds := commands{
		registeredComms: make(map[string]func(*state, command) error),
	}
	
	//register new user CLI commands
	cmds.Register("login", handlerLogin)
	cmds.Register("register", handlerRegister)
	cmds.Register("reset", handlerReset)
	cmds.Register("users", handlerUsers)
	cmds.Register("agg", handlerAgg)
	cmds.Register("addfeed", handlerAddFeed)
	cmds.Register("feeds", handlerFeeds)

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
		log.Fatal(err)
	}
}

# blog_aggregator
Boot.dev blog aggre'Gator' guided project

Requirements:
- Golang 
    - Install:  curl -sS https://webi.sh/golang | sh
- Postgres
    - Install:  sudo apt update
                sudo apt install postgresql postgresql-contrib

Installing aggreGator CLI
go install https://github.com/Sebastien-Johnson/blog_aggregator

Setting up config file
- Setup json file in home directory '~/.gatorconfig.json'

{
  "db_url": "postgres://example"
}

- Export a config struct representing this file structure
- Create a config.go file in internal/config within your project root and export to main
- Export a Read function to read the gatorconfig file and convert it to a config struct
- Export a SetUser method on the config struct that adds a current_user_name field and write its contents into the gatorconfig file
- Update the 'main' function to red the config, set the username and printe the config contents to the terminal
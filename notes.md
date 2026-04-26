PostgresSql
- Open source database
- Default server :5432
- 'System user' password seperate form 'database user' password

Migrations
- A set of changes for a database table
- 'up' migrations progress the db schema
- 'down' migrations regress the db schema in the case something breaks

Goose
- Go CLI tool to manage db migrations
- Migrations are .sql files with queries and comments
- migrations made from go cli
- Changes made manually from within the psql server are not seen/known by goose cli commands

SQLC
- Generates go code form sql queries that applications can use to interact with databases
- Configured in a .yaml file in project root
- must tell it where to look for schema and queries and where to post the generated code

Postgres driver
- Shows program how to talk to db
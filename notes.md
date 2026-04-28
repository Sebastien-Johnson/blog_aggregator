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

RSS Feed
- "Really Simple Syndication" 
- A way to get the latest content from a website in a structured format (XML)
- XML is unmarshalled into structs

'Many to Many' unique constrain ex:
CREATE TABLE product_suppliers (
  product_id INTEGER,
  supplier_id INTEGER,
  UNIQUE(product_id, supplier_id),
  FOREIGN KEY (product_id) REFERENCES products (id),
  FOREIGN KEY (supplier_id) REFERENCES suppliers (id)
);


SELECT
    feed_follows.*,
    <something>.name AS feed_name,
    <something>.name AS user_name
FROM feed_follows
INNER JOIN <table> ON <feed_follows column> = <table column>
INNER JOIN <table> ON <feed_follows column> = <table column>
WHERE <which column> = $1;


Middleware
- A way to wrap a function with additional functionality. It is a common pattern that allows us to write DRY code.
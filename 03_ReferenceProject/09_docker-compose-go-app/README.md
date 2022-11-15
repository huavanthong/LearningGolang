# Reference
    Original tutorial. Refer: [here](https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/)

# Getting Started
Step 1: Let's also init our Go module from within the container, run the follow command 
```
$ docker compose run --rm app go mod init github.com/huavanthong/MasterGolang/03_ReferenceProject/09_docker-compose-go-app
```

Step 2: We should run our Air init from inside of our app container 
```
$ docker compose run --rm app air init

  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ 1.40.4, built with Go 1.18.3

.air.toml file created to the current directory with the default settings
```

Step 3: Docker compose up to start all services
```
docker compose up
```

Step 4: Run our application
Open brower, and enter:
```
http://localhost:3000/

Hello World
```
Or run curl command:
```bash
curl -F "name=my-item" http://localhost:3000/items
curl -i -X POST -H "Content-Type: application/json" -d '{"name":"my-item"}' http://localhost:3000/items
```
### Set up GO application project with AIR
Init our Go module from within the container, run the follow command 
```bash
$ docker compose run --rm app go mod init github.com/huavanthong/MasterGolang/03_ReferenceProject/09_docker-compose-go-app
```

Should run our Air init from inside of our app container.
```bash
$ docker compose run --rm app air init
```
### Usage knowledge
Item 1: If you want to run your application by docker command
```bash
$ docker compose run --rm app go mod tidy
```

Item 2: Suppose you got an error as below, and if you update your app, and you want to update on your docker images.
**Error**
```
goapp               | go: downloading github.com/russross/blackfriday/v2 v2.1.0  
goapp               | apiserver/items.go:8:2: no required module provides package github.com/bobbytables/go-and-compose/storage; to add it:
goapp               |   go get github.com/bobbytables/go-and-compose/storage     
goapp               | storage/storage.go:8:2: no required module provides package github.com/lib/pq; to add it:
goapp               |   go get github.com/lib/pq
goapp               | failed to build, error: exit status 1
```
**Solution**
```
$ docker compose run --rm app go get github.com/bobbytables/go-and-compose/storage
$ docker compose run --rm app go get github.com/lib/pq
```

### Database and Migration knowledge
Item 1: To run a migration against the database.
```
$ docker compose --profile tools run migrate
```

Item 2: To create items by migrating to database
```
$ docker compose --profile tools run create-migration create_items
[+] Running 1/0
 ⠿ Container go-and-compose_db_1 Created
[+] Running 1/1
 ⠿ Container go-and-compose_db_1  Started
/tmp/migrations/20210828163618_create_items.up.sql
/tmp/migrations/20210828163618_create_items.down.sql
```

Item 3: Up my table to database
```sql
CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;

CREATE TABLE items(
  id uuid DEFAULT public.gen_random_uuid() NOT NULL,
  name character varying NOT NULL
)
```
Error at Item 3: ERROR:  function public.gen_random_uuid() does not exist. More details: [here](https://stackoverflow.com/questions/35959265/postgresql-function-gen-random-uuid-not-working).  
To fix this error
```sql
CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;

CREATE TABLE items(
  id uuid DEFAULT gen_random_uuid() NOT NULL,
  name character varying NOT NULL
)
```

Item 4: Down my table to database
```sql
DROP TABLE items;
DROP EXTENSION pgcrypto;
```
Item 5: To migrate our database to new version
```bash
$ docker compose --profile tools run migrate
```

Item 6: To access database on Postgres image
```
$ docker compose exec db psql -U local-dev -d api
```

Item 6: List all items from our database
```
api=# \d items;
```
### Docker Compose Import
There are a few things to callout about this docker-compose.yml so we understand how these gears mesh together.
    - The services.app.build.target value is set to "dev" - This is the same dev that is in our Dockerfile stages.
    - The services.app.volumes[0] value is mounting the current directory to the same WORKDIR in our Dockerfile "dev" stage.

### TOML knowledge with AIR command
#### full_bin
To tell air to run the command with start as an argument.
```diff
- full_bin = ""
+ full_bin = "./tmp/main start"
```
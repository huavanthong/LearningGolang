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

Step 3: Trying it all out
### Docker Compose Import
There are a few things to callout about this docker-compose.yml so we understand how these gears mesh together.
    - The services.app.build.target value is set to "dev" - This is the same dev that is in our Dockerfile stages.
    - The services.app.volumes[0] value is mounting the current directory to the same WORKDIR in our Dockerfile "dev" stage.


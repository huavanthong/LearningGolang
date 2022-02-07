This is the example repository for [this blog post](https://tutorialedge.net/golang/go-docker-tutorial/)


## Setup  
### Build and run image
To get help from docker
```
$ docker --help
```

To build docker on window. Open CML and enter:
```
$ docker build -t my-go-app .
```
To run image and pass in the porst we want to map to and the image we wish to run.
When you run this -p option, means that we're gonna bins a port on my machine to a port within the docker application.
Therefore, this command below show you: port 8080 on your machine is gonna map to port 8081 inside.
So if you hit localhost 8080 on brower, it's actually going to route that to port 8081 inside the specific container. 
```
$ docker run -p 8080:8081 -it my-go-app
```
To check the result:
```
Open brower:  
  http://localhost:8080/
or
Command line: 
$ curl http://localhost:8080/

```
### Check image
To verify that our image exists on our machine 
```
$ docker images my-go-app
```
To check list of the existence images on docker.
```
$ docker images 
$ docker images -a
```

### Check container  
To check all of containers in images.
```
$docker container ls
CONTAINER ID   IMAGE                                 COMMAND                  CREATED          STATUS          PORTS                NAMES
95700e0f6884   my-go-app:latest                      "/app/main"              26 minutes ago   Up 26 minutes                        nervous_colden
```
To check all container running on all images
```
$docker ps -a
D:\Working\02_Learning\03_ReferenceProject\MasterGolang\03_ReferenceProject\07_Docker-compose-Go-Nodejs>docker ps -a
CONTAINER ID   IMAGE                                 COMMAND                  CREATED          STATUS                      PORTS                NAMES
dcacb7a35b8f   af1205224676                          "/bin/sh -c ./main"      17 minutes ago   Exited (2) 11 minutes ago                        07_docker-compose-go-nodejs_webreactjs_1
e961a2ee38d0   my-go-app:latest                      "/app/main"              4 weeks ago      Exited (255) 4 weeks ago                         naughty_wu
0d2bb06b5aa4   my-go-app                             "/app/main"              4 weeks ago      Exited (2) 4 weeks ago                           hopeful_lamarr
8344e953d36c   docker101tutorial:latest              "/docker-entrypoint.…"   4 weeks ago      Exited (255) 4 weeks ago    80/tcp               beautiful_liskov
2aa785f1317f   docker/dev-environments-go:stable-1   "sleep infinity"         4 weeks ago      Exited (255) 4 weeks ago                         funny_curran
3b8e4fe7797c   docker101tutorial                     "/docker-entrypoint.…"   4 weeks ago      Exited (255) 4 weeks ago    0.0.0.0:80->80/tcp   docker-tutorial
a5a6baaea1db   alpine/git                            "git clone https://g…"   4 weeks ago      Exited (0) 4 weeks ago                           repo
```
### Check volumne
To check all volumne in docker
```
$docker volume ls
D:\Working\02_Learning\03_ReferenceProject\MasterGolang\03_ReferenceProject\07_Docker-compose-Go-Nodejs>docker volume ls
DRIVER    VOLUME NAME
local     5767f1bc77df3f031bc1c73d2885fe402e567963bf2112b4b5c695f31c9f3f65
local     volume-single-dev-env-affectionate_hoover
local     vsCodeServerVolume-single-dev-env-affectionate_hoover
```
To remove volumne on docker  
```
$docker volume rm volume_name volume_name
```
  
### Remove images
To stop container on images
```
$ docker stop 95700e0f6884
```
To remove container in images
```
$ docker rm 95700e0f6884
```
To remove image on docker
```
$ docker rmi my-go-app
=> If it happened error: "conflict: unable to remove repository reference"
$ docker rmi -f my-go-app
```
To remove all images  
```
$ docker rmi $(docker images -a -q)
```
## Issue knowledge
## 1. Issue related to build docker on Window
#### Problem
```
$ docker build -t my-go-app .

=> error: 
error during connect: This error may indicate that the docker daemon is not running.: ........  The system cannot find the file specified.

```
#### Solution
Step 1: Run Docker  
Step 2: Build docker again  

## 2. Docker Desktop Is Shutting Down
#### Problem
```
Open docker by start UI
or
Open docker by command line (CLI)
$ docker build -t my-go-app .

=> error: Docker Desktop Is Shutting Down

```
#### Solution
Solve it following steps:  
Step 1: Docker should be removed from the "Add or remove programmes" list.  
Step 2: Start your computer again.  
Step 3: Install Docker with Administrative privileges (and not by running the installer directly)  
Step 4: Run Docker as Admin  
  
Refer:   https://latestnews.fresherslive.com/articles/docker-failed-to-initialize-docker-desktop-is-shutting-down-a-guide-on-why-docker-failed-to-initialize-293394

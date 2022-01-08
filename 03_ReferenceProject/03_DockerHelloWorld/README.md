This is the example repository for [this blog post](https://tutorialedge.net/golang/go-docker-tutorial/)


## Setup
To get help from docker
```
$ docker --help
```

To build docker on window. Open CML and enter:
```
$ docker build -t my-go-app .
```

To verify that our image exists on our machine 
```
$ docker images
```

To run image and pass in the porst we want to map to and the image we wish to run
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

To check all of containers in images.
```
$docker container ls
CONTAINER ID   IMAGE                                 COMMAND                  CREATED          STATUS          PORTS                NAMES
95700e0f6884   my-go-app:latest                      "/app/main"              26 minutes ago   Up 26 minutes                        nervous_colden
```

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

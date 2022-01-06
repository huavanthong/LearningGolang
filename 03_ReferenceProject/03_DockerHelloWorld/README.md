This is the example repository for [this blog post](https://tutorialedge.net/golang/go-docker-tutorial/)


## Setup
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
$  docker run -p 8080:8081 -it my-go-app
```

To check the result:
```
Open brower: http://localhost:8080/
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